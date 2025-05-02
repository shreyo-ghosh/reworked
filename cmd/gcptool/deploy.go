package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// CommandExecutor is an interface for executing commands
type CommandExecutor interface {
	Execute(name string, args ...string) ([]byte, error)
}

// RealCommandExecutor executes actual system commands
type RealCommandExecutor struct{}

func (e *RealCommandExecutor) Execute(name string, args ...string) ([]byte, error) {
	if name == "gcloud" {
		// Use the full path to gcloud.ps1
		gcloudPath := "C:\\Program Files (x86)\\Google\\Cloud SDK\\google-cloud-sdk\\bin\\gcloud.ps1"
		psArgs := []string{
			"-NoProfile",
			"-NonInteractive",
			"-Command",
			"& { " + gcloudPath + " " + strings.Join(args, " ") + " }",
		}
		fmt.Printf("Running PowerShell command: powershell.exe %s\n", strings.Join(psArgs, " "))
		cmd := exec.Command("powershell.exe", psArgs...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("PowerShell error: %v\n", err)
			fmt.Printf("PowerShell output: %s\n", string(output))
		}
		return output, err
	}
	cmd := exec.Command(name, args...)
	return cmd.CombinedOutput()
}

var defaultExecutor CommandExecutor = &RealCommandExecutor{}

// isValidEnvironment checks if the provided environment is valid
func isValidEnvironment(env string) bool {
	validEnvs := map[string]bool{
		"sandbox": true,
		"dev":     true,
		"pro":     true,
	}
	return validEnvs[env]
}

// deployFunction deploys a function to the specified environment
func deployFunction(functionName, env, version string) error {
	return deployFunctionWithExecutor(functionName, env, version, defaultExecutor)
}

// deployFunctionWithExecutor deploys a function using the provided command executor
func deployFunctionWithExecutor(functionName, env, version string, executor CommandExecutor) error {
	if !isValidEnvironment(env) {
		return fmt.Errorf("invalid environment: %s", env)
	}

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

	// Prepare deployment command arguments
	args := []string{
		"functions",
		"deploy",
		functionName + suffix,
		"--runtime", "go122",
		"--trigger-http",
		"--allow-unauthenticated",
		"--region", "us-central1",
		"--project", "calm-cab-458210-t2",
		"--memory", memory,
		"--source", filepath.Join("functions", functionName),
		"--entry-point", "HelloWorld",
	}

	// Execute the command
	output, err := executor.Execute("gcloud", args...)
	if err != nil {
		return fmt.Errorf("error deploying function: %v\nOutput: %s", err, string(output))
	}

	fmt.Printf("Deployment successful!\nOutput: %s\n", string(output))
	return nil
}

// DeployCmd represents the deploy command
var DeployCmd = &cobra.Command{
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

		fmt.Printf("Deploying function %s to environment %s in region %s (project: %s)\n",
			functionName, env, region, project)

		err := deployFunction(functionName, env, "1.0.0")
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
	},
}

func init() {
	DeployCmd.Flags().StringP("env", "e", "sandbox", "Environment to deploy to (sandbox/staging/production)")
	DeployCmd.Flags().StringP("region", "r", "us-central1", "GCP region to deploy to")
	DeployCmd.Flags().StringP("project", "p", "calm-cab-458210-t2", "GCP project ID")
}
