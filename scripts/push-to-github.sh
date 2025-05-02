#!/bin/bash

# Initialize git repository if not already done
if [ ! -d ".git" ]; then
    git init
fi

# Add all files
git add .

# Create initial commit
git commit -m "Initial commit: GCP Cloud Function Deployment Tool"

# Add remote repository
git remote add origin https://github.com/yourusername/Carbonquest-assignment.git

# Push to main branch
git push -u origin main

# Create and push tags
git tag v1.0.0
git push --tags 