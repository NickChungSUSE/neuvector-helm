// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudwatchlogs provides the client and types for making API
// requests to Amazon CloudWatch Logs.
//
// You can use Amazon CloudWatch Logs to monitor, store, and access your log
// files from EC2 instances, CloudTrail, and other sources. You can then retrieve
// the associated log data from CloudWatch Logs using the CloudWatch console,
// CloudWatch Logs commands in the Amazon Web Services CLI, CloudWatch Logs
// API, or CloudWatch Logs SDK.
//
// You can use CloudWatch Logs to:
//
//   - Monitor logs from EC2 instances in real-time: You can use CloudWatch
//     Logs to monitor applications and systems using log data. For example,
//     CloudWatch Logs can track the number of errors that occur in your application
//     logs and send you a notification whenever the rate of errors exceeds a
//     threshold that you specify. CloudWatch Logs uses your log data for monitoring
//     so no code changes are required. For example, you can monitor application
//     logs for specific literal terms (such as "NullReferenceException") or
//     count the number of occurrences of a literal term at a particular position
//     in log data (such as "404" status codes in an Apache access log). When
//     the term you are searching for is found, CloudWatch Logs reports the data
//     to a CloudWatch metric that you specify.
//
//   - Monitor CloudTrail logged events: You can create alarms in CloudWatch
//     and receive notifications of particular API activity as captured by CloudTrail.
//     You can use the notification to perform troubleshooting.
//
//   - Archive log data: You can use CloudWatch Logs to store your log data
//     in highly durable storage. You can change the log retention setting so
//     that any log events older than this setting are automatically deleted.
//     The CloudWatch Logs agent makes it easy to quickly send both rotated and
//     non-rotated log data off of a host and into the log service. You can then
//     access the raw log data when you need it.
//
// See https://docs.aws.amazon.com/goto/WebAPI/logs-2014-03-28 for more information on this service.
//
// See cloudwatchlogs package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudwatchlogs/
//
// # Using the Client
//
// To contact Amazon CloudWatch Logs with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon CloudWatch Logs client CloudWatchLogs for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudwatchlogs/#New
package cloudwatchlogs