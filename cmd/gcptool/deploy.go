package main

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy [function-name]",
	Short: "Deploy a Cloud Function",
	Long: `Deploy a Cloud Function to Google Cloud Platform.
Specify the function name and environment to deploy to.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		functionName := args[0]
		env, _ := cmd.Flags().GetString("env")
		region, _ := cmd.Flags().GetString("region")
		project, _ := cmd.Flags().GetString("project")

		// Set function suffix based on environment
		suffix := ""
		switch env {
		case "production":
			suffix = "-prod"
		case "staging":
			suffix = "-staging"
		default:
			suffix = "-sandbox"
		}

		// Set memory based on environment
		memory := "128MB"
		switch env {
		case "production":
			memory = "512MB"
		case "staging":
			memory = "256MB"
		}

		fmt.Printf("Deploying function %s to environment %s in region %s (project: %s)\n",
			functionName, env, region, project)

		// Construct the deployment command
		deployCmd := exec.Command("gcloud", "functions", "deploy",
			functionName+suffix,
			"--runtime", "go121",
			"--trigger-http",
			"--allow-unauthenticated",
			"--region", region,
			"--project", project,
			"--memory", memory,
			"--source", filepath.Join("functions", functionName),
			"--entry-point", "HelloWorld",
		)

		// Execute the command
		output, err := deployCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error deploying function: %v\nOutput: %s\n", err, string(output))
			return
		}

		fmt.Printf("Deployment successful!\nOutput: %s\n", string(output))
	},
}

func init() {
	deployCmd.Flags().StringP("env", "e", "sandbox", "Environment to deploy to (sandbox/staging/production)")
	deployCmd.Flags().StringP("region", "r", "us-central1", "GCP region to deploy to")
	deployCmd.Flags().StringP("project", "p", "calm-cab-458210-t2", "GCP project ID")
}
