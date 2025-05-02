module github.com/shreyo-ghosh/carbonquest-assignment/functions/hello-world

go 1.22

require (
	github.com/shreyo-ghosh/carbonquest-assignment/functions/hello-world/pkg/errors v0.0.0
	github.com/shreyo-ghosh/carbonquest-assignment/functions/hello-world/pkg/monitoring v0.0.0
)

replace (
	github.com/shreyo-ghosh/carbonquest-assignment/functions/hello-world/pkg/errors => ./pkg/errors
	github.com/shreyo-ghosh/carbonquest-assignment/functions/hello-world/pkg/monitoring => ./pkg/monitoring
) 