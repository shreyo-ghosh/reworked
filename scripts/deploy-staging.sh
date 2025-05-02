#!/bin/bash

# Set variables
PROJECT_ID="calm-cab-458210-t2"
FUNCTION_NAME="hello-world-staging"
REGION="us-central1"
RUNTIME="go121"
MEMORY="256MB"
TIMEOUT="60s"
MIN_INSTANCES="0"
MAX_INSTANCES="10"

# Enable required APIs
echo "Enabling required APIs..."
gcloud services enable cloudfunctions.googleapis.com
gcloud services enable cloudbuild.googleapis.com
gcloud services enable artifactregistry.googleapis.com

# Deploy the function
echo "Deploying function to staging..."
gcloud functions deploy $FUNCTION_NAME \
    --runtime $RUNTIME \
    --trigger-http \
    --allow-unauthenticated \
    --region $REGION \
    --source functions/hello-world \
    --entry-point HelloWorld \
    --memory $MEMORY \
    --timeout $TIMEOUT \
    --min-instances $MIN_INSTANCES \
    --max-instances $MAX_INSTANCES \
    --project $PROJECT_ID

# Get function URL
FUNCTION_URL=$(gcloud functions describe $FUNCTION_NAME \
    --region $REGION \
    --format='get(httpsTrigger.url)')

echo "Function deployed successfully!"
echo "Function URL: $FUNCTION_URL"

# Test the function
echo "Testing the function..."
curl -s $FUNCTION_URL

# Check deployment status
echo "Checking deployment status..."
gcloud functions describe $FUNCTION_NAME \
    --region $REGION \
    --format='table(name,status,versionId,updateTime)' 