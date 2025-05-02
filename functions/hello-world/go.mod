module carbonquest

go 1.22

require (
	carbonquest/pkg/errors v0.0.0
	carbonquest/pkg/monitoring v0.0.0
)

replace (
	carbonquest/pkg/errors => ./pkg/errors
	carbonquest/pkg/monitoring => ./pkg/monitoring
) 