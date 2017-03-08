// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudwatchlogsiface provides an interface to enable mocking the Amazon CloudWatch Logs service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudwatchlogsiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

// CloudWatchLogsAPI provides an interface to enable mocking the
// cloudwatchlogs.CloudWatchLogs service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon CloudWatch Logs.
//    func myFunc(svc cloudwatchlogsiface.CloudWatchLogsAPI) bool {
//        // Make svc.CancelExportTask request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cloudwatchlogs.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCloudWatchLogsClient struct {
//        cloudwatchlogsiface.CloudWatchLogsAPI
//    }
//    func (m *mockCloudWatchLogsClient) CancelExportTask(input *cloudwatchlogs.CancelExportTaskInput) (*cloudwatchlogs.CancelExportTaskOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCloudWatchLogsClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type CloudWatchLogsAPI interface {
	CancelExportTaskRequest(*cloudwatchlogs.CancelExportTaskInput) (*request.Request, *cloudwatchlogs.CancelExportTaskOutput)

	CancelExportTask(*cloudwatchlogs.CancelExportTaskInput) (*cloudwatchlogs.CancelExportTaskOutput, error)

	CreateExportTaskRequest(*cloudwatchlogs.CreateExportTaskInput) (*request.Request, *cloudwatchlogs.CreateExportTaskOutput)

	CreateExportTask(*cloudwatchlogs.CreateExportTaskInput) (*cloudwatchlogs.CreateExportTaskOutput, error)

	CreateLogGroupRequest(*cloudwatchlogs.CreateLogGroupInput) (*request.Request, *cloudwatchlogs.CreateLogGroupOutput)

	CreateLogGroup(*cloudwatchlogs.CreateLogGroupInput) (*cloudwatchlogs.CreateLogGroupOutput, error)

	CreateLogStreamRequest(*cloudwatchlogs.CreateLogStreamInput) (*request.Request, *cloudwatchlogs.CreateLogStreamOutput)

	CreateLogStream(*cloudwatchlogs.CreateLogStreamInput) (*cloudwatchlogs.CreateLogStreamOutput, error)

	DeleteDestinationRequest(*cloudwatchlogs.DeleteDestinationInput) (*request.Request, *cloudwatchlogs.DeleteDestinationOutput)

	DeleteDestination(*cloudwatchlogs.DeleteDestinationInput) (*cloudwatchlogs.DeleteDestinationOutput, error)

	DeleteLogGroupRequest(*cloudwatchlogs.DeleteLogGroupInput) (*request.Request, *cloudwatchlogs.DeleteLogGroupOutput)

	DeleteLogGroup(*cloudwatchlogs.DeleteLogGroupInput) (*cloudwatchlogs.DeleteLogGroupOutput, error)

	DeleteLogStreamRequest(*cloudwatchlogs.DeleteLogStreamInput) (*request.Request, *cloudwatchlogs.DeleteLogStreamOutput)

	DeleteLogStream(*cloudwatchlogs.DeleteLogStreamInput) (*cloudwatchlogs.DeleteLogStreamOutput, error)

	DeleteMetricFilterRequest(*cloudwatchlogs.DeleteMetricFilterInput) (*request.Request, *cloudwatchlogs.DeleteMetricFilterOutput)

	DeleteMetricFilter(*cloudwatchlogs.DeleteMetricFilterInput) (*cloudwatchlogs.DeleteMetricFilterOutput, error)

	DeleteRetentionPolicyRequest(*cloudwatchlogs.DeleteRetentionPolicyInput) (*request.Request, *cloudwatchlogs.DeleteRetentionPolicyOutput)

	DeleteRetentionPolicy(*cloudwatchlogs.DeleteRetentionPolicyInput) (*cloudwatchlogs.DeleteRetentionPolicyOutput, error)

	DeleteSubscriptionFilterRequest(*cloudwatchlogs.DeleteSubscriptionFilterInput) (*request.Request, *cloudwatchlogs.DeleteSubscriptionFilterOutput)

	DeleteSubscriptionFilter(*cloudwatchlogs.DeleteSubscriptionFilterInput) (*cloudwatchlogs.DeleteSubscriptionFilterOutput, error)

	DescribeDestinationsRequest(*cloudwatchlogs.DescribeDestinationsInput) (*request.Request, *cloudwatchlogs.DescribeDestinationsOutput)

	DescribeDestinations(*cloudwatchlogs.DescribeDestinationsInput) (*cloudwatchlogs.DescribeDestinationsOutput, error)

	DescribeDestinationsPages(*cloudwatchlogs.DescribeDestinationsInput, func(*cloudwatchlogs.DescribeDestinationsOutput, bool) bool) error

	DescribeExportTasksRequest(*cloudwatchlogs.DescribeExportTasksInput) (*request.Request, *cloudwatchlogs.DescribeExportTasksOutput)

	DescribeExportTasks(*cloudwatchlogs.DescribeExportTasksInput) (*cloudwatchlogs.DescribeExportTasksOutput, error)

	DescribeLogGroupsRequest(*cloudwatchlogs.DescribeLogGroupsInput) (*request.Request, *cloudwatchlogs.DescribeLogGroupsOutput)

	DescribeLogGroups(*cloudwatchlogs.DescribeLogGroupsInput) (*cloudwatchlogs.DescribeLogGroupsOutput, error)

	DescribeLogGroupsPages(*cloudwatchlogs.DescribeLogGroupsInput, func(*cloudwatchlogs.DescribeLogGroupsOutput, bool) bool) error

	DescribeLogStreamsRequest(*cloudwatchlogs.DescribeLogStreamsInput) (*request.Request, *cloudwatchlogs.DescribeLogStreamsOutput)

	DescribeLogStreams(*cloudwatchlogs.DescribeLogStreamsInput) (*cloudwatchlogs.DescribeLogStreamsOutput, error)

	DescribeLogStreamsPages(*cloudwatchlogs.DescribeLogStreamsInput, func(*cloudwatchlogs.DescribeLogStreamsOutput, bool) bool) error

	DescribeMetricFiltersRequest(*cloudwatchlogs.DescribeMetricFiltersInput) (*request.Request, *cloudwatchlogs.DescribeMetricFiltersOutput)

	DescribeMetricFilters(*cloudwatchlogs.DescribeMetricFiltersInput) (*cloudwatchlogs.DescribeMetricFiltersOutput, error)

	DescribeMetricFiltersPages(*cloudwatchlogs.DescribeMetricFiltersInput, func(*cloudwatchlogs.DescribeMetricFiltersOutput, bool) bool) error

	DescribeSubscriptionFiltersRequest(*cloudwatchlogs.DescribeSubscriptionFiltersInput) (*request.Request, *cloudwatchlogs.DescribeSubscriptionFiltersOutput)

	DescribeSubscriptionFilters(*cloudwatchlogs.DescribeSubscriptionFiltersInput) (*cloudwatchlogs.DescribeSubscriptionFiltersOutput, error)

	DescribeSubscriptionFiltersPages(*cloudwatchlogs.DescribeSubscriptionFiltersInput, func(*cloudwatchlogs.DescribeSubscriptionFiltersOutput, bool) bool) error

	FilterLogEventsRequest(*cloudwatchlogs.FilterLogEventsInput) (*request.Request, *cloudwatchlogs.FilterLogEventsOutput)

	FilterLogEvents(*cloudwatchlogs.FilterLogEventsInput) (*cloudwatchlogs.FilterLogEventsOutput, error)

	FilterLogEventsPages(*cloudwatchlogs.FilterLogEventsInput, func(*cloudwatchlogs.FilterLogEventsOutput, bool) bool) error

	GetLogEventsRequest(*cloudwatchlogs.GetLogEventsInput) (*request.Request, *cloudwatchlogs.GetLogEventsOutput)

	GetLogEvents(*cloudwatchlogs.GetLogEventsInput) (*cloudwatchlogs.GetLogEventsOutput, error)

	GetLogEventsPages(*cloudwatchlogs.GetLogEventsInput, func(*cloudwatchlogs.GetLogEventsOutput, bool) bool) error

	ListTagsLogGroupRequest(*cloudwatchlogs.ListTagsLogGroupInput) (*request.Request, *cloudwatchlogs.ListTagsLogGroupOutput)

	ListTagsLogGroup(*cloudwatchlogs.ListTagsLogGroupInput) (*cloudwatchlogs.ListTagsLogGroupOutput, error)

	PutDestinationRequest(*cloudwatchlogs.PutDestinationInput) (*request.Request, *cloudwatchlogs.PutDestinationOutput)

	PutDestination(*cloudwatchlogs.PutDestinationInput) (*cloudwatchlogs.PutDestinationOutput, error)

	PutDestinationPolicyRequest(*cloudwatchlogs.PutDestinationPolicyInput) (*request.Request, *cloudwatchlogs.PutDestinationPolicyOutput)

	PutDestinationPolicy(*cloudwatchlogs.PutDestinationPolicyInput) (*cloudwatchlogs.PutDestinationPolicyOutput, error)

	PutLogEventsRequest(*cloudwatchlogs.PutLogEventsInput) (*request.Request, *cloudwatchlogs.PutLogEventsOutput)

	PutLogEvents(*cloudwatchlogs.PutLogEventsInput) (*cloudwatchlogs.PutLogEventsOutput, error)

	PutMetricFilterRequest(*cloudwatchlogs.PutMetricFilterInput) (*request.Request, *cloudwatchlogs.PutMetricFilterOutput)

	PutMetricFilter(*cloudwatchlogs.PutMetricFilterInput) (*cloudwatchlogs.PutMetricFilterOutput, error)

	PutRetentionPolicyRequest(*cloudwatchlogs.PutRetentionPolicyInput) (*request.Request, *cloudwatchlogs.PutRetentionPolicyOutput)

	PutRetentionPolicy(*cloudwatchlogs.PutRetentionPolicyInput) (*cloudwatchlogs.PutRetentionPolicyOutput, error)

	PutSubscriptionFilterRequest(*cloudwatchlogs.PutSubscriptionFilterInput) (*request.Request, *cloudwatchlogs.PutSubscriptionFilterOutput)

	PutSubscriptionFilter(*cloudwatchlogs.PutSubscriptionFilterInput) (*cloudwatchlogs.PutSubscriptionFilterOutput, error)

	TagLogGroupRequest(*cloudwatchlogs.TagLogGroupInput) (*request.Request, *cloudwatchlogs.TagLogGroupOutput)

	TagLogGroup(*cloudwatchlogs.TagLogGroupInput) (*cloudwatchlogs.TagLogGroupOutput, error)

	TestMetricFilterRequest(*cloudwatchlogs.TestMetricFilterInput) (*request.Request, *cloudwatchlogs.TestMetricFilterOutput)

	TestMetricFilter(*cloudwatchlogs.TestMetricFilterInput) (*cloudwatchlogs.TestMetricFilterOutput, error)

	UntagLogGroupRequest(*cloudwatchlogs.UntagLogGroupInput) (*request.Request, *cloudwatchlogs.UntagLogGroupOutput)

	UntagLogGroup(*cloudwatchlogs.UntagLogGroupInput) (*cloudwatchlogs.UntagLogGroupOutput, error)
}

var _ CloudWatchLogsAPI = (*cloudwatchlogs.CloudWatchLogs)(nil)
