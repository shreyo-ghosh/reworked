package monitoring

import (
	"context"
	"fmt"
	"log"
	"time"

	monitoring "cloud.google.com/go/monitoring/apiv3/v2"
	monitoringpb "cloud.google.com/go/monitoring/apiv3/v2/monitoringpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// MetricsClient wraps the Cloud Monitoring client
type MetricsClient struct {
	client    *monitoring.MetricClient
	projectID string
}

// NewMetricsClient creates a new metrics client
func NewMetricsClient(ctx context.Context, projectID string) (*MetricsClient, error) {
	client, err := monitoring.NewMetricClient(ctx)
	if err != nil {
		return nil, err
	}
	return &MetricsClient{
		client:    client,
		projectID: projectID,
	}, nil
}

// RecordLatency records function execution latency
func (m *MetricsClient) RecordLatency(ctx context.Context, functionName string, latency time.Duration) error {
	now := time.Now()
	req := &monitoringpb.CreateTimeSeriesRequest{
		Name: fmt.Sprintf("projects/%s", m.projectID),
		TimeSeries: []*monitoringpb.TimeSeries{
			{
				Metric: &monitoringpb.Metric{
					Type: "custom.googleapis.com/function/latency",
					Labels: map[string]string{
						"function_name": functionName,
					},
				},
				Points: []*monitoringpb.Point{
					{
						Interval: &monitoringpb.TimeInterval{
							EndTime: &timestamppb.Timestamp{
								Seconds: now.Unix(),
							},
						},
						Value: &monitoringpb.TypedValue{
							Value: &monitoringpb.TypedValue_DoubleValue{
								DoubleValue: latency.Seconds(),
							},
						},
					},
				},
			},
		},
	}

	return m.client.CreateTimeSeries(ctx, req)
}

// LogFunctionExecution logs function execution details
func LogFunctionExecution(functionName string, startTime time.Time, err error) {
	latency := time.Since(startTime)
	log.Printf("Function: %s, Latency: %v, Error: %v", functionName, latency, err)
}
