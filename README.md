# Cloud Function Deployment Tool

A Go-based tool for deploying and managing Google Cloud Functions.

## Features

- Deploy Cloud Functions to different environments (sandbox, dev, pro)
- Simple command-line interface
- GitHub Actions integration for automated deployments

## Prerequisites

- Go 1.22 or later
- Google Cloud SDK
- A Google Cloud project with Cloud Functions API enabled
- GitHub repository with Workload Identity Federation configured

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/shreyo-ghosh/reworked.git
   cd reworked
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Configure GitHub Secrets:
   - `GCP_PROJECT_ID`: Your Google Cloud project ID
   - `GCP_WORKLOAD_IDENTITY_PROVIDER`: The full identifier of your Workload Identity Provider
   - `GCP_SERVICE_ACCOUNT`: The email of your service account

## Local Development

To run the hello-world function locally:

```bash
cd functions/hello-world
go run main.go
```

The function will be available at `http://localhost:8080`.

## Deployment

The application is automatically deployed to Google Cloud Functions when changes are pushed to the main branch. The deployment is handled by GitHub Actions.

### Manual Deployment

To deploy manually:

```bash
go run cmd/gcptool/main.go deploy hello-world --env sandbox
```

## License

MIT
