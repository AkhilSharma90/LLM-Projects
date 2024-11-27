# AI Content Creator Agent

This project automates content creation and social media management using AI, specifically via Google's Gemini API and Facebook. It leverages current news trends and learns from past popular posts to generate engaging content.

## Features

- **AI-Powered Content Generation:** Creates social media posts using the Gemini API, incorporating company information, industry news, and analysis of past successful posts.
- **Facebook Integration:** Posts generated content to Facebook and replies to comments.
- **News Integration:** Utilizes the News API to gather relevant industry news for content generation.
- **Scheduled Posting (In Development):** Designed for integration with Google Cloud Scheduler to automate posting. This feature requires further configuration.
- **Persistent Storage:** Stores company information and generated content in Google Cloud Firestore.

## Technologies

- **Backend:** Go (Golang), Gin web framework
- **AI:** Google Gemini API
- **News:** News API
- **Social Media:** Facebook Graph API
- **Database:** Google Cloud Firestore
- **Scheduling:** Google Cloud Scheduler (Integration in progress)

## Setup and Installation

### Prerequisites

1. **Go:** Install [Go](https://go.dev/doc/install).
2. **Google Cloud Project:** Create a GCP project and enable the Firestore API and, optionally, Cloud Scheduler if you plan to use the scheduling feature.
3. **Service Account:** Create a service account with permissions to access Firestore and Cloud Scheduler (if using). Download its JSON key file.
4. **API Keys:** Obtain API keys for:
   - [Google Gemini API](https://developers.generativeai.google/)
   - [News API](https://newsapi.org/)
   - Facebook Graph API: Set up a Facebook App and obtain an access token with the necessary permissions to manage your Facebook page.

### Project Setup

1. **Clone the repository:**
   ```bash
   git clone
   ```
2. **Change directory:**
   ```bash
   cd hard-content-creator-agent
   ```
3. **Service Account Key:** Place your `service-account-key.json` file in the `internal/infrastructure/config` directory.
4. **Environment Variables:** Create a `.env` file in the project's root directory:
   ```plaintext
   PROJECT_ID=<your-gcp-project-id>
   GEMINI_API_KEY=<your-gemini-api-key>
   NEWS_API_KEY=<your-news-api-key>
   SERVICE_ACCOUNT_KEY_PATH=internal/infrastructure/config/service-account-key.json
   DB_NAME=<your-firestore-db-name>  # Can be any name. Doesn't create the database, just used in the code.
   ```

### Code Adjustments (Important)

- **`internal/domain/usecases/company_usecase.go`:** Replace `<your-backend-server-url>` with the public URL of your backend server (e.g., `http://your-domain.com:8080` or `http://localhost:8080` for local testing).
- **`cmd/api/main.go`:** Uncomment and update the Facebook API initialization:
  ```go
  facebookAPI := api.NewFacebookAPIService()
  socialMediaUsecase := usecases.NewSocialMediaMgtUsecase(contentUsecase, companyUsecase, facebookAPI)
  ```
- **`internal/domain/usecases/social_media_mgt_usecase.go`:** Uncomment the lines of code that make the Facebook post. *Be very careful with this, especially in production.*
- **`internal/adapters/repositories/firestore_company_repo.go`:** Change `Companys` to `companies` in all Firestore-related collections in this file and `firestore_content_repo.go` file too.

### Build and Run

- **Using Makefile (Recommended):**
  ```bash
  make run
  ```
- **Directly:**
  ```bash
  go run cmd/api/main.go
  ```

## API Usage (with cURL)

1. **Register a Company:**
   ```bash
   curl -X POST -H "Content-Type: application/json" \
   -d '{
     "companyName": "Your Company Name",
     "industry": "Your Industry",
     "missionStatement": "Your Mission Statement",
     "brandVoice": "Your Brand Voice",
     "logoURL": "Your Logo URL",
     "targetAudience": "Your Target Audience",
     "contentGoals": ["Goal 1", "Goal 2"],
     "keyMessages": ["Message 1", "Message 2"],
     "facebookPageID": "Your Facebook Page ID",
     "facebookAccessToken": "Your Facebook Access Token",
     "postingFrequency": "Posting Frequency"
   }' http://localhost:8080/company/register
   ```

2. **Get a Company:**
   ```bash
   curl http://localhost:8080/company/<company_id>
   ```

3. **Generate and Post Content:**
   ```bash
   curl -X POST http://localhost:8080/post-content/<company_id>
   ```

4. **Reply to Comments:**
   ```bash
   curl -X POST http://localhost:8080/reply/<company_id>
   ```