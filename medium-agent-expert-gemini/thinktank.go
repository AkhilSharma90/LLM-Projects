package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/tmc/langchaingo/llms"
)

const DefaultConcurrency = 5

type ThinkTank struct {
	Prompt      string
	Model       llms.Model
	Experts     []ExpertInterface
	Concurrency int
}

func NewThinkTank(model llms.Model, experts ...ExpertInterface) *ThinkTank {
	return &ThinkTank{
		Prompt:      "Weigh and judge the experts answers against each other.  Eliminate contradictions. Then provide a comprehensive and definitive answer to the question: %s",
		Model:       model,
		Experts:     experts,
		Concurrency: DefaultConcurrency,
	}
}

type ThinkTankResult struct {
	ExpertResults []*ExpertResult
}

func (tt *ThinkTank) Evaluate(ctx context.Context, input string) *ThinkTankResult {

	var results []*ExpertResult
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, expert := range tt.Experts {
		wg.Add(1)

		go func(expert ExpertInterface) {
			defer wg.Done()
			result, err := expert.Evaluate(ctx, input)
			result.Err = err

			mu.Lock()
			results = append(results, result)
			mu.Unlock()
		}(expert)
	}

	wg.Wait()

	return &ThinkTankResult{
		ExpertResults: results,
	}
}

func (tt *ThinkTank) Answer(ctx context.Context, input string) (string, error) {
	result := tt.Evaluate(ctx, input)

	var expertsReports = ""

	// Next we let each expert chime in
	for _, expertResult := range result.ExpertResults {
		if expertResult.Err == nil {

			expertsReports += fmt.Sprintf("%s: %s\n\n", expertResult.Expert.Name, expertResult.Text)
		}
	}

	// Concatenate all expert reports into one message as context
	expertOpinions := llms.MessageContent{
		Role:  llms.ChatMessageTypeSystem,
		Parts: []llms.ContentPart{llms.TextPart(expertsReports)},
	}

	// Next we instruct the LLM to re-answer the question in light of
	// the new expert opinions
	thinkTankPrompt := fmt.Sprintf(tt.Prompt, input)

	systemMessage := llms.MessageContent{
		Role:  llms.ChatMessageTypeHuman,
		Parts: []llms.ContentPart{llms.TextPart(thinkTankPrompt)},
	}

	var messages = []llms.MessageContent{systemMessage, expertOpinions}

	content, err := tt.Model.GenerateContent(ctx, messages)

	if err != nil {
		return "", err
	}

	return content.Choices[0].Content, nil
}
