# CarbonQuest GCP Cloud Function Deployment Tool

A command line tool for deploying and managing GCP Cloud Functions with automated CI/CD pipeline.

## Features

- Deploy cloud functions to different environments (sandbox, dev, pro)
- Get detailed information about deployed functions
- Automated testing and deployment through GitHub Actions
- Support for versioning and clean builds

## Prerequisites

- Go 1.21 or later
- Google Cloud SDK
- GCP project with Cloud Functions API enabled
- GitHub repository with GCP credentials configured

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/Carbonquest-assignment.git
cd Carbonquest-assignment
```

2. Install dependencies:
```bash
go mod download
```

3. Build the tool:
```bash
go build -o gcptool cmd/gcptool/main.go
```

## Usage

### Deploy a Function
```bash
./gcptool deploy <function-name> -e <environment> -v <revision> -c
```

Options:
- `-e, --environment`: Target environment (sandbox, dev, pro)
- `-v, --revision`: Version/revision number
- `-c, --clean`: Clean and rebuild before deploying

### Describe a Function
```bash
./gcptool describe <function-name>
```

## GitHub Actions Setup

1. Create a new GitHub repository named "Carbonquest-assignment"

2. Add GCP credentials to GitHub Secrets:
   - Go to your repository settings
   - Navigate to Secrets and Variables > Actions
   - Add a new secret named `GCP_CREDENTIALS`
   - Paste your GCP service account key JSON

3. Push your code to the repository:
```bash
git init
git add .
git commit -m "Initial commit"
git branch -M main
git remote add origin https://github.com/yourusername/Carbonquest-assignment.git
git push -u origin main
```

## Development

### Running Tests
```bash
go test -v ./...
```

### Adding New Features
1. Create a new branch
2. Make your changes
3. Run tests
4. Create a pull request

## License

MIT License.
