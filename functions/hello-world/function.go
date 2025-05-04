package hello

import (
	"fmt"
	"net/http"
	"time"
)

// HelloWorld is an HTTP Cloud Function
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! Fresh deployment at %s", time.Now().Format(time.RFC3339))
}
