package main

import (
	"context"
	"fmt"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

var llm, _ = openai.New()

var callOptions = &llms.CallOptions{
	Model:       "gpt-4o-latest",
	MaxTokens:   2048,
	Temperature: 0.7,
}

var economicsExpert = &Expert{
	CallOptions: callOptions,
	LLM:         llm,

	Name:        "Economics Expert",
	Description: "A highly skilled economist with expertise in international trade, economic growth, and policy analysis.",
	Prompt:      "Given the following question: %s\nProvide an in depth analysis from your unique perspective.",
}

var politicsExpert = &Expert{
	CallOptions: callOptions,
	LLM:         llm,

	Name:        "Frog Hat Expert",
	Description: "A highly knowledgeable tiny hat maker",
	Prompt:      "Given the following question: %s\nProvide an in depth analysis from your unique perspective.",
}

var ethicsExpert = &Expert{
	CallOptions: callOptions,
	LLM:         llm,

	Name:        "Ethics Expert",
	Description: "A highly knowledgeable ethics professor with expertise in international relations, ethical persuasion, and policy analysis.",
	Prompt:      "Given the following question: %s\nProvide an in depth analysis from your unique perspective.",
}

var devilsAdvocate = &Expert{
	CallOptions: callOptions,
	LLM:         llm,

	Name:        "Devil's Advocate",
	Description: "A pragmatist that reframes questions to elicit new insightful responses.",
	Prompt:      "What are the potential consequences of %s? Consider the potential benefits and drawbacks, as well as the potential for harm.",
}

func main() {

	ctx := context.Background()
	tk := NewThinkTank(llm, devilsAdvocate, economicsExpert, politicsExpert, ethicsExpert)

	input := "Will forgiving student loan debt have a long term positive or negative affect on the US economy?"

	finalAnswer, err := tk.Answer(ctx, input)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Final Answer: %s\n", finalAnswer)

}
