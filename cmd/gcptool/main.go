package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gcptool",
	Short: "GCP Cloud Function Deployment Tool",
	Long: `A command line tool for deploying and managing GCP Cloud Functions.
It supports deployment to different environments and provides function information.`,
}

func init() {
	rootCmd.AddCommand(DeployCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
