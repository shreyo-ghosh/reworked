package main

import (
	"log"
	"os"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	"github.com/shreyo-ghosh/reworked/functions/hello-world/function"
)

func main() {
	// Register the function
	funcframework.RegisterHTTPFunction("/", function.HelloWorld)

	// Start the function
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
