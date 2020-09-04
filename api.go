package sqsextendedclient

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/query"
)

const opAddPermission = "AddPermission"

// AddPermissionRequest generates a "aws/request.Request" representing the
// client's request for the AddPermission operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See AddPermission for more information on using the AddPermission
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the AddPermissionRequest method.
//    req, resp := client.AddPermissionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/AddPermission
func (c *SQSExtended) AddPermissionRequest(input *AddPermissionInput) (req *request.Request, output *AddPermissionOutput) {
	op := &request.Operation{
		Name:       opAddPermission,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AddPermissionInput{}
	}

	output = &AddPermissionOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(query.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

// AddPermission API operation for Amazon Simple Queue Service.
//
// Adds a permission to a queue for a specific principal (https://docs.aws.amazon.com/general/latest/gr/glos-chap.html#P).
// This allows sharing access to the queue.
//
// When you create a queue, you have full control access rights for the queue.
// Only you, the owner of the queue, can grant or deny permissions to the queue.
// For more information about these permissions, see Allow Developers to Write
// Messages to a Shared Queue (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-writing-an-sqs-policy.html#write-messages-to-shared-queue)
// in the Amazon Simple Queue Service Developer Guide.
//
//    * AddPermission generates a policy for you. You can use SetQueueAttributes
//    to upload your policy. For more information, see Using Custom Policies
//    with the Amazon SQS Access Policy Language (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-creating-custom-policies.html)
//    in the Amazon Simple Queue Service Developer Guide.
//
//    * An Amazon SQS policy can have a maximum of 7 actions.
//
//    * To remove the ability to change queue permissions, you must deny permission
//    to the AddPermission, RemovePermission, and SetQueueAttributes actions
//    in your IAM policy.
//
// Some actions take lists of parameters. These lists are specified using the
// param.n notation. Values of n are integers starting from 1. For example,
// a parameter list with two elements looks like this:
//
// &Attribute.1=first
//
// &Attribute.2=second
//
// Cross-account permissions don't apply to this action. For more information,
// see Grant Cross-Account Permissions to a Role and a User Name (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon Simple Queue Service Developer Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation AddPermission for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeOverLimit "OverLimit"
//   The specified action violates a limit. For example, ReceiveMessage returns
//   this error if the maximum number of inflight messages is reached and AddPermission
//   returns this error if the maximum number of permissions for the queue is
//   reached.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/AddPermission
func (c *SQSExtended) AddPermission(input *AddPermissionInput) (*AddPermissionOutput, error) {
	req, out := c.AddPermissionRequest(input)
	return out, req.Send()
}

// AddPermissionWithContext is the same as AddPermission with the addition of
// the ability to pass a context and additional request options.
//
// See AddPermission for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) AddPermissionWithContext(ctx aws.Context, input *AddPermissionInput, opts ...request.Option) (*AddPermissionOutput, error) {
	req, out := c.AddPermissionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opChangeMessageVisibility = "ChangeMessageVisibility"

// ChangeMessageVisibilityRequest generates a "aws/request.Request" representing the
// client's request for the ChangeMessageVisibility operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ChangeMessageVisibility for more information on using the ChangeMessageVisibility
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ChangeMessageVisibilityRequest method.
//    req, resp := client.ChangeMessageVisibilityRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ChangeMessageVisibility
func (c *SQSExtended) ChangeMessageVisibilityRequest(input *ChangeMessageVisibilityInput) (req *request.Request, output *ChangeMessageVisibilityOutput) {
	op := &request.Operation{
		Name:       opChangeMessageVisibility,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ChangeMessageVisibilityInput{}
	}

	output = &ChangeMessageVisibilityOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(query.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

// ChangeMessageVisibility API operation for Amazon Simple Queue Service.
//
// Changes the visibility timeout of a specified message in a queue to a new
// value. The default visibility timeout for a message is 30 seconds. The minimum
// is 0 seconds. The maximum is 12 hours. For more information, see Visibility
// Timeout (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html)
// in the Amazon Simple Queue Service Developer Guide.
//
// For example, you have a message with a visibility timeout of 5 minutes. After
// 3 minutes, you call ChangeMessageVisibility with a timeout of 10 minutes.
// You can continue to call ChangeMessageVisibility to extend the visibility
// timeout to the maximum allowed time. If you try to extend the visibility
// timeout beyond the maximum, your request is rejected.
//
// An Amazon SQS message has three basic states:
//
// Sent to a queue by a producer.
//
// Received from the queue by a consumer.
//
// Deleted from the queue.
//
// A message is considered to be stored after it is sent to a queue by a producer,
// but not yet received from the queue by a consumer (that is, between states
// 1 and 2). There is no limit to the number of stored messages. A message is
// considered to be in flight after it is received from a queue by a consumer,
// but not yet deleted from the queue (that is, between states 2 and 3). There
// is a limit to the number of inflight messages.
//
// Limits that apply to inflight messages are unrelated to the unlimited number
// of stored messages.
//
// For most standard queues (depending on queue traffic and message backlog),
// there can be a maximum of approximately 120,000 inflight messages (received
// from a queue by a consumer, but not yet deleted from the queue). If you reach
// this limit, Amazon SQS returns the OverLimit error message. To avoid reaching
// the limit, you should delete messages from the queue after they're processed.
// You can also increase the number of queues you use to process your messages.
// To request a limit increase, file a support request (https://console.aws.amazon.com/support/home#/case/create?issueType=service-limit-increase&limitType=service-code-sqs).
//
// For FIFO queues, there can be a maximum of 20,000 inflight messages (received
// from a queue by a consumer, but not yet deleted from the queue). If you reach
// this limit, Amazon SQS returns no error messages.
//
// If you attempt to set the VisibilityTimeout to a value greater than the maximum
// time left, Amazon SQS returns an error. Amazon SQS doesn't automatically
// recalculate and increase the timeout to the maximum remaining time.
//
// Unlike with a queue, when you change the visibility timeout for a specific
// message the timeout value is applied immediately but isn't saved in memory
// for that message. If you don't delete a message after it is received, the
// visibility timeout for the message reverts to the original timeout value
// (not to the value you set using the ChangeMessageVisibility action) the next
// time the message is received.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation ChangeMessageVisibility for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeMessageNotInflight "AWS.SimpleQueueService.MessageNotInflight"
//   The specified message isn't in flight.
//
//   * ErrCodeReceiptHandleIsInvalid "ReceiptHandleIsInvalid"
//   The specified receipt handle isn't valid.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ChangeMessageVisibility
func (c *SQSExtended) ChangeMessageVisibility(input *ChangeMessageVisibilityInput) (*ChangeMessageVisibilityOutput, error) {
	req, out := c.ChangeMessageVisibilityRequest(input)
	return out, req.Send()
}

// ChangeMessageVisibilityWithContext is the same as ChangeMessageVisibility with the addition of
// the ability to pass a context and additional request options.
//
// See ChangeMessageVisibility for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) ChangeMessageVisibilityWithContext(ctx aws.Context, input *ChangeMessageVisibilityInput, opts ...request.Option) (*ChangeMessageVisibilityOutput, error) {
	req, out := c.ChangeMessageVisibilityRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opChangeMessageVisibilityBatch = "ChangeMessageVisibilityBatch"

// ChangeMessageVisibilityBatchRequest generates a "aws/request.Request" representing the
// client's request for the ChangeMessageVisibilityBatch operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ChangeMessageVisibilityBatch for more information on using the ChangeMessageVisibilityBatch
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ChangeMessageVisibilityBatchRequest method.
//    req, resp := client.ChangeMessageVisibilityBatchRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ChangeMessageVisibilityBatch
func (c *SQSExtended) ChangeMessageVisibilityBatchRequest(input *ChangeMessageVisibilityBatchInput) (req *request.Request, output *ChangeMessageVisibilityBatchOutput) {
	op := &request.Operation{
		Name:       opChangeMessageVisibilityBatch,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ChangeMessageVisibilityBatchInput{}
	}

	output = &ChangeMessageVisibilityBatchOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ChangeMessageVisibilityBatch API operation for Amazon Simple Queue Service.
//
// Changes the visibility timeout of multiple messages. This is a batch version
// of ChangeMessageVisibility. The result of the action on each message is reported
// individually in the response. You can send up to 10 ChangeMessageVisibility
// requests with each ChangeMessageVisibilityBatch action.
//
// Because the batch request can result in a combination of successful and unsuccessful
// actions, you should check for batch errors even when the call returns an
// HTTP status code of 200.
//
// Some actions take lists of parameters. These lists are specified using the
// param.n notation. Values of n are integers starting from 1. For example,
// a parameter list with two elements looks like this:
//
// &Attribute.1=first
//
// &Attribute.2=second
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation ChangeMessageVisibilityBatch for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeTooManyEntriesInBatchRequest "AWS.SimpleQueueService.TooManyEntriesInBatchRequest"
//   The batch request contains more entries than permissible.
//
//   * ErrCodeEmptyBatchRequest "AWS.SimpleQueueService.EmptyBatchRequest"
//   The batch request doesn't contain any entries.
//
//   * ErrCodeBatchEntryIdsNotDistinct "AWS.SimpleQueueService.BatchEntryIdsNotDistinct"
//   Two or more batch entries in the request have the same Id.
//
//   * ErrCodeInvalidBatchEntryId "AWS.SimpleQueueService.InvalidBatchEntryId"
//   The Id of a batch entry in a batch request doesn't abide by the specification.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ChangeMessageVisibilityBatch
func (c *SQSExtended) ChangeMessageVisibilityBatch(input *ChangeMessageVisibilityBatchInput) (*ChangeMessageVisibilityBatchOutput, error) {
	req, out := c.ChangeMessageVisibilityBatchRequest(input)
	return out, req.Send()
}

// ChangeMessageVisibilityBatchWithContext is the same as ChangeMessageVisibilityBatch with the addition of
// the ability to pass a context and additional request options.
//
// See ChangeMessageVisibilityBatch for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) ChangeMessageVisibilityBatchWithContext(ctx aws.Context, input *ChangeMessageVisibilityBatchInput, opts ...request.Option) (*ChangeMessageVisibilityBatchOutput, error) {
	req, out := c.ChangeMessageVisibilityBatchRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateQueue = "CreateQueue"

// CreateQueueRequest generates a "aws/request.Request" representing the
// client's request for the CreateQueue operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See CreateQueue for more information on using the CreateQueue
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the CreateQueueRequest method.
//    req, resp := client.CreateQueueRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/CreateQueue
func (c *SQSExtended) CreateQueueRequest(input *CreateQueueInput) (req *request.Request, output *CreateQueueOutput) {
	op := &request.Operation{
		Name:       opCreateQueue,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateQueueInput{}
	}

	output = &CreateQueueOutput{}
	req = c.newRequest(op, input, output)
	return
}

// CreateQueue API operation for Amazon Simple Queue Service.
//
// Creates a new standard or FIFO queue. You can pass one or more attributes
// in the request. Keep the following caveats in mind:
//
//    * If you don't specify the FifoQueue attribute, Amazon SQS creates a standard
//    queue. You can't change the queue type after you create it and you can't
//    convert an existing standard queue into a FIFO queue. You must either
//    create a new FIFO queue for your application or delete your existing standard
//    queue and recreate it as a FIFO queue. For more information, see Moving
//    From a Standard Queue to a FIFO Queue (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-moving)
//    in the Amazon Simple Queue Service Developer Guide.
//
//    * If you don't provide a value for an attribute, the queue is created
//    with the default value for the attribute.
//
//    * If you delete a queue, you must wait at least 60 seconds before creating
//    a queue with the same name.
//
// To successfully create a new queue, you must provide a queue name that adheres
// to the limits related to queues (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/limits-queues.html)
// and is unique within the scope of your queues.
//
// To get the queue URL, use the GetQueueUrl action. GetQueueUrl requires only
// the QueueName parameter. be aware of existing queue names:
//
//    * If you provide the name of an existing queue along with the exact names
//    and values of all the queue's attributes, CreateQueue returns the queue
//    URL for the existing queue.
//
//    * If the queue name, attribute names, or attribute values don't match
//    an existing queue, CreateQueue returns an error.
//
// Some actions take lists of parameters. These lists are specified using the
// param.n notation. Values of n are integers starting from 1. For example,
// a parameter list with two elements looks like this:
//
// &Attribute.1=first
//
// &Attribute.2=second
//
// Cross-account permissions don't apply to this action. For more information,
// see Grant Cross-Account Permissions to a Role and a User Name (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon Simple Queue Service Developer Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation CreateQueue for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeQueueDeletedRecently "AWS.SimpleQueueService.QueueDeletedRecently"
//   You must wait 60 seconds after deleting a queue before you can create another
//   queue with the same name.
//
//   * ErrCodeQueueNameExists "QueueAlreadyExists"
//   A queue with this name already exists. Amazon SQS returns this error only
//   if the request includes attributes whose values differ from those of the
//   existing queue.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/CreateQueue
func (c *SQSExtended) CreateQueue(input *CreateQueueInput) (*CreateQueueOutput, error) {
	req, out := c.CreateQueueRequest(input)
	return out, req.Send()
}

// CreateQueueWithContext is the same as CreateQueue with the addition of
// the ability to pass a context and additional request options.
//
// See CreateQueue for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) CreateQueueWithContext(ctx aws.Context, input *CreateQueueInput, opts ...request.Option) (*CreateQueueOutput, error) {
	req, out := c.CreateQueueRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteMessage = "DeleteMessage"

// DeleteMessageRequest generates a "aws/request.Request" representing the
// client's request for the DeleteMessage operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DeleteMessage for more information on using the DeleteMessage
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DeleteMessageRequest method.
//    req, resp := client.DeleteMessageRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/DeleteMessage
func (c *SQSExtended) DeleteMessageRequest(input *DeleteMessageInput) (req *request.Request, output *DeleteMessageOutput) {
	op := &request.Operation{
		Name:       opDeleteMessage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteMessageInput{}
	}

	output = &DeleteMessageOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(query.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

// DeleteMessage API operation for Amazon Simple Queue Service.
//
// Deletes the specified message from the specified queue. To select the message
// to delete, use the ReceiptHandle of the message (not the MessageId which
// you receive when you send the message). Amazon SQS can delete a message from
// a queue even if a visibility timeout setting causes the message to be locked
// by another consumer. Amazon SQS automatically deletes messages left in a
// queue longer than the retention period configured for the queue.
//
// The ReceiptHandle is associated with a specific instance of receiving a message.
// If you receive a message more than once, the ReceiptHandle is different each
// time you receive a message. When you use the DeleteMessage action, you must
// provide the most recently received ReceiptHandle for the message (otherwise,
// the request succeeds, but the message might not be deleted).
//
// For standard queues, it is possible to receive a message even after you delete
// it. This might happen on rare occasions if one of the servers which stores
// a copy of the message is unavailable when you send the request to delete
// the message. The copy remains on the server and might be returned to you
// during a subsequent receive request. You should ensure that your application
// is idempotent, so that receiving a message more than once does not cause
// issues.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation DeleteMessage for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInvalidIdFormat "InvalidIdFormat"
//   The specified receipt handle isn't valid for the current version.
//
//   * ErrCodeReceiptHandleIsInvalid "ReceiptHandleIsInvalid"
//   The specified receipt handle isn't valid.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/DeleteMessage
func (c *SQSExtended) DeleteMessage(input *DeleteMessageInput) (*DeleteMessageOutput, error) {
	req, out := c.DeleteMessageRequest(input)
	return out, req.Send()
}

// DeleteMessageWithContext is the same as DeleteMessage with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteMessage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) DeleteMessageWithContext(ctx aws.Context, input *DeleteMessageInput, opts ...request.Option) (*DeleteMessageOutput, error) {
	req, out := c.DeleteMessageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteMessageBatch = "DeleteMessageBatch"

// DeleteMessageBatchRequest generates a "aws/request.Request" representing the
// client's request for the DeleteMessageBatch operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DeleteMessageBatch for more information on using the DeleteMessageBatch
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DeleteMessageBatchRequest method.
//    req, resp := client.DeleteMessageBatchRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/DeleteMessageBatch
func (c *SQSExtended) DeleteMessageBatchRequest(input *DeleteMessageBatchInput) (req *request.Request, output *DeleteMessageBatchOutput) {
	op := &request.Operation{
		Name:       opDeleteMessageBatch,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteMessageBatchInput{}
	}

	output = &DeleteMessageBatchOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DeleteMessageBatch API operation for Amazon Simple Queue Service.
//
// Deletes up to ten messages from the specified queue. This is a batch version
// of DeleteMessage. The result of the action on each message is reported individually
// in the response.
//
// Because the batch request can result in a combination of successful and unsuccessful
// actions, you should check for batch errors even when the call returns an
// HTTP status code of 200.
//
// Some actions take lists of parameters. These lists are specified using the
// param.n notation. Values of n are integers starting from 1. For example,
// a parameter list with two elements looks like this:
//
// &Attribute.1=first
//
// &Attribute.2=second
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation DeleteMessageBatch for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeTooManyEntriesInBatchRequest "AWS.SimpleQueueService.TooManyEntriesInBatchRequest"
//   The batch request contains more entries than permissible.
//
//   * ErrCodeEmptyBatchRequest "AWS.SimpleQueueService.EmptyBatchRequest"
//   The batch request doesn't contain any entries.
//
//   * ErrCodeBatchEntryIdsNotDistinct "AWS.SimpleQueueService.BatchEntryIdsNotDistinct"
//   Two or more batch entries in the request have the same Id.
//
//   * ErrCodeInvalidBatchEntryId "AWS.SimpleQueueService.InvalidBatchEntryId"
//   The Id of a batch entry in a batch request doesn't abide by the specification.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/DeleteMessageBatch
func (c *SQSExtended) DeleteMessageBatch(input *DeleteMessageBatchInput) (*DeleteMessageBatchOutput, error) {
	req, out := c.DeleteMessageBatchRequest(input)
	return out, req.Send()
}

// DeleteMessageBatchWithContext is the same as DeleteMessageBatch with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteMessageBatch for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) DeleteMessageBatchWithContext(ctx aws.Context, input *DeleteMessageBatchInput, opts ...request.Option) (*DeleteMessageBatchOutput, error) {
	req, out := c.DeleteMessageBatchRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteQueue = "DeleteQueue"

// DeleteQueueRequest generates a "aws/request.Request" representing the
// client's request for the DeleteQueue operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DeleteQueue for more information on using the DeleteQueue
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DeleteQueueRequest method.
//    req, resp := client.DeleteQueueRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/DeleteQueue
func (c *SQSExtended) DeleteQueueRequest(input *DeleteQueueInput) (req *request.Request, output *DeleteQueueOutput) {
	op := &request.Operation{
		Name:       opDeleteQueue,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteQueueInput{}
	}

	output = &DeleteQueueOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(query.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

// DeleteQueue API operation for Amazon Simple Queue Service.
//
// Deletes the queue specified by the QueueUrl, regardless of the queue's contents.
// If the specified queue doesn't exist, Amazon SQS returns a successful response.
//
// Be careful with the DeleteQueue action: When you delete a queue, any messages
// in the queue are no longer available.
//
// When you delete a queue, the deletion process takes up to 60 seconds. Requests
// you send involving that queue during the 60 seconds might succeed. For example,
// a SendMessage request might succeed, but after 60 seconds the queue and the
// message you sent no longer exist.
//
// When you delete a queue, you must wait at least 60 seconds before creating
// a queue with the same name.
//
// Cross-account permissions don't apply to this action. For more information,
// see Grant Cross-Account Permissions to a Role and a User Name (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon Simple Queue Service Developer Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation DeleteQueue for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/DeleteQueue
func (c *SQSExtended) DeleteQueue(input *DeleteQueueInput) (*DeleteQueueOutput, error) {
	req, out := c.DeleteQueueRequest(input)
	return out, req.Send()
}

// DeleteQueueWithContext is the same as DeleteQueue with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteQueue for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) DeleteQueueWithContext(ctx aws.Context, input *DeleteQueueInput, opts ...request.Option) (*DeleteQueueOutput, error) {
	req, out := c.DeleteQueueRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetQueueAttributes = "GetQueueAttributes"

// GetQueueAttributesRequest generates a "aws/request.Request" representing the
// client's request for the GetQueueAttributes operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetQueueAttributes for more information on using the GetQueueAttributes
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetQueueAttributesRequest method.
//    req, resp := client.GetQueueAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/GetQueueAttributes
func (c *SQSExtended) GetQueueAttributesRequest(input *GetQueueAttributesInput) (req *request.Request, output *GetQueueAttributesOutput) {
	op := &request.Operation{
		Name:       opGetQueueAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetQueueAttributesInput{}
	}

	output = &GetQueueAttributesOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetQueueAttributes API operation for Amazon Simple Queue Service.
//
// Gets attributes for the specified queue.
//
// To determine whether a queue is FIFO (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html),
// you can check whether QueueName ends with the .fifo suffix.
//
// Some actions take lists of parameters. These lists are specified using the
// param.n notation. Values of n are integers starting from 1. For example,
// a parameter list with two elements looks like this:
//
// &Attribute.1=first
//
// &Attribute.2=second
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation GetQueueAttributes for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInvalidAttributeName "InvalidAttributeName"
//   The specified attribute doesn't exist.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/GetQueueAttributes
func (c *SQSExtended) GetQueueAttributes(input *GetQueueAttributesInput) (*GetQueueAttributesOutput, error) {
	req, out := c.GetQueueAttributesRequest(input)
	return out, req.Send()
}

// GetQueueAttributesWithContext is the same as GetQueueAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See GetQueueAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) GetQueueAttributesWithContext(ctx aws.Context, input *GetQueueAttributesInput, opts ...request.Option) (*GetQueueAttributesOutput, error) {
	req, out := c.GetQueueAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetQueueUrl = "GetQueueUrl"

// GetQueueUrlRequest generates a "aws/request.Request" representing the
// client's request for the GetQueueUrl operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetQueueUrl for more information on using the GetQueueUrl
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetQueueUrlRequest method.
//    req, resp := client.GetQueueUrlRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/GetQueueUrl
func (c *SQSExtended) GetQueueUrlRequest(input *GetQueueUrlInput) (req *request.Request, output *GetQueueUrlOutput) {
	op := &request.Operation{
		Name:       opGetQueueUrl,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &GetQueueUrlInput{}
	}

	output = &GetQueueUrlOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetQueueUrl API operation for Amazon Simple Queue Service.
//
// Returns the URL of an existing Amazon SQS queue.
//
// To access a queue that belongs to another AWS account, use the QueueOwnerAWSAccountId
// parameter to specify the account ID of the queue's owner. The queue's owner
// must grant you permission to access the queue. For more information about
// shared queue access, see AddPermission or see Allow Developers to Write Messages
// to a Shared Queue (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-writing-an-sqs-policy.html#write-messages-to-shared-queue)
// in the Amazon Simple Queue Service Developer Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation GetQueueUrl for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeQueueDoesNotExist "AWS.SimpleQueueService.NonExistentQueue"
//   The specified queue doesn't exist.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/GetQueueUrl
func (c *SQSExtended) GetQueueUrl(input *GetQueueUrlInput) (*GetQueueUrlOutput, error) {
	req, out := c.GetQueueUrlRequest(input)
	return out, req.Send()
}

// GetQueueUrlWithContext is the same as GetQueueUrl with the addition of
// the ability to pass a context and additional request options.
//
// See GetQueueUrl for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) GetQueueUrlWithContext(ctx aws.Context, input *GetQueueUrlInput, opts ...request.Option) (*GetQueueUrlOutput, error) {
	req, out := c.GetQueueUrlRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListDeadLetterSourceQueues = "ListDeadLetterSourceQueues"

// ListDeadLetterSourceQueuesRequest generates a "aws/request.Request" representing the
// client's request for the ListDeadLetterSourceQueues operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ListDeadLetterSourceQueues for more information on using the ListDeadLetterSourceQueues
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ListDeadLetterSourceQueuesRequest method.
//    req, resp := client.ListDeadLetterSourceQueuesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ListDeadLetterSourceQueues
func (c *SQSExtended) ListDeadLetterSourceQueuesRequest(input *ListDeadLetterSourceQueuesInput) (req *request.Request, output *ListDeadLetterSourceQueuesOutput) {
	op := &request.Operation{
		Name:       opListDeadLetterSourceQueues,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListDeadLetterSourceQueuesInput{}
	}

	output = &ListDeadLetterSourceQueuesOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ListDeadLetterSourceQueues API operation for Amazon Simple Queue Service.
//
// Returns a list of your queues that have the RedrivePolicy queue attribute
// configured with a dead-letter queue.
//
// For more information about using dead-letter queues, see Using Amazon SQS
// Dead-Letter Queues (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html)
// in the Amazon Simple Queue Service Developer Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation ListDeadLetterSourceQueues for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeQueueDoesNotExist "AWS.SimpleQueueService.NonExistentQueue"
//   The specified queue doesn't exist.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ListDeadLetterSourceQueues
func (c *SQSExtended) ListDeadLetterSourceQueues(input *ListDeadLetterSourceQueuesInput) (*ListDeadLetterSourceQueuesOutput, error) {
	req, out := c.ListDeadLetterSourceQueuesRequest(input)
	return out, req.Send()
}

// ListDeadLetterSourceQueuesWithContext is the same as ListDeadLetterSourceQueues with the addition of
// the ability to pass a context and additional request options.
//
// See ListDeadLetterSourceQueues for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) ListDeadLetterSourceQueuesWithContext(ctx aws.Context, input *ListDeadLetterSourceQueuesInput, opts ...request.Option) (*ListDeadLetterSourceQueuesOutput, error) {
	req, out := c.ListDeadLetterSourceQueuesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListQueueTags = "ListQueueTags"

// ListQueueTagsRequest generates a "aws/request.Request" representing the
// client's request for the ListQueueTags operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ListQueueTags for more information on using the ListQueueTags
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ListQueueTagsRequest method.
//    req, resp := client.ListQueueTagsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ListQueueTags
func (c *SQSExtended) ListQueueTagsRequest(input *ListQueueTagsInput) (req *request.Request, output *ListQueueTagsOutput) {
	op := &request.Operation{
		Name:       opListQueueTags,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListQueueTagsInput{}
	}

	output = &ListQueueTagsOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ListQueueTags API operation for Amazon Simple Queue Service.
//
// List all cost allocation tags added to the specified Amazon SQS queue. For
// an overview, see Tagging Your Amazon SQS Queues (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-tags.html)
// in the Amazon Simple Queue Service Developer Guide.
//
// Cross-account permissions don't apply to this action. For more information,
// see Grant Cross-Account Permissions to a Role and a User Name (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon Simple Queue Service Developer Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation ListQueueTags for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ListQueueTags
func (c *SQSExtended) ListQueueTags(input *ListQueueTagsInput) (*ListQueueTagsOutput, error) {
	req, out := c.ListQueueTagsRequest(input)
	return out, req.Send()
}

// ListQueueTagsWithContext is the same as ListQueueTags with the addition of
// the ability to pass a context and additional request options.
//
// See ListQueueTags for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) ListQueueTagsWithContext(ctx aws.Context, input *ListQueueTagsInput, opts ...request.Option) (*ListQueueTagsOutput, error) {
	req, out := c.ListQueueTagsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListQueues = "ListQueues"

// ListQueuesRequest generates a "aws/request.Request" representing the
// client's request for the ListQueues operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ListQueues for more information on using the ListQueues
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ListQueuesRequest method.
//    req, resp := client.ListQueuesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ListQueues
func (c *SQSExtended) ListQueuesRequest(input *ListQueuesInput) (req *request.Request, output *ListQueuesOutput) {
	op := &request.Operation{
		Name:       opListQueues,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ListQueuesInput{}
	}

	output = &ListQueuesOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ListQueues API operation for Amazon Simple Queue Service.
//
// Returns a list of your queues. The maximum number of queues that can be returned
// is 1,000. If you specify a value for the optional QueueNamePrefix parameter,
// only queues with a name that begins with the specified value are returned.
//
// Cross-account permissions don't apply to this action. For more information,
// see Grant Cross-Account Permissions to a Role and a User Name (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon Simple Queue Service Developer Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation ListQueues for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ListQueues
func (c *SQSExtended) ListQueues(input *ListQueuesInput) (*ListQueuesOutput, error) {
	req, out := c.ListQueuesRequest(input)
	return out, req.Send()
}

// ListQueuesWithContext is the same as ListQueues with the addition of
// the ability to pass a context and additional request options.
//
// See ListQueues for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) ListQueuesWithContext(ctx aws.Context, input *ListQueuesInput, opts ...request.Option) (*ListQueuesOutput, error) {
	req, out := c.ListQueuesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opPurgeQueue = "PurgeQueue"

// PurgeQueueRequest generates a "aws/request.Request" representing the
// client's request for the PurgeQueue operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See PurgeQueue for more information on using the PurgeQueue
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the PurgeQueueRequest method.
//    req, resp := client.PurgeQueueRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/PurgeQueue
func (c *SQSExtended) PurgeQueueRequest(input *PurgeQueueInput) (req *request.Request, output *PurgeQueueOutput) {
	op := &request.Operation{
		Name:       opPurgeQueue,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PurgeQueueInput{}
	}

	output = &PurgeQueueOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(query.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

// PurgeQueue API operation for Amazon Simple Queue Service.
//
// Deletes the messages in a queue specified by the QueueURL parameter.
//
// When you use the PurgeQueue action, you can't retrieve any messages deleted
// from a queue.
//
// The message deletion process takes up to 60 seconds. We recommend waiting
// for 60 seconds regardless of your queue's size.
//
// Messages sent to the queue before you call PurgeQueue might be received but
// are deleted within the next minute.
//
// Messages sent to the queue after you call PurgeQueue might be deleted while
// the queue is being purged.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation PurgeQueue for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeQueueDoesNotExist "AWS.SimpleQueueService.NonExistentQueue"
//   The specified queue doesn't exist.
//
//   * ErrCodePurgeQueueInProgress "AWS.SimpleQueueService.PurgeQueueInProgress"
//   Indicates that the specified queue previously received a PurgeQueue request
//   within the last 60 seconds (the time it can take to delete the messages in
//   the queue).
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/PurgeQueue
func (c *SQSExtended) PurgeQueue(input *PurgeQueueInput) (*PurgeQueueOutput, error) {
	req, out := c.PurgeQueueRequest(input)
	return out, req.Send()
}

// PurgeQueueWithContext is the same as PurgeQueue with the addition of
// the ability to pass a context and additional request options.
//
// See PurgeQueue for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) PurgeQueueWithContext(ctx aws.Context, input *PurgeQueueInput, opts ...request.Option) (*PurgeQueueOutput, error) {
	req, out := c.PurgeQueueRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opReceiveMessage = "ReceiveMessage"

// ReceiveMessageRequest generates a "aws/request.Request" representing the
// client's request for the ReceiveMessage operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ReceiveMessage for more information on using the ReceiveMessage
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ReceiveMessageRequest method.
//    req, resp := client.ReceiveMessageRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ReceiveMessage
func (c *SQSExtended) ReceiveMessageRequest(input *ReceiveMessageInput) (req *request.Request, output *ReceiveMessageOutput) {
	op := &request.Operation{
		Name:       opReceiveMessage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ReceiveMessageInput{}
	}

	output = &ReceiveMessageOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ReceiveMessage API operation for Amazon Simple Queue Service.
//
// Retrieves one or more messages (up to 10), from the specified queue. Using
// the WaitTimeSeconds parameter enables long-poll support. For more information,
// see Amazon SQS Long Polling (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-long-polling.html)
// in the Amazon Simple Queue Service Developer Guide.
//
// Short poll is the default behavior where a weighted random set of machines
// is sampled on a ReceiveMessage call. Thus, only the messages on the sampled
// machines are returned. If the number of messages in the queue is small (fewer
// than 1,000), you most likely get fewer messages than you requested per ReceiveMessage
// call. If the number of messages in the queue is extremely small, you might
// not receive any messages in a particular ReceiveMessage response. If this
// happens, repeat the request.
//
// For each message returned, the response includes the following:
//
//    * The message body.
//
//    * An MD5 digest of the message body. For information about MD5, see RFC1321
//    (https://www.ietf.org/rfc/rfc1321.txt).
//
//    * The MessageId you received when you sent the message to the queue.
//
//    * The receipt handle.
//
//    * The message attributes.
//
//    * An MD5 digest of the message attributes.
//
// The receipt handle is the identifier you must provide when deleting the message.
// For more information, see Queue and Message Identifiers (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-message-identifiers.html)
// in the Amazon Simple Queue Service Developer Guide.
//
// You can provide the VisibilityTimeout parameter in your request. The parameter
// is applied to the messages that Amazon SQS returns in the response. If you
// don't include the parameter, the overall visibility timeout for the queue
// is used for the returned messages. For more information, see Visibility Timeout
// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html)
// in the Amazon Simple Queue Service Developer Guide.
//
// A message that isn't deleted or a message whose visibility isn't extended
// before the visibility timeout expires counts as a failed receive. Depending
// on the configuration of the queue, the message might be sent to the dead-letter
// queue.
//
// In the future, new attributes might be added. If you write code that calls
// this action, we recommend that you structure your code so that it can handle
// new attributes gracefully.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation ReceiveMessage for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeOverLimit "OverLimit"
//   The specified action violates a limit. For example, ReceiveMessage returns
//   this error if the maximum number of inflight messages is reached and AddPermission
//   returns this error if the maximum number of permissions for the queue is
//   reached.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/ReceiveMessage
func (c *SQSExtended) ReceiveMessage(input *ReceiveMessageInput) (*ReceiveMessageOutput, error) {
	req, out := c.ReceiveMessageRequest(input)
	return out, req.Send()
}

// ReceiveMessageWithContext is the same as ReceiveMessage with the addition of
// the ability to pass a context and additional request options.
//
// See ReceiveMessage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) ReceiveMessageWithContext(ctx aws.Context, input *ReceiveMessageInput, opts ...request.Option) (*ReceiveMessageOutput, error) {
	req, out := c.ReceiveMessageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opRemovePermission = "RemovePermission"

// RemovePermissionRequest generates a "aws/request.Request" representing the
// client's request for the RemovePermission operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See RemovePermission for more information on using the RemovePermission
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the RemovePermissionRequest method.
//    req, resp := client.RemovePermissionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/RemovePermission
func (c *SQSExtended) RemovePermissionRequest(input *RemovePermissionInput) (req *request.Request, output *RemovePermissionOutput) {
	op := &request.Operation{
		Name:       opRemovePermission,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &RemovePermissionInput{}
	}

	output = &RemovePermissionOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(query.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

// RemovePermission API operation for Amazon Simple Queue Service.
//
// Revokes any permissions in the queue policy that matches the specified Label
// parameter.
//
//    * Only the owner of a queue can remove permissions from it.
//
//    * Cross-account permissions don't apply to this action. For more information,
//    see Grant Cross-Account Permissions to a Role and a User Name (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
//    in the Amazon Simple Queue Service Developer Guide.
//
//    * To remove the ability to change queue permissions, you must deny permission
//    to the AddPermission, RemovePermission, and SetQueueAttributes actions
//    in your IAM policy.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation RemovePermission for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/RemovePermission
func (c *SQSExtended) RemovePermission(input *RemovePermissionInput) (*RemovePermissionOutput, error) {
	req, out := c.RemovePermissionRequest(input)
	return out, req.Send()
}

// RemovePermissionWithContext is the same as RemovePermission with the addition of
// the ability to pass a context and additional request options.
//
// See RemovePermission for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) RemovePermissionWithContext(ctx aws.Context, input *RemovePermissionInput, opts ...request.Option) (*RemovePermissionOutput, error) {
	req, out := c.RemovePermissionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opSendMessage = "SendMessage"

// SendMessageRequest generates a "aws/request.Request" representing the
// client's request for the SendMessage operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See SendMessage for more information on using the SendMessage
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the SendMessageRequest method.
//    req, resp := client.SendMessageRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/SendMessage
func (c *SQSExtended) SendMessageRequest(input *SendMessageInput) (req *request.Request, output *SendMessageOutput) {
	op := &request.Operation{
		Name:       opSendMessage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SendMessageInput{}
	}

	output = &SendMessageOutput{}
	req = c.newRequest(op, input, output)
	return
}

// SendMessage API operation for Amazon Simple Queue Service.
//
// Delivers a message to the specified queue.
//
// A message can include only XML, JSON, and unformatted text. The following
// Unicode characters are allowed:
//
// #x9 | #xA | #xD | #x20 to #xD7FF | #xE000 to #xFFFD | #x10000 to #x10FFFF
//
// Any characters not included in this list will be rejected. For more information,
// see the W3C specification for characters (http://www.w3.org/TR/REC-xml/#charsets).
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation SendMessage for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInvalidMessageContents "InvalidMessageContents"
//   The message contains characters outside the allowed set.
//
//   * ErrCodeUnsupportedOperation "AWS.SimpleQueueService.UnsupportedOperation"
//   Error code 400. Unsupported operation.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/SendMessage
func (c *SQSExtended) SendMessage(input *SendMessageInput) (*SendMessageOutput, error) {
	req, out := c.SendMessageRequest(input)
	return out, req.Send()
}

// SendMessageWithContext is the same as SendMessage with the addition of
// the ability to pass a context and additional request options.
//
// See SendMessage for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) SendMessageWithContext(ctx aws.Context, input *SendMessageInput, opts ...request.Option) (*SendMessageOutput, error) {
	req, out := c.SendMessageRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opSendMessageBatch = "SendMessageBatch"

// SendMessageBatchRequest generates a "aws/request.Request" representing the
// client's request for the SendMessageBatch operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See SendMessageBatch for more information on using the SendMessageBatch
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the SendMessageBatchRequest method.
//    req, resp := client.SendMessageBatchRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/SendMessageBatch
func (c *SQSExtended) SendMessageBatchRequest(input *SendMessageBatchInput) (req *request.Request, output *SendMessageBatchOutput) {
	op := &request.Operation{
		Name:       opSendMessageBatch,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SendMessageBatchInput{}
	}

	output = &SendMessageBatchOutput{}
	req = c.newRequest(op, input, output)
	return
}

// SendMessageBatch API operation for Amazon Simple Queue Service.
//
// Delivers up to ten messages to the specified queue. This is a batch version
// of SendMessage. For a FIFO queue, multiple messages within a single batch
// are enqueued in the order they are sent.
//
// The result of sending each message is reported individually in the response.
// Because the batch request can result in a combination of successful and unsuccessful
// actions, you should check for batch errors even when the call returns an
// HTTP status code of 200.
//
// The maximum allowed individual message size and the maximum total payload
// size (the sum of the individual lengths of all of the batched messages) are
// both 256 KB (262,144 bytes).
//
// A message can include only XML, JSON, and unformatted text. The following
// Unicode characters are allowed:
//
// #x9 | #xA | #xD | #x20 to #xD7FF | #xE000 to #xFFFD | #x10000 to #x10FFFF
//
// Any characters not included in this list will be rejected. For more information,
// see the W3C specification for characters (http://www.w3.org/TR/REC-xml/#charsets).
//
// If you don't specify the DelaySeconds parameter for an entry, Amazon SQS
// uses the default value for the queue.
//
// Some actions take lists of parameters. These lists are specified using the
// param.n notation. Values of n are integers starting from 1. For example,
// a parameter list with two elements looks like this:
//
// &Attribute.1=first
//
// &Attribute.2=second
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation SendMessageBatch for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeTooManyEntriesInBatchRequest "AWS.SimpleQueueService.TooManyEntriesInBatchRequest"
//   The batch request contains more entries than permissible.
//
//   * ErrCodeEmptyBatchRequest "AWS.SimpleQueueService.EmptyBatchRequest"
//   The batch request doesn't contain any entries.
//
//   * ErrCodeBatchEntryIdsNotDistinct "AWS.SimpleQueueService.BatchEntryIdsNotDistinct"
//   Two or more batch entries in the request have the same Id.
//
//   * ErrCodeBatchRequestTooLong "AWS.SimpleQueueService.BatchRequestTooLong"
//   The length of all the messages put together is more than the limit.
//
//   * ErrCodeInvalidBatchEntryId "AWS.SimpleQueueService.InvalidBatchEntryId"
//   The Id of a batch entry in a batch request doesn't abide by the specification.
//
//   * ErrCodeUnsupportedOperation "AWS.SimpleQueueService.UnsupportedOperation"
//   Error code 400. Unsupported operation.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/SendMessageBatch
func (c *SQSExtended) SendMessageBatch(input *SendMessageBatchInput) (*SendMessageBatchOutput, error) {
	req, out := c.SendMessageBatchRequest(input)
	return out, req.Send()
}

// SendMessageBatchWithContext is the same as SendMessageBatch with the addition of
// the ability to pass a context and additional request options.
//
// See SendMessageBatch for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) SendMessageBatchWithContext(ctx aws.Context, input *SendMessageBatchInput, opts ...request.Option) (*SendMessageBatchOutput, error) {
	req, out := c.SendMessageBatchRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opSetQueueAttributes = "SetQueueAttributes"

// SetQueueAttributesRequest generates a "aws/request.Request" representing the
// client's request for the SetQueueAttributes operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See SetQueueAttributes for more information on using the SetQueueAttributes
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the SetQueueAttributesRequest method.
//    req, resp := client.SetQueueAttributesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/SetQueueAttributes
func (c *SQSExtended) SetQueueAttributesRequest(input *SetQueueAttributesInput) (req *request.Request, output *SetQueueAttributesOutput) {
	op := &request.Operation{
		Name:       opSetQueueAttributes,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetQueueAttributesInput{}
	}

	output = &SetQueueAttributesOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(query.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

// SetQueueAttributes API operation for Amazon Simple Queue Service.
//
// Sets the value of one or more queue attributes. When you change a queue's
// attributes, the change can take up to 60 seconds for most of the attributes
// to propagate throughout the Amazon SQS system. Changes made to the MessageRetentionPeriod
// attribute can take up to 15 minutes.
//
//    * In the future, new attributes might be added. If you write code that
//    calls this action, we recommend that you structure your code so that it
//    can handle new attributes gracefully.
//
//    * Cross-account permissions don't apply to this action. For more information,
//    see Grant Cross-Account Permissions to a Role and a User Name (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
//    in the Amazon Simple Queue Service Developer Guide.
//
//    * To remove the ability to change queue permissions, you must deny permission
//    to the AddPermission, RemovePermission, and SetQueueAttributes actions
//    in your IAM policy.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation SetQueueAttributes for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInvalidAttributeName "InvalidAttributeName"
//   The specified attribute doesn't exist.
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/SetQueueAttributes
func (c *SQSExtended) SetQueueAttributes(input *SetQueueAttributesInput) (*SetQueueAttributesOutput, error) {
	req, out := c.SetQueueAttributesRequest(input)
	return out, req.Send()
}

// SetQueueAttributesWithContext is the same as SetQueueAttributes with the addition of
// the ability to pass a context and additional request options.
//
// See SetQueueAttributes for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) SetQueueAttributesWithContext(ctx aws.Context, input *SetQueueAttributesInput, opts ...request.Option) (*SetQueueAttributesOutput, error) {
	req, out := c.SetQueueAttributesRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opTagQueue = "TagQueue"

// TagQueueRequest generates a "aws/request.Request" representing the
// client's request for the TagQueue operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See TagQueue for more information on using the TagQueue
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the TagQueueRequest method.
//    req, resp := client.TagQueueRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/TagQueue
func (c *SQSExtended) TagQueueRequest(input *TagQueueInput) (req *request.Request, output *TagQueueOutput) {
	op := &request.Operation{
		Name:       opTagQueue,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &TagQueueInput{}
	}

	output = &TagQueueOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(query.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

// TagQueue API operation for Amazon Simple Queue Service.
//
// Add cost allocation tags to the specified Amazon SQS queue. For an overview,
// see Tagging Your Amazon SQS Queues (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-tags.html)
// in the Amazon Simple Queue Service Developer Guide.
//
// When you use queue tags, keep the following guidelines in mind:
//
//    * Adding more than 50 tags to a queue isn't recommended.
//
//    * Tags don't have any semantic meaning. Amazon SQS interprets tags as
//    character strings.
//
//    * Tags are case-sensitive.
//
//    * A new tag with a key identical to that of an existing tag overwrites
//    the existing tag.
//
// For a full list of tag restrictions, see Limits Related to Queues (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-limits.html#limits-queues)
// in the Amazon Simple Queue Service Developer Guide.
//
// Cross-account permissions don't apply to this action. For more information,
// see Grant Cross-Account Permissions to a Role and a User Name (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon Simple Queue Service Developer Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation TagQueue for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/TagQueue
func (c *SQSExtended) TagQueue(input *TagQueueInput) (*TagQueueOutput, error) {
	req, out := c.TagQueueRequest(input)
	return out, req.Send()
}

// TagQueueWithContext is the same as TagQueue with the addition of
// the ability to pass a context and additional request options.
//
// See TagQueue for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) TagQueueWithContext(ctx aws.Context, input *TagQueueInput, opts ...request.Option) (*TagQueueOutput, error) {
	req, out := c.TagQueueRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opUntagQueue = "UntagQueue"

// UntagQueueRequest generates a "aws/request.Request" representing the
// client's request for the UntagQueue operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See UntagQueue for more information on using the UntagQueue
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the UntagQueueRequest method.
//    req, resp := client.UntagQueueRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/UntagQueue
func (c *SQSExtended) UntagQueueRequest(input *UntagQueueInput) (req *request.Request, output *UntagQueueOutput) {
	op := &request.Operation{
		Name:       opUntagQueue,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UntagQueueInput{}
	}

	output = &UntagQueueOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Swap(query.UnmarshalHandler.Name, protocol.UnmarshalDiscardBodyHandler)
	return
}

// UntagQueue API operation for Amazon Simple Queue Service.
//
// Remove cost allocation tags from the specified Amazon SQS queue. For an overview,
// see Tagging Your Amazon SQS Queues (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-tags.html)
// in the Amazon Simple Queue Service Developer Guide.
//
// Cross-account permissions don't apply to this action. For more information,
// see Grant Cross-Account Permissions to a Role and a User Name (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
// in the Amazon Simple Queue Service Developer Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Simple Queue Service's
// API operation UntagQueue for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/sqs-2012-11-05/UntagQueue
func (c *SQSExtended) UntagQueue(input *UntagQueueInput) (*UntagQueueOutput, error) {
	req, out := c.UntagQueueRequest(input)
	return out, req.Send()
}

// UntagQueueWithContext is the same as UntagQueue with the addition of
// the ability to pass a context and additional request options.
//
// See UntagQueue for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *SQSExtended) UntagQueueWithContext(ctx aws.Context, input *UntagQueueInput, opts ...request.Option) (*UntagQueueOutput, error) {
	req, out := c.UntagQueueRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AddPermissionInput struct {
	_ struct{} `type:"structure"`

	// The AWS account number of the principal (https://docs.aws.amazon.com/general/latest/gr/glos-chap.html#P)
	// who is given permission. The principal must have an AWS account, but does
	// not need to be signed up for Amazon SQS. For information about locating the
	// AWS account identification, see Your AWS Identifiers (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-making-api-requests.html#sqs-api-request-authentication)
	// in the Amazon Simple Queue Service Developer Guide.
	//
	// AWSAccountIds is a required field
	AWSAccountIds []*string `locationNameList:"AWSAccountId" type:"list" flattened:"true" required:"true"`

	// The action the client wants to allow for the specified principal. Valid values:
	// the name of any action or *.
	//
	// For more information about these actions, see Overview of Managing Access
	// Permissions to Your Amazon Simple Queue Service Resource (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-overview-of-managing-access.html)
	// in the Amazon Simple Queue Service Developer Guide.
	//
	// Specifying SendMessage, DeleteMessage, or ChangeMessageVisibility for ActionName.n
	// also grants permissions for the corresponding batch versions of those actions:
	// SendMessageBatch, DeleteMessageBatch, and ChangeMessageVisibilityBatch.
	//
	// Actions is a required field
	Actions []*string `locationNameList:"ActionName" type:"list" flattened:"true" required:"true"`

	// The unique identification of the permission you're setting (for example,
	// AliceSendMessage). Maximum 80 characters. Allowed characters include alphanumeric
	// characters, hyphens (-), and underscores (_).
	//
	// Label is a required field
	Label *string `type:"string" required:"true"`

	// The URL of the Amazon SQS queue to which permissions are added.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`
}

// String returns the string representation
func (s AddPermissionInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AddPermissionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddPermissionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "AddPermissionInput"}
	if s.AWSAccountIds == nil {
		invalidParams.Add(request.NewErrParamRequired("AWSAccountIds"))
	}
	if s.Actions == nil {
		invalidParams.Add(request.NewErrParamRequired("Actions"))
	}
	if s.Label == nil {
		invalidParams.Add(request.NewErrParamRequired("Label"))
	}
	if s.QueueUrl == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueUrl"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAWSAccountIds sets the AWSAccountIds field's value.
func (s *AddPermissionInput) SetAWSAccountIds(v []*string) *AddPermissionInput {
	s.AWSAccountIds = v
	return s
}

// SetActions sets the Actions field's value.
func (s *AddPermissionInput) SetActions(v []*string) *AddPermissionInput {
	s.Actions = v
	return s
}

// SetLabel sets the Label field's value.
func (s *AddPermissionInput) SetLabel(v string) *AddPermissionInput {
	s.Label = &v
	return s
}

// SetQueueUrl sets the QueueUrl field's value.
func (s *AddPermissionInput) SetQueueUrl(v string) *AddPermissionInput {
	s.QueueUrl = &v
	return s
}

type AddPermissionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s AddPermissionOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s AddPermissionOutput) GoString() string {
	return s.String()
}

// Gives a detailed description of the result of an action on each entry in
// the request.
type BatchResultErrorEntry struct {
	_ struct{} `type:"structure"`

	// An error code representing why the action failed on this entry.
	//
	// Code is a required field
	Code *string `type:"string" required:"true"`

	// The Id of an entry in a batch request.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`

	// A message explaining why the action failed on this entry.
	Message *string `type:"string"`

	// Specifies whether the error happened due to the producer.
	//
	// SenderFault is a required field
	SenderFault *bool `type:"boolean" required:"true"`
}

// String returns the string representation
func (s BatchResultErrorEntry) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s BatchResultErrorEntry) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *BatchResultErrorEntry) SetCode(v string) *BatchResultErrorEntry {
	s.Code = &v
	return s
}

// SetId sets the Id field's value.
func (s *BatchResultErrorEntry) SetId(v string) *BatchResultErrorEntry {
	s.Id = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *BatchResultErrorEntry) SetMessage(v string) *BatchResultErrorEntry {
	s.Message = &v
	return s
}

// SetSenderFault sets the SenderFault field's value.
func (s *BatchResultErrorEntry) SetSenderFault(v bool) *BatchResultErrorEntry {
	s.SenderFault = &v
	return s
}

type ChangeMessageVisibilityBatchInput struct {
	_ struct{} `type:"structure"`

	// A list of receipt handles of the messages for which the visibility timeout
	// must be changed.
	//
	// Entries is a required field
	Entries []*ChangeMessageVisibilityBatchRequestEntry `locationNameList:"ChangeMessageVisibilityBatchRequestEntry" type:"list" flattened:"true" required:"true"`

	// The URL of the Amazon SQS queue whose messages' visibility is changed.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ChangeMessageVisibilityBatchInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ChangeMessageVisibilityBatchInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ChangeMessageVisibilityBatchInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ChangeMessageVisibilityBatchInput"}
	if s.Entries == nil {
		invalidParams.Add(request.NewErrParamRequired("Entries"))
	}
	if s.QueueUrl == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueUrl"))
	}
	if s.Entries != nil {
		for i, v := range s.Entries {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Entries", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEntries sets the Entries field's value.
func (s *ChangeMessageVisibilityBatchInput) SetEntries(v []*ChangeMessageVisibilityBatchRequestEntry) *ChangeMessageVisibilityBatchInput {
	s.Entries = v
	return s
}

// SetQueueUrl sets the QueueUrl field's value.
func (s *ChangeMessageVisibilityBatchInput) SetQueueUrl(v string) *ChangeMessageVisibilityBatchInput {
	s.QueueUrl = &v
	return s
}

// For each message in the batch, the response contains a ChangeMessageVisibilityBatchResultEntry
// tag if the message succeeds or a BatchResultErrorEntry tag if the message
// fails.
type ChangeMessageVisibilityBatchOutput struct {
	_ struct{} `type:"structure"`

	// A list of BatchResultErrorEntry items.
	//
	// Failed is a required field
	Failed []*BatchResultErrorEntry `locationNameList:"BatchResultErrorEntry" type:"list" flattened:"true" required:"true"`

	// A list of ChangeMessageVisibilityBatchResultEntry items.
	//
	// Successful is a required field
	Successful []*ChangeMessageVisibilityBatchResultEntry `locationNameList:"ChangeMessageVisibilityBatchResultEntry" type:"list" flattened:"true" required:"true"`
}

// String returns the string representation
func (s ChangeMessageVisibilityBatchOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ChangeMessageVisibilityBatchOutput) GoString() string {
	return s.String()
}

// SetFailed sets the Failed field's value.
func (s *ChangeMessageVisibilityBatchOutput) SetFailed(v []*BatchResultErrorEntry) *ChangeMessageVisibilityBatchOutput {
	s.Failed = v
	return s
}

// SetSuccessful sets the Successful field's value.
func (s *ChangeMessageVisibilityBatchOutput) SetSuccessful(v []*ChangeMessageVisibilityBatchResultEntry) *ChangeMessageVisibilityBatchOutput {
	s.Successful = v
	return s
}

// Encloses a receipt handle and an entry id for each message in ChangeMessageVisibilityBatch.
//
// All of the following list parameters must be prefixed with ChangeMessageVisibilityBatchRequestEntry.n,
// where n is an integer value starting with 1. For example, a parameter list
// for this action might look like this:
//
// &ChangeMessageVisibilityBatchRequestEntry.1.Id=change_visibility_msg_2
//
// &ChangeMessageVisibilityBatchRequestEntry.1.ReceiptHandle=your_receipt_handle
//
// &ChangeMessageVisibilityBatchRequestEntry.1.VisibilityTimeout=45
type ChangeMessageVisibilityBatchRequestEntry struct {
	_ struct{} `type:"structure"`

	// An identifier for this particular receipt handle used to communicate the
	// result.
	//
	// The Ids of a batch request need to be unique within a request
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`

	// A receipt handle.
	//
	// ReceiptHandle is a required field
	ReceiptHandle *string `type:"string" required:"true"`

	// The new value (in seconds) for the message's visibility timeout.
	VisibilityTimeout *int64 `type:"integer"`
}

// String returns the string representation
func (s ChangeMessageVisibilityBatchRequestEntry) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ChangeMessageVisibilityBatchRequestEntry) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ChangeMessageVisibilityBatchRequestEntry) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ChangeMessageVisibilityBatchRequestEntry"}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}
	if s.ReceiptHandle == nil {
		invalidParams.Add(request.NewErrParamRequired("ReceiptHandle"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetId sets the Id field's value.
func (s *ChangeMessageVisibilityBatchRequestEntry) SetId(v string) *ChangeMessageVisibilityBatchRequestEntry {
	s.Id = &v
	return s
}

// SetReceiptHandle sets the ReceiptHandle field's value.
func (s *ChangeMessageVisibilityBatchRequestEntry) SetReceiptHandle(v string) *ChangeMessageVisibilityBatchRequestEntry {
	s.ReceiptHandle = &v
	return s
}

// SetVisibilityTimeout sets the VisibilityTimeout field's value.
func (s *ChangeMessageVisibilityBatchRequestEntry) SetVisibilityTimeout(v int64) *ChangeMessageVisibilityBatchRequestEntry {
	s.VisibilityTimeout = &v
	return s
}

// Encloses the Id of an entry in ChangeMessageVisibilityBatch.
type ChangeMessageVisibilityBatchResultEntry struct {
	_ struct{} `type:"structure"`

	// Represents a message whose visibility timeout has been changed successfully.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ChangeMessageVisibilityBatchResultEntry) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ChangeMessageVisibilityBatchResultEntry) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *ChangeMessageVisibilityBatchResultEntry) SetId(v string) *ChangeMessageVisibilityBatchResultEntry {
	s.Id = &v
	return s
}

type ChangeMessageVisibilityInput struct {
	_ struct{} `type:"structure"`

	// The URL of the Amazon SQS queue whose message's visibility is changed.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`

	// The receipt handle associated with the message whose visibility timeout is
	// changed. This parameter is returned by the ReceiveMessage action.
	//
	// ReceiptHandle is a required field
	ReceiptHandle *string `type:"string" required:"true"`

	// The new value for the message's visibility timeout (in seconds). Values values:
	// 0 to 43200. Maximum: 12 hours.
	//
	// VisibilityTimeout is a required field
	VisibilityTimeout *int64 `type:"integer" required:"true"`
}

// String returns the string representation
func (s ChangeMessageVisibilityInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ChangeMessageVisibilityInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ChangeMessageVisibilityInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ChangeMessageVisibilityInput"}
	if s.QueueUrl == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueUrl"))
	}
	if s.ReceiptHandle == nil {
		invalidParams.Add(request.NewErrParamRequired("ReceiptHandle"))
	}
	if s.VisibilityTimeout == nil {
		invalidParams.Add(request.NewErrParamRequired("VisibilityTimeout"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetQueueUrl sets the QueueUrl field's value.
func (s *ChangeMessageVisibilityInput) SetQueueUrl(v string) *ChangeMessageVisibilityInput {
	s.QueueUrl = &v
	return s
}

// SetReceiptHandle sets the ReceiptHandle field's value.
func (s *ChangeMessageVisibilityInput) SetReceiptHandle(v string) *ChangeMessageVisibilityInput {
	s.ReceiptHandle = &v
	return s
}

// SetVisibilityTimeout sets the VisibilityTimeout field's value.
func (s *ChangeMessageVisibilityInput) SetVisibilityTimeout(v int64) *ChangeMessageVisibilityInput {
	s.VisibilityTimeout = &v
	return s
}

type ChangeMessageVisibilityOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ChangeMessageVisibilityOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ChangeMessageVisibilityOutput) GoString() string {
	return s.String()
}

type CreateQueueInput struct {
	_ struct{} `type:"structure"`

	// A map of attributes with their corresponding values.
	//
	// The following lists the names, descriptions, and values of the special request
	// parameters that the CreateQueue action uses:
	//
	//    * DelaySeconds - The length of time, in seconds, for which the delivery
	//    of all messages in the queue is delayed. Valid values: An integer from
	//    0 to 900 seconds (15 minutes). Default: 0.
	//
	//    * MaximumMessageSize - The limit of how many bytes a message can contain
	//    before Amazon SQS rejects it. Valid values: An integer from 1,024 bytes
	//    (1 KiB) to 262,144 bytes (256 KiB). Default: 262,144 (256 KiB).
	//
	//    * MessageRetentionPeriod - The length of time, in seconds, for which Amazon
	//    SQS retains a message. Valid values: An integer from 60 seconds (1 minute)
	//    to 1,209,600 seconds (14 days). Default: 345,600 (4 days).
	//
	//    * Policy - The queue's policy. A valid AWS policy. For more information
	//    about policy structure, see Overview of AWS IAM Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/PoliciesOverview.html)
	//    in the Amazon IAM User Guide.
	//
	//    * ReceiveMessageWaitTimeSeconds - The length of time, in seconds, for
	//    which a ReceiveMessage action waits for a message to arrive. Valid values:
	//    An integer from 0 to 20 (seconds). Default: 0.
	//
	//    * RedrivePolicy - The string that includes the parameters for the dead-letter
	//    queue functionality of the source queue. For more information about the
	//    redrive policy and dead-letter queues, see Using Amazon SQS Dead-Letter
	//    Queues (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html)
	//    in the Amazon Simple Queue Service Developer Guide. deadLetterTargetArn
	//    - The Amazon Resource Name (ARN) of the dead-letter queue to which Amazon
	//    SQS moves messages after the value of maxReceiveCount is exceeded. maxReceiveCount
	//    - The number of times a message is delivered to the source queue before
	//    being moved to the dead-letter queue. When the ReceiveCount for a message
	//    exceeds the maxReceiveCount for a queue, Amazon SQS moves the message
	//    to the dead-letter-queue. The dead-letter queue of a FIFO queue must also
	//    be a FIFO queue. Similarly, the dead-letter queue of a standard queue
	//    must also be a standard queue.
	//
	//    * VisibilityTimeout - The visibility timeout for the queue, in seconds.
	//    Valid values: An integer from 0 to 43,200 (12 hours). Default: 30. For
	//    more information about the visibility timeout, see Visibility Timeout
	//    (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html)
	//    in the Amazon Simple Queue Service Developer Guide.
	//
	// The following attributes apply only to server-side-encryption (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html):
	//
	//    * KmsMasterKeyId - The ID of an AWS-managed customer master key (CMK)
	//    for Amazon SQS or a custom CMK. For more information, see Key Terms (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
	//    While the alias of the AWS-managed CMK for Amazon SQS is always alias/aws/sqs,
	//    the alias of a custom CMK can, for example, be alias/MyAlias . For more
	//    examples, see KeyId (https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters)
	//    in the AWS Key Management Service API Reference.
	//
	//    * KmsDataKeyReusePeriodSeconds - The length of time, in seconds, for which
	//    Amazon SQS can reuse a data key (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#data-keys)
	//    to encrypt or decrypt messages before calling AWS KMS again. An integer
	//    representing seconds, between 60 seconds (1 minute) and 86,400 seconds
	//    (24 hours). Default: 300 (5 minutes). A shorter time period provides better
	//    security but results in more calls to KMS which might incur charges after
	//    Free Tier. For more information, see How Does the Data Key Reuse Period
	//    Work? (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-how-does-the-data-key-reuse-period-work).
	//
	// The following attributes apply only to FIFO (first-in-first-out) queues (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html):
	//
	//    * FifoQueue - Designates a queue as FIFO. Valid values: true, false. If
	//    you don't specify the FifoQueue attribute, Amazon SQS creates a standard
	//    queue. You can provide this attribute only during queue creation. You
	//    can't change it for an existing queue. When you set this attribute, you
	//    must also provide the MessageGroupId for your messages explicitly. For
	//    more information, see FIFO Queue Logic (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-understanding-logic)
	//    in the Amazon Simple Queue Service Developer Guide.
	//
	//    * ContentBasedDeduplication - Enables content-based deduplication. Valid
	//    values: true, false. For more information, see Exactly-Once Processing
	//    (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
	//    in the Amazon Simple Queue Service Developer Guide. Every message must
	//    have a unique MessageDeduplicationId, You may provide a MessageDeduplicationId
	//    explicitly. If you aren't able to provide a MessageDeduplicationId and
	//    you enable ContentBasedDeduplication for your queue, Amazon SQS uses a
	//    SHA-256 hash to generate the MessageDeduplicationId using the body of
	//    the message (but not the attributes of the message). If you don't provide
	//    a MessageDeduplicationId and the queue doesn't have ContentBasedDeduplication
	//    set, the action fails with an error. If the queue has ContentBasedDeduplication
	//    set, your MessageDeduplicationId overrides the generated one. When ContentBasedDeduplication
	//    is in effect, messages with identical content sent within the deduplication
	//    interval are treated as duplicates and only one copy of the message is
	//    delivered. If you send one message with ContentBasedDeduplication enabled
	//    and then another message with a MessageDeduplicationId that is the same
	//    as the one generated for the first MessageDeduplicationId, the two messages
	//    are treated as duplicates and only one copy of the message is delivered.
	Attributes map[string]*string `locationName:"Attribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true"`

	// The name of the new queue. The following limits apply to this name:
	//
	//    * A queue name can have up to 80 characters.
	//
	//    * Valid values: alphanumeric characters, hyphens (-), and underscores
	//    (_).
	//
	//    * A FIFO queue name must end with the .fifo suffix.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueName is a required field
	QueueName *string `type:"string" required:"true"`

	// Add cost allocation tags to the specified Amazon SQS queue. For an overview,
	// see Tagging Your Amazon SQS Queues (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-tags.html)
	// in the Amazon Simple Queue Service Developer Guide.
	//
	// When you use queue tags, keep the following guidelines in mind:
	//
	//    * Adding more than 50 tags to a queue isn't recommended.
	//
	//    * Tags don't have any semantic meaning. Amazon SQS interprets tags as
	//    character strings.
	//
	//    * Tags are case-sensitive.
	//
	//    * A new tag with a key identical to that of an existing tag overwrites
	//    the existing tag.
	//
	// For a full list of tag restrictions, see Limits Related to Queues (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-limits.html#limits-queues)
	// in the Amazon Simple Queue Service Developer Guide.
	//
	// To be able to tag a queue on creation, you must have the sqs:CreateQueue
	// and sqs:TagQueue permissions.
	//
	// Cross-account permissions don't apply to this action. For more information,
	// see Grant Cross-Account Permissions to a Role and a User Name (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-customer-managed-policy-examples.html#grant-cross-account-permissions-to-role-and-user-name)
	// in the Amazon Simple Queue Service Developer Guide.
	Tags map[string]*string `locationName:"Tag" locationNameKey:"Key" locationNameValue:"Value" type:"map" flattened:"true"`
}

// String returns the string representation
func (s CreateQueueInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateQueueInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateQueueInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateQueueInput"}
	if s.QueueName == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAttributes sets the Attributes field's value.
func (s *CreateQueueInput) SetAttributes(v map[string]*string) *CreateQueueInput {
	s.Attributes = v
	return s
}

// SetQueueName sets the QueueName field's value.
func (s *CreateQueueInput) SetQueueName(v string) *CreateQueueInput {
	s.QueueName = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *CreateQueueInput) SetTags(v map[string]*string) *CreateQueueInput {
	s.Tags = v
	return s
}

// Returns the QueueUrl attribute of the created queue.
type CreateQueueOutput struct {
	_ struct{} `type:"structure"`

	// The URL of the created Amazon SQS queue.
	QueueUrl *string `type:"string"`
}

// String returns the string representation
func (s CreateQueueOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateQueueOutput) GoString() string {
	return s.String()
}

// SetQueueUrl sets the QueueUrl field's value.
func (s *CreateQueueOutput) SetQueueUrl(v string) *CreateQueueOutput {
	s.QueueUrl = &v
	return s
}

type DeleteMessageBatchInput struct {
	_ struct{} `type:"structure"`

	// A list of receipt handles for the messages to be deleted.
	//
	// Entries is a required field
	Entries []*DeleteMessageBatchRequestEntry `locationNameList:"DeleteMessageBatchRequestEntry" type:"list" flattened:"true" required:"true"`

	// The URL of the Amazon SQS queue from which messages are deleted.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMessageBatchInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteMessageBatchInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMessageBatchInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteMessageBatchInput"}
	if s.Entries == nil {
		invalidParams.Add(request.NewErrParamRequired("Entries"))
	}
	if s.QueueUrl == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueUrl"))
	}
	if s.Entries != nil {
		for i, v := range s.Entries {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Entries", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEntries sets the Entries field's value.
func (s *DeleteMessageBatchInput) SetEntries(v []*DeleteMessageBatchRequestEntry) *DeleteMessageBatchInput {
	s.Entries = v
	return s
}

// SetQueueUrl sets the QueueUrl field's value.
func (s *DeleteMessageBatchInput) SetQueueUrl(v string) *DeleteMessageBatchInput {
	s.QueueUrl = &v
	return s
}

// For each message in the batch, the response contains a DeleteMessageBatchResultEntry
// tag if the message is deleted or a BatchResultErrorEntry tag if the message
// can't be deleted.
type DeleteMessageBatchOutput struct {
	_ struct{} `type:"structure"`

	// A list of BatchResultErrorEntry items.
	//
	// Failed is a required field
	Failed []*BatchResultErrorEntry `locationNameList:"BatchResultErrorEntry" type:"list" flattened:"true" required:"true"`

	// A list of DeleteMessageBatchResultEntry items.
	//
	// Successful is a required field
	Successful []*DeleteMessageBatchResultEntry `locationNameList:"DeleteMessageBatchResultEntry" type:"list" flattened:"true" required:"true"`
}

// String returns the string representation
func (s DeleteMessageBatchOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteMessageBatchOutput) GoString() string {
	return s.String()
}

// SetFailed sets the Failed field's value.
func (s *DeleteMessageBatchOutput) SetFailed(v []*BatchResultErrorEntry) *DeleteMessageBatchOutput {
	s.Failed = v
	return s
}

// SetSuccessful sets the Successful field's value.
func (s *DeleteMessageBatchOutput) SetSuccessful(v []*DeleteMessageBatchResultEntry) *DeleteMessageBatchOutput {
	s.Successful = v
	return s
}

// Encloses a receipt handle and an identifier for it.
type DeleteMessageBatchRequestEntry struct {
	_ struct{} `type:"structure"`

	// An identifier for this particular receipt handle. This is used to communicate
	// the result.
	//
	// The Ids of a batch request need to be unique within a request
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`

	// A receipt handle.
	//
	// ReceiptHandle is a required field
	ReceiptHandle *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMessageBatchRequestEntry) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteMessageBatchRequestEntry) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMessageBatchRequestEntry) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteMessageBatchRequestEntry"}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}
	if s.ReceiptHandle == nil {
		invalidParams.Add(request.NewErrParamRequired("ReceiptHandle"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetId sets the Id field's value.
func (s *DeleteMessageBatchRequestEntry) SetId(v string) *DeleteMessageBatchRequestEntry {
	s.Id = &v
	return s
}

// SetReceiptHandle sets the ReceiptHandle field's value.
func (s *DeleteMessageBatchRequestEntry) SetReceiptHandle(v string) *DeleteMessageBatchRequestEntry {
	s.ReceiptHandle = &v
	return s
}

// Encloses the Id of an entry in DeleteMessageBatch.
type DeleteMessageBatchResultEntry struct {
	_ struct{} `type:"structure"`

	// Represents a successfully deleted message.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMessageBatchResultEntry) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteMessageBatchResultEntry) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *DeleteMessageBatchResultEntry) SetId(v string) *DeleteMessageBatchResultEntry {
	s.Id = &v
	return s
}

type DeleteMessageInput struct {
	_ struct{} `type:"structure"`

	// The URL of the Amazon SQS queue from which messages are deleted.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`

	// The receipt handle associated with the message to delete.
	//
	// ReceiptHandle is a required field
	ReceiptHandle *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMessageInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteMessageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMessageInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteMessageInput"}
	if s.QueueUrl == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueUrl"))
	}
	if s.ReceiptHandle == nil {
		invalidParams.Add(request.NewErrParamRequired("ReceiptHandle"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetQueueUrl sets the QueueUrl field's value.
func (s *DeleteMessageInput) SetQueueUrl(v string) *DeleteMessageInput {
	s.QueueUrl = &v
	return s
}

// SetReceiptHandle sets the ReceiptHandle field's value.
func (s *DeleteMessageInput) SetReceiptHandle(v string) *DeleteMessageInput {
	s.ReceiptHandle = &v
	return s
}

type DeleteMessageOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteMessageOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteMessageOutput) GoString() string {
	return s.String()
}

type DeleteQueueInput struct {
	_ struct{} `type:"structure"`

	// The URL of the Amazon SQS queue to delete.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteQueueInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteQueueInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteQueueInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteQueueInput"}
	if s.QueueUrl == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueUrl"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetQueueUrl sets the QueueUrl field's value.
func (s *DeleteQueueInput) SetQueueUrl(v string) *DeleteQueueInput {
	s.QueueUrl = &v
	return s
}

type DeleteQueueOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteQueueOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteQueueOutput) GoString() string {
	return s.String()
}

type GetQueueAttributesInput struct {
	_ struct{} `type:"structure"`

	// A list of attributes for which to retrieve information.
	//
	// In the future, new attributes might be added. If you write code that calls
	// this action, we recommend that you structure your code so that it can handle
	// new attributes gracefully.
	//
	// The following attributes are supported:
	//
	//    * All - Returns all values.
	//
	//    * ApproximateNumberOfMessages - Returns the approximate number of messages
	//    available for retrieval from the queue.
	//
	//    * ApproximateNumberOfMessagesDelayed - Returns the approximate number
	//    of messages in the queue that are delayed and not available for reading
	//    immediately. This can happen when the queue is configured as a delay queue
	//    or when a message has been sent with a delay parameter.
	//
	//    * ApproximateNumberOfMessagesNotVisible - Returns the approximate number
	//    of messages that are in flight. Messages are considered to be in flight
	//    if they have been sent to a client but have not yet been deleted or have
	//    not yet reached the end of their visibility window.
	//
	//    * CreatedTimestamp - Returns the time when the queue was created in seconds
	//    (epoch time (http://en.wikipedia.org/wiki/Unix_time)).
	//
	//    * DelaySeconds - Returns the default delay on the queue in seconds.
	//
	//    * LastModifiedTimestamp - Returns the time when the queue was last changed
	//    in seconds (epoch time (http://en.wikipedia.org/wiki/Unix_time)).
	//
	//    * MaximumMessageSize - Returns the limit of how many bytes a message can
	//    contain before Amazon SQS rejects it.
	//
	//    * MessageRetentionPeriod - Returns the length of time, in seconds, for
	//    which Amazon SQS retains a message.
	//
	//    * Policy - Returns the policy of the queue.
	//
	//    * QueueArn - Returns the Amazon resource name (ARN) of the queue.
	//
	//    * ReceiveMessageWaitTimeSeconds - Returns the length of time, in seconds,
	//    for which the ReceiveMessage action waits for a message to arrive.
	//
	//    * RedrivePolicy - Returns the string that includes the parameters for
	//    dead-letter queue functionality of the source queue. For more information
	//    about the redrive policy and dead-letter queues, see Using Amazon SQS
	//    Dead-Letter Queues (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html)
	//    in the Amazon Simple Queue Service Developer Guide. deadLetterTargetArn
	//    - The Amazon Resource Name (ARN) of the dead-letter queue to which Amazon
	//    SQS moves messages after the value of maxReceiveCount is exceeded. maxReceiveCount
	//    - The number of times a message is delivered to the source queue before
	//    being moved to the dead-letter queue. When the ReceiveCount for a message
	//    exceeds the maxReceiveCount for a queue, Amazon SQS moves the message
	//    to the dead-letter-queue.
	//
	//    * VisibilityTimeout - Returns the visibility timeout for the queue. For
	//    more information about the visibility timeout, see Visibility Timeout
	//    (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html)
	//    in the Amazon Simple Queue Service Developer Guide.
	//
	// The following attributes apply only to server-side-encryption (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html):
	//
	//    * KmsMasterKeyId - Returns the ID of an AWS-managed customer master key
	//    (CMK) for Amazon SQS or a custom CMK. For more information, see Key Terms
	//    (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
	//
	//    * KmsDataKeyReusePeriodSeconds - Returns the length of time, in seconds,
	//    for which Amazon SQS can reuse a data key to encrypt or decrypt messages
	//    before calling AWS KMS again. For more information, see How Does the Data
	//    Key Reuse Period Work? (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-how-does-the-data-key-reuse-period-work).
	//
	// The following attributes apply only to FIFO (first-in-first-out) queues (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html):
	//
	//    * FifoQueue - Returns whether the queue is FIFO. For more information,
	//    see FIFO Queue Logic (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-understanding-logic)
	//    in the Amazon Simple Queue Service Developer Guide. To determine whether
	//    a queue is FIFO (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html),
	//    you can check whether QueueName ends with the .fifo suffix.
	//
	//    * ContentBasedDeduplication - Returns whether content-based deduplication
	//    is enabled for the queue. For more information, see Exactly-Once Processing
	//    (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
	//    in the Amazon Simple Queue Service Developer Guide.
	AttributeNames []*string `locationNameList:"AttributeName" type:"list" flattened:"true"`

	// The URL of the Amazon SQS queue whose attribute information is retrieved.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetQueueAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetQueueAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetQueueAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetQueueAttributesInput"}
	if s.QueueUrl == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueUrl"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAttributeNames sets the AttributeNames field's value.
func (s *GetQueueAttributesInput) SetAttributeNames(v []*string) *GetQueueAttributesInput {
	s.AttributeNames = v
	return s
}

// SetQueueUrl sets the QueueUrl field's value.
func (s *GetQueueAttributesInput) SetQueueUrl(v string) *GetQueueAttributesInput {
	s.QueueUrl = &v
	return s
}

// A list of returned queue attributes.
type GetQueueAttributesOutput struct {
	_ struct{} `type:"structure"`

	// A map of attributes to their respective values.
	Attributes map[string]*string `locationName:"Attribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true"`
}

// String returns the string representation
func (s GetQueueAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetQueueAttributesOutput) GoString() string {
	return s.String()
}

// SetAttributes sets the Attributes field's value.
func (s *GetQueueAttributesOutput) SetAttributes(v map[string]*string) *GetQueueAttributesOutput {
	s.Attributes = v
	return s
}

type GetQueueUrlInput struct {
	_ struct{} `type:"structure"`

	// The name of the queue whose URL must be fetched. Maximum 80 characters. Valid
	// values: alphanumeric characters, hyphens (-), and underscores (_).
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueName is a required field
	QueueName *string `type:"string" required:"true"`

	// The AWS account ID of the account that created the queue.
	QueueOwnerAWSAccountId *string `type:"string"`
}

// String returns the string representation
func (s GetQueueUrlInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetQueueUrlInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetQueueUrlInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetQueueUrlInput"}
	if s.QueueName == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetQueueName sets the QueueName field's value.
func (s *GetQueueUrlInput) SetQueueName(v string) *GetQueueUrlInput {
	s.QueueName = &v
	return s
}

// SetQueueOwnerAWSAccountId sets the QueueOwnerAWSAccountId field's value.
func (s *GetQueueUrlInput) SetQueueOwnerAWSAccountId(v string) *GetQueueUrlInput {
	s.QueueOwnerAWSAccountId = &v
	return s
}

// For more information, see Interpreting Responses (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-api-responses.html)
// in the Amazon Simple Queue Service Developer Guide.
type GetQueueUrlOutput struct {
	_ struct{} `type:"structure"`

	// The URL of the queue.
	QueueUrl *string `type:"string"`
}

// String returns the string representation
func (s GetQueueUrlOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetQueueUrlOutput) GoString() string {
	return s.String()
}

// SetQueueUrl sets the QueueUrl field's value.
func (s *GetQueueUrlOutput) SetQueueUrl(v string) *GetQueueUrlOutput {
	s.QueueUrl = &v
	return s
}

type ListDeadLetterSourceQueuesInput struct {
	_ struct{} `type:"structure"`

	// The URL of a dead-letter queue.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListDeadLetterSourceQueuesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListDeadLetterSourceQueuesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListDeadLetterSourceQueuesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListDeadLetterSourceQueuesInput"}
	if s.QueueUrl == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueUrl"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetQueueUrl sets the QueueUrl field's value.
func (s *ListDeadLetterSourceQueuesInput) SetQueueUrl(v string) *ListDeadLetterSourceQueuesInput {
	s.QueueUrl = &v
	return s
}

// A list of your dead letter source queues.
type ListDeadLetterSourceQueuesOutput struct {
	_ struct{} `type:"structure"`

	// A list of source queue URLs that have the RedrivePolicy queue attribute configured
	// with a dead-letter queue.
	//
	// QueueUrls is a required field
	QueueUrls []*string `locationName:"queueUrls" locationNameList:"QueueUrl" type:"list" flattened:"true" required:"true"`
}

// String returns the string representation
func (s ListDeadLetterSourceQueuesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListDeadLetterSourceQueuesOutput) GoString() string {
	return s.String()
}

// SetQueueUrls sets the QueueUrls field's value.
func (s *ListDeadLetterSourceQueuesOutput) SetQueueUrls(v []*string) *ListDeadLetterSourceQueuesOutput {
	s.QueueUrls = v
	return s
}

type ListQueueTagsInput struct {
	_ struct{} `type:"structure"`

	// The URL of the queue.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ListQueueTagsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListQueueTagsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListQueueTagsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListQueueTagsInput"}
	if s.QueueUrl == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueUrl"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetQueueUrl sets the QueueUrl field's value.
func (s *ListQueueTagsInput) SetQueueUrl(v string) *ListQueueTagsInput {
	s.QueueUrl = &v
	return s
}

type ListQueueTagsOutput struct {
	_ struct{} `type:"structure"`

	// The list of all tags added to the specified queue.
	Tags map[string]*string `locationName:"Tag" locationNameKey:"Key" locationNameValue:"Value" type:"map" flattened:"true"`
}

// String returns the string representation
func (s ListQueueTagsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListQueueTagsOutput) GoString() string {
	return s.String()
}

// SetTags sets the Tags field's value.
func (s *ListQueueTagsOutput) SetTags(v map[string]*string) *ListQueueTagsOutput {
	s.Tags = v
	return s
}

type ListQueuesInput struct {
	_ struct{} `type:"structure"`

	// A string to use for filtering the list results. Only those queues whose name
	// begins with the specified string are returned.
	//
	// Queue URLs and names are case-sensitive.
	QueueNamePrefix *string `type:"string"`
}

// String returns the string representation
func (s ListQueuesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListQueuesInput) GoString() string {
	return s.String()
}

// SetQueueNamePrefix sets the QueueNamePrefix field's value.
func (s *ListQueuesInput) SetQueueNamePrefix(v string) *ListQueuesInput {
	s.QueueNamePrefix = &v
	return s
}

// A list of your queues.
type ListQueuesOutput struct {
	_ struct{} `type:"structure"`

	// A list of queue URLs, up to 1,000 entries.
	QueueUrls []*string `locationNameList:"QueueUrl" type:"list" flattened:"true"`
}

// String returns the string representation
func (s ListQueuesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListQueuesOutput) GoString() string {
	return s.String()
}

// SetQueueUrls sets the QueueUrls field's value.
func (s *ListQueuesOutput) SetQueueUrls(v []*string) *ListQueuesOutput {
	s.QueueUrls = v
	return s
}

// An Amazon SQS message.
type Message struct {
	_ struct{} `type:"structure"`

	// A map of the attributes requested in ReceiveMessage to their respective values.
	// Supported attributes:
	//
	//    * ApproximateReceiveCount
	//
	//    * ApproximateFirstReceiveTimestamp
	//
	//    * MessageDeduplicationId
	//
	//    * MessageGroupId
	//
	//    * SenderId
	//
	//    * SentTimestamp
	//
	//    * SequenceNumber
	//
	// ApproximateFirstReceiveTimestamp and SentTimestamp are each returned as an
	// integer representing the epoch time (http://en.wikipedia.org/wiki/Unix_time)
	// in milliseconds.
	Attributes map[string]*string `locationName:"Attribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true"`

	// The message's contents (not URL-encoded).
	Body *string `type:"string"`

	// An MD5 digest of the non-URL-encoded message body string.
	MD5OfBody *string `type:"string"`

	// An MD5 digest of the non-URL-encoded message attribute string. You can use
	// this attribute to verify that Amazon SQS received the message correctly.
	// Amazon SQS URL-decodes the message before creating the MD5 digest. For information
	// about MD5, see RFC1321 (https://www.ietf.org/rfc/rfc1321.txt).
	MD5OfMessageAttributes *string `type:"string"`

	// Each message attribute consists of a Name, Type, and Value. For more information,
	// see Amazon SQS Message Attributes (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-message-attributes.html)
	// in the Amazon Simple Queue Service Developer Guide.
	MessageAttributes map[string]*MessageAttributeValue `locationName:"MessageAttribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true"`

	// A unique identifier for the message. A MessageIdis considered unique across
	// all AWS accounts for an extended period of time.
	MessageId *string `type:"string"`

	// An identifier associated with the act of receiving the message. A new receipt
	// handle is returned every time you receive a message. When deleting a message,
	// you provide the last received receipt handle to delete the message.
	ReceiptHandle *string `type:"string"`
}

// String returns the string representation
func (s Message) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Message) GoString() string {
	return s.String()
}

// SetAttributes sets the Attributes field's value.
func (s *Message) SetAttributes(v map[string]*string) *Message {
	s.Attributes = v
	return s
}

// SetBody sets the Body field's value.
func (s *Message) SetBody(v string) *Message {
	s.Body = &v
	return s
}

// SetMD5OfBody sets the MD5OfBody field's value.
func (s *Message) SetMD5OfBody(v string) *Message {
	s.MD5OfBody = &v
	return s
}

// SetMD5OfMessageAttributes sets the MD5OfMessageAttributes field's value.
func (s *Message) SetMD5OfMessageAttributes(v string) *Message {
	s.MD5OfMessageAttributes = &v
	return s
}

// SetMessageAttributes sets the MessageAttributes field's value.
func (s *Message) SetMessageAttributes(v map[string]*MessageAttributeValue) *Message {
	s.MessageAttributes = v
	return s
}

// SetMessageId sets the MessageId field's value.
func (s *Message) SetMessageId(v string) *Message {
	s.MessageId = &v
	return s
}

// SetReceiptHandle sets the ReceiptHandle field's value.
func (s *Message) SetReceiptHandle(v string) *Message {
	s.ReceiptHandle = &v
	return s
}

// The user-specified message attribute value. For string data types, the Value
// attribute has the same restrictions on the content as the message body. For
// more information, see SendMessage.
//
// Name, type, value and the message body must not be empty or null. All parts
// of the message attribute, including Name, Type, and Value, are part of the
// message size restriction (256 KB or 262,144 bytes).
type MessageAttributeValue struct {
	_ struct{} `type:"structure"`

	// Not implemented. Reserved for future use.
	BinaryListValues [][]byte `locationName:"BinaryListValue" locationNameList:"BinaryListValue" type:"list" flattened:"true"`

	// Binary type attributes can store any binary data, such as compressed data,
	// encrypted data, or images.
	//
	// BinaryValue is automatically base64 encoded/decoded by the SDK.
	BinaryValue []byte `type:"blob"`

	// Amazon SQS supports the following logical data types: String, Number, and
	// Binary. For the Number data type, you must use StringValue.
	//
	// You can also append custom labels. For more information, see Amazon SQS Message
	// Attributes (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-message-attributes.html)
	// in the Amazon Simple Queue Service Developer Guide.
	//
	// DataType is a required field
	DataType *string `type:"string" required:"true"`

	// Not implemented. Reserved for future use.
	StringListValues []*string `locationName:"StringListValue" locationNameList:"StringListValue" type:"list" flattened:"true"`

	// Strings are Unicode with UTF-8 binary encoding. For a list of code values,
	// see ASCII Printable Characters (http://en.wikipedia.org/wiki/ASCII#ASCII_printable_characters).
	StringValue *string `type:"string"`
}

// String returns the string representation
func (s MessageAttributeValue) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s MessageAttributeValue) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MessageAttributeValue) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "MessageAttributeValue"}
	if s.DataType == nil {
		invalidParams.Add(request.NewErrParamRequired("DataType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBinaryListValues sets the BinaryListValues field's value.
func (s *MessageAttributeValue) SetBinaryListValues(v [][]byte) *MessageAttributeValue {
	s.BinaryListValues = v
	return s
}

// SetBinaryValue sets the BinaryValue field's value.
func (s *MessageAttributeValue) SetBinaryValue(v []byte) *MessageAttributeValue {
	s.BinaryValue = v
	return s
}

// SetDataType sets the DataType field's value.
func (s *MessageAttributeValue) SetDataType(v string) *MessageAttributeValue {
	s.DataType = &v
	return s
}

// SetStringListValues sets the StringListValues field's value.
func (s *MessageAttributeValue) SetStringListValues(v []*string) *MessageAttributeValue {
	s.StringListValues = v
	return s
}

// SetStringValue sets the StringValue field's value.
func (s *MessageAttributeValue) SetStringValue(v string) *MessageAttributeValue {
	s.StringValue = &v
	return s
}

// The user-specified message system attribute value. For string data types,
// the Value attribute has the same restrictions on the content as the message
// body. For more information, see SendMessage.
//
// Name, type, value and the message body must not be empty or null.
type MessageSystemAttributeValue struct {
	_ struct{} `type:"structure"`

	// Not implemented. Reserved for future use.
	BinaryListValues [][]byte `locationName:"BinaryListValue" locationNameList:"BinaryListValue" type:"list" flattened:"true"`

	// Binary type attributes can store any binary data, such as compressed data,
	// encrypted data, or images.
	//
	// BinaryValue is automatically base64 encoded/decoded by the SDK.
	BinaryValue []byte `type:"blob"`

	// Amazon SQS supports the following logical data types: String, Number, and
	// Binary. For the Number data type, you must use StringValue.
	//
	// You can also append custom labels. For more information, see Amazon SQS Message
	// Attributes (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-message-attributes.html)
	// in the Amazon Simple Queue Service Developer Guide.
	//
	// DataType is a required field
	DataType *string `type:"string" required:"true"`

	// Not implemented. Reserved for future use.
	StringListValues []*string `locationName:"StringListValue" locationNameList:"StringListValue" type:"list" flattened:"true"`

	// Strings are Unicode with UTF-8 binary encoding. For a list of code values,
	// see ASCII Printable Characters (http://en.wikipedia.org/wiki/ASCII#ASCII_printable_characters).
	StringValue *string `type:"string"`
}

// String returns the string representation
func (s MessageSystemAttributeValue) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s MessageSystemAttributeValue) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MessageSystemAttributeValue) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "MessageSystemAttributeValue"}
	if s.DataType == nil {
		invalidParams.Add(request.NewErrParamRequired("DataType"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBinaryListValues sets the BinaryListValues field's value.
func (s *MessageSystemAttributeValue) SetBinaryListValues(v [][]byte) *MessageSystemAttributeValue {
	s.BinaryListValues = v
	return s
}

// SetBinaryValue sets the BinaryValue field's value.
func (s *MessageSystemAttributeValue) SetBinaryValue(v []byte) *MessageSystemAttributeValue {
	s.BinaryValue = v
	return s
}

// SetDataType sets the DataType field's value.
func (s *MessageSystemAttributeValue) SetDataType(v string) *MessageSystemAttributeValue {
	s.DataType = &v
	return s
}

// SetStringListValues sets the StringListValues field's value.
func (s *MessageSystemAttributeValue) SetStringListValues(v []*string) *MessageSystemAttributeValue {
	s.StringListValues = v
	return s
}

// SetStringValue sets the StringValue field's value.
func (s *MessageSystemAttributeValue) SetStringValue(v string) *MessageSystemAttributeValue {
	s.StringValue = &v
	return s
}

type PurgeQueueInput struct {
	_ struct{} `type:"structure"`

	// The URL of the queue from which the PurgeQueue action deletes messages.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`
}

// String returns the string representation
func (s PurgeQueueInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PurgeQueueInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PurgeQueueInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PurgeQueueInput"}
	if s.QueueUrl == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueUrl"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetQueueUrl sets the QueueUrl field's value.
func (s *PurgeQueueInput) SetQueueUrl(v string) *PurgeQueueInput {
	s.QueueUrl = &v
	return s
}

type PurgeQueueOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PurgeQueueOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PurgeQueueOutput) GoString() string {
	return s.String()
}

type ReceiveMessageInput struct {
	_ struct{} `type:"structure"`

	// A list of attributes that need to be returned along with each message. These
	// attributes include:
	//
	//    * All - Returns all values.
	//
	//    * ApproximateFirstReceiveTimestamp - Returns the time the message was
	//    first received from the queue (epoch time (http://en.wikipedia.org/wiki/Unix_time)
	//    in milliseconds).
	//
	//    * ApproximateReceiveCount - Returns the number of times a message has
	//    been received from the queue but not deleted.
	//
	//    * AWSTraceHeader - Returns the AWS X-Ray trace header string.
	//
	//    * SenderId For an IAM user, returns the IAM user ID, for example ABCDEFGHI1JKLMNOPQ23R.
	//    For an IAM role, returns the IAM role ID, for example ABCDE1F2GH3I4JK5LMNOP:i-a123b456.
	//
	//    * SentTimestamp - Returns the time the message was sent to the queue (epoch
	//    time (http://en.wikipedia.org/wiki/Unix_time) in milliseconds).
	//
	//    * MessageDeduplicationId - Returns the value provided by the producer
	//    that calls the SendMessage action.
	//
	//    * MessageGroupId - Returns the value provided by the producer that calls
	//    the SendMessage action. Messages with the same MessageGroupId are returned
	//    in sequence.
	//
	//    * SequenceNumber - Returns the value provided by Amazon SQS.
	AttributeNames []*string `locationNameList:"AttributeName" type:"list" flattened:"true"`

	// The maximum number of messages to return. Amazon SQS never returns more messages
	// than this value (however, fewer messages might be returned). Valid values:
	// 1 to 10. Default: 1.
	MaxNumberOfMessages *int64 `type:"integer"`

	// The name of the message attribute, where N is the index.
	//
	//    * The name can contain alphanumeric characters and the underscore (_),
	//    hyphen (-), and period (.).
	//
	//    * The name is case-sensitive and must be unique among all attribute names
	//    for the message.
	//
	//    * The name must not start with AWS-reserved prefixes such as AWS. or Amazon.
	//    (or any casing variants).
	//
	//    * The name must not start or end with a period (.), and it should not
	//    have periods in succession (..).
	//
	//    * The name can be up to 256 characters long.
	//
	// When using ReceiveMessage, you can send a list of attribute names to receive,
	// or you can return all of the attributes by specifying All or .* in your request.
	// You can also use all message attributes starting with a prefix, for example
	// bar.*.
	MessageAttributeNames []*string `locationNameList:"MessageAttributeName" type:"list" flattened:"true"`

	// The URL of the Amazon SQS queue from which messages are received.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`

	// This parameter applies only to FIFO (first-in-first-out) queues.
	//
	// The token used for deduplication of ReceiveMessage calls. If a networking
	// issue occurs after a ReceiveMessage action, and instead of a response you
	// receive a generic error, you can retry the same action with an identical
	// ReceiveRequestAttemptId to retrieve the same set of messages, even if their
	// visibility timeout has not yet expired.
	//
	//    * You can use ReceiveRequestAttemptId only for 5 minutes after a ReceiveMessage
	//    action.
	//
	//    * When you set FifoQueue, a caller of the ReceiveMessage action can provide
	//    a ReceiveRequestAttemptId explicitly.
	//
	//    * If a caller of the ReceiveMessage action doesn't provide a ReceiveRequestAttemptId,
	//    Amazon SQS generates a ReceiveRequestAttemptId.
	//
	//    * You can retry the ReceiveMessage action with the same ReceiveRequestAttemptId
	//    if none of the messages have been modified (deleted or had their visibility
	//    changes).
	//
	//    * During a visibility timeout, subsequent calls with the same ReceiveRequestAttemptId
	//    return the same messages and receipt handles. If a retry occurs within
	//    the deduplication interval, it resets the visibility timeout. For more
	//    information, see Visibility Timeout (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html)
	//    in the Amazon Simple Queue Service Developer Guide. If a caller of the
	//    ReceiveMessage action still processes messages when the visibility timeout
	//    expires and messages become visible, another worker consuming from the
	//    same queue can receive the same messages and therefore process duplicates.
	//    Also, if a consumer whose message processing time is longer than the visibility
	//    timeout tries to delete the processed messages, the action fails with
	//    an error. To mitigate this effect, ensure that your application observes
	//    a safe threshold before the visibility timeout expires and extend the
	//    visibility timeout as necessary.
	//
	//    * While messages with a particular MessageGroupId are invisible, no more
	//    messages belonging to the same MessageGroupId are returned until the visibility
	//    timeout expires. You can still receive messages with another MessageGroupId
	//    as long as it is also visible.
	//
	//    * If a caller of ReceiveMessage can't track the ReceiveRequestAttemptId,
	//    no retries work until the original visibility timeout expires. As a result,
	//    delays might occur but the messages in the queue remain in a strict order.
	//
	// The length of ReceiveRequestAttemptId is 128 characters. ReceiveRequestAttemptId
	// can contain alphanumeric characters (a-z, A-Z, 0-9) and punctuation (!"#$%&'()*+,-./:;<=>?@[\]^_`{|}~).
	//
	// For best practices of using ReceiveRequestAttemptId, see Using the ReceiveRequestAttemptId
	// Request Parameter (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/using-receiverequestattemptid-request-parameter.html)
	// in the Amazon Simple Queue Service Developer Guide.
	ReceiveRequestAttemptId *string `type:"string"`

	// The duration (in seconds) that the received messages are hidden from subsequent
	// retrieve requests after being retrieved by a ReceiveMessage request.
	VisibilityTimeout *int64 `type:"integer"`

	// The duration (in seconds) for which the call waits for a message to arrive
	// in the queue before returning. If a message is available, the call returns
	// sooner than WaitTimeSeconds. If no messages are available and the wait time
	// expires, the call returns successfully with an empty list of messages.
	WaitTimeSeconds *int64 `type:"integer"`
}

// String returns the string representation
func (s ReceiveMessageInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ReceiveMessageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ReceiveMessageInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ReceiveMessageInput"}
	if s.QueueUrl == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueUrl"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAttributeNames sets the AttributeNames field's value.
func (s *ReceiveMessageInput) SetAttributeNames(v []*string) *ReceiveMessageInput {
	s.AttributeNames = v
	return s
}

// SetMaxNumberOfMessages sets the MaxNumberOfMessages field's value.
func (s *ReceiveMessageInput) SetMaxNumberOfMessages(v int64) *ReceiveMessageInput {
	s.MaxNumberOfMessages = &v
	return s
}

// SetMessageAttributeNames sets the MessageAttributeNames field's value.
func (s *ReceiveMessageInput) SetMessageAttributeNames(v []*string) *ReceiveMessageInput {
	s.MessageAttributeNames = v
	return s
}

// SetQueueUrl sets the QueueUrl field's value.
func (s *ReceiveMessageInput) SetQueueUrl(v string) *ReceiveMessageInput {
	s.QueueUrl = &v
	return s
}

// SetReceiveRequestAttemptId sets the ReceiveRequestAttemptId field's value.
func (s *ReceiveMessageInput) SetReceiveRequestAttemptId(v string) *ReceiveMessageInput {
	s.ReceiveRequestAttemptId = &v
	return s
}

// SetVisibilityTimeout sets the VisibilityTimeout field's value.
func (s *ReceiveMessageInput) SetVisibilityTimeout(v int64) *ReceiveMessageInput {
	s.VisibilityTimeout = &v
	return s
}

// SetWaitTimeSeconds sets the WaitTimeSeconds field's value.
func (s *ReceiveMessageInput) SetWaitTimeSeconds(v int64) *ReceiveMessageInput {
	s.WaitTimeSeconds = &v
	return s
}

// A list of received messages.
type ReceiveMessageOutput struct {
	_ struct{} `type:"structure"`

	// A list of messages.
	Messages []*Message `locationNameList:"Message" type:"list" flattened:"true"`
}

// String returns the string representation
func (s ReceiveMessageOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ReceiveMessageOutput) GoString() string {
	return s.String()
}

// SetMessages sets the Messages field's value.
func (s *ReceiveMessageOutput) SetMessages(v []*Message) *ReceiveMessageOutput {
	s.Messages = v
	return s
}

type RemovePermissionInput struct {
	_ struct{} `type:"structure"`

	// The identification of the permission to remove. This is the label added using
	// the AddPermission action.
	//
	// Label is a required field
	Label *string `type:"string" required:"true"`

	// The URL of the Amazon SQS queue from which permissions are removed.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`
}

// String returns the string representation
func (s RemovePermissionInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RemovePermissionInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RemovePermissionInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "RemovePermissionInput"}
	if s.Label == nil {
		invalidParams.Add(request.NewErrParamRequired("Label"))
	}
	if s.QueueUrl == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueUrl"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetLabel sets the Label field's value.
func (s *RemovePermissionInput) SetLabel(v string) *RemovePermissionInput {
	s.Label = &v
	return s
}

// SetQueueUrl sets the QueueUrl field's value.
func (s *RemovePermissionInput) SetQueueUrl(v string) *RemovePermissionInput {
	s.QueueUrl = &v
	return s
}

type RemovePermissionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s RemovePermissionOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s RemovePermissionOutput) GoString() string {
	return s.String()
}

type SendMessageBatchInput struct {
	_ struct{} `type:"structure"`

	// A list of SendMessageBatchRequestEntry items.
	//
	// Entries is a required field
	Entries []*SendMessageBatchRequestEntry `locationNameList:"SendMessageBatchRequestEntry" type:"list" flattened:"true" required:"true"`

	// The URL of the Amazon SQS queue to which batched messages are sent.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SendMessageBatchInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SendMessageBatchInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendMessageBatchInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SendMessageBatchInput"}
	if s.Entries == nil {
		invalidParams.Add(request.NewErrParamRequired("Entries"))
	}
	if s.QueueUrl == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueUrl"))
	}
	if s.Entries != nil {
		for i, v := range s.Entries {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Entries", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetEntries sets the Entries field's value.
func (s *SendMessageBatchInput) SetEntries(v []*SendMessageBatchRequestEntry) *SendMessageBatchInput {
	s.Entries = v
	return s
}

// SetQueueUrl sets the QueueUrl field's value.
func (s *SendMessageBatchInput) SetQueueUrl(v string) *SendMessageBatchInput {
	s.QueueUrl = &v
	return s
}

// For each message in the batch, the response contains a SendMessageBatchResultEntry
// tag if the message succeeds or a BatchResultErrorEntry tag if the message
// fails.
type SendMessageBatchOutput struct {
	_ struct{} `type:"structure"`

	// A list of BatchResultErrorEntry items with error details about each message
	// that can't be enqueued.
	//
	// Failed is a required field
	Failed []*BatchResultErrorEntry `locationNameList:"BatchResultErrorEntry" type:"list" flattened:"true" required:"true"`

	// A list of SendMessageBatchResultEntry items.
	//
	// Successful is a required field
	Successful []*SendMessageBatchResultEntry `locationNameList:"SendMessageBatchResultEntry" type:"list" flattened:"true" required:"true"`
}

// String returns the string representation
func (s SendMessageBatchOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SendMessageBatchOutput) GoString() string {
	return s.String()
}

// SetFailed sets the Failed field's value.
func (s *SendMessageBatchOutput) SetFailed(v []*BatchResultErrorEntry) *SendMessageBatchOutput {
	s.Failed = v
	return s
}

// SetSuccessful sets the Successful field's value.
func (s *SendMessageBatchOutput) SetSuccessful(v []*SendMessageBatchResultEntry) *SendMessageBatchOutput {
	s.Successful = v
	return s
}

// Contains the details of a single Amazon SQS message along with an Id.
type SendMessageBatchRequestEntry struct {
	_ struct{} `type:"structure"`

	// The length of time, in seconds, for which a specific message is delayed.
	// Valid values: 0 to 900. Maximum: 15 minutes. Messages with a positive DelaySeconds
	// value become available for processing after the delay period is finished.
	// If you don't specify a value, the default value for the queue is applied.
	//
	// When you set FifoQueue, you can't set DelaySeconds per message. You can set
	// this parameter only on a queue level.
	DelaySeconds *int64 `type:"integer"`

	// An identifier for a message in this batch used to communicate the result.
	//
	// The Ids of a batch request need to be unique within a request
	//
	// This identifier can have up to 80 characters. The following characters are
	// accepted: alphanumeric characters, hyphens(-), and underscores (_).
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`

	// Each message attribute consists of a Name, Type, and Value. For more information,
	// see Amazon SQS Message Attributes (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-message-attributes.html)
	// in the Amazon Simple Queue Service Developer Guide.
	MessageAttributes map[string]*MessageAttributeValue `locationName:"MessageAttribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true"`

	// The body of the message.
	//
	// MessageBody is a required field
	MessageBody *string `type:"string" required:"true"`

	// This parameter applies only to FIFO (first-in-first-out) queues.
	//
	// The token used for deduplication of messages within a 5-minute minimum deduplication
	// interval. If a message with a particular MessageDeduplicationId is sent successfully,
	// subsequent messages with the same MessageDeduplicationId are accepted successfully
	// but aren't delivered. For more information, see Exactly-Once Processing (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
	// in the Amazon Simple Queue Service Developer Guide.
	//
	//    * Every message must have a unique MessageDeduplicationId, You may provide
	//    a MessageDeduplicationId explicitly. If you aren't able to provide a MessageDeduplicationId
	//    and you enable ContentBasedDeduplication for your queue, Amazon SQS uses
	//    a SHA-256 hash to generate the MessageDeduplicationId using the body of
	//    the message (but not the attributes of the message). If you don't provide
	//    a MessageDeduplicationId and the queue doesn't have ContentBasedDeduplication
	//    set, the action fails with an error. If the queue has ContentBasedDeduplication
	//    set, your MessageDeduplicationId overrides the generated one.
	//
	//    * When ContentBasedDeduplication is in effect, messages with identical
	//    content sent within the deduplication interval are treated as duplicates
	//    and only one copy of the message is delivered.
	//
	//    * If you send one message with ContentBasedDeduplication enabled and then
	//    another message with a MessageDeduplicationId that is the same as the
	//    one generated for the first MessageDeduplicationId, the two messages are
	//    treated as duplicates and only one copy of the message is delivered.
	//
	// The MessageDeduplicationId is available to the consumer of the message (this
	// can be useful for troubleshooting delivery issues).
	//
	// If a message is sent successfully but the acknowledgement is lost and the
	// message is resent with the same MessageDeduplicationId after the deduplication
	// interval, Amazon SQS can't detect duplicate messages.
	//
	// Amazon SQS continues to keep track of the message deduplication ID even after
	// the message is received and deleted.
	//
	// The length of MessageDeduplicationId is 128 characters. MessageDeduplicationId
	// can contain alphanumeric characters (a-z, A-Z, 0-9) and punctuation (!"#$%&'()*+,-./:;<=>?@[\]^_`{|}~).
	//
	// For best practices of using MessageDeduplicationId, see Using the MessageDeduplicationId
	// Property (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/using-messagededuplicationid-property.html)
	// in the Amazon Simple Queue Service Developer Guide.
	MessageDeduplicationId *string `type:"string"`

	// This parameter applies only to FIFO (first-in-first-out) queues.
	//
	// The tag that specifies that a message belongs to a specific message group.
	// Messages that belong to the same message group are processed in a FIFO manner
	// (however, messages in different message groups might be processed out of
	// order). To interleave multiple ordered streams within a single queue, use
	// MessageGroupId values (for example, session data for multiple users). In
	// this scenario, multiple consumers can process the queue, but the session
	// data of each user is processed in a FIFO fashion.
	//
	//    * You must associate a non-empty MessageGroupId with a message. If you
	//    don't provide a MessageGroupId, the action fails.
	//
	//    * ReceiveMessage might return messages with multiple MessageGroupId values.
	//    For each MessageGroupId, the messages are sorted by time sent. The caller
	//    can't specify a MessageGroupId.
	//
	// The length of MessageGroupId is 128 characters. Valid values: alphanumeric
	// characters and punctuation (!"#$%&'()*+,-./:;<=>?@[\]^_`{|}~).
	//
	// For best practices of using MessageGroupId, see Using the MessageGroupId
	// Property (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/using-messagegroupid-property.html)
	// in the Amazon Simple Queue Service Developer Guide.
	//
	// MessageGroupId is required for FIFO queues. You can't use it for Standard
	// queues.
	MessageGroupId *string `type:"string"`

	// The message system attribute to send Each message system attribute consists
	// of a Name, Type, and Value.
	//
	//    * Currently, the only supported message system attribute is AWSTraceHeader.
	//    Its type must be String and its value must be a correctly formatted AWS
	//    X-Ray trace string.
	//
	//    * The size of a message system attribute doesn't count towards the total
	//    size of a message.
	MessageSystemAttributes map[string]*MessageSystemAttributeValue `locationName:"MessageSystemAttribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true"`
}

// String returns the string representation
func (s SendMessageBatchRequestEntry) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SendMessageBatchRequestEntry) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendMessageBatchRequestEntry) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SendMessageBatchRequestEntry"}
	if s.Id == nil {
		invalidParams.Add(request.NewErrParamRequired("Id"))
	}
	if s.MessageBody == nil {
		invalidParams.Add(request.NewErrParamRequired("MessageBody"))
	}
	if s.MessageAttributes != nil {
		for i, v := range s.MessageAttributes {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "MessageAttributes", i), err.(request.ErrInvalidParams))
			}
		}
	}
	if s.MessageSystemAttributes != nil {
		for i, v := range s.MessageSystemAttributes {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "MessageSystemAttributes", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDelaySeconds sets the DelaySeconds field's value.
func (s *SendMessageBatchRequestEntry) SetDelaySeconds(v int64) *SendMessageBatchRequestEntry {
	s.DelaySeconds = &v
	return s
}

// SetId sets the Id field's value.
func (s *SendMessageBatchRequestEntry) SetId(v string) *SendMessageBatchRequestEntry {
	s.Id = &v
	return s
}

// SetMessageAttributes sets the MessageAttributes field's value.
func (s *SendMessageBatchRequestEntry) SetMessageAttributes(v map[string]*MessageAttributeValue) *SendMessageBatchRequestEntry {
	s.MessageAttributes = v
	return s
}

// SetMessageBody sets the MessageBody field's value.
func (s *SendMessageBatchRequestEntry) SetMessageBody(v string) *SendMessageBatchRequestEntry {
	s.MessageBody = &v
	return s
}

// SetMessageDeduplicationId sets the MessageDeduplicationId field's value.
func (s *SendMessageBatchRequestEntry) SetMessageDeduplicationId(v string) *SendMessageBatchRequestEntry {
	s.MessageDeduplicationId = &v
	return s
}

// SetMessageGroupId sets the MessageGroupId field's value.
func (s *SendMessageBatchRequestEntry) SetMessageGroupId(v string) *SendMessageBatchRequestEntry {
	s.MessageGroupId = &v
	return s
}

// SetMessageSystemAttributes sets the MessageSystemAttributes field's value.
func (s *SendMessageBatchRequestEntry) SetMessageSystemAttributes(v map[string]*MessageSystemAttributeValue) *SendMessageBatchRequestEntry {
	s.MessageSystemAttributes = v
	return s
}

// Encloses a MessageId for a successfully-enqueued message in a SendMessageBatch.
type SendMessageBatchResultEntry struct {
	_ struct{} `type:"structure"`

	// An identifier for the message in this batch.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`

	// An MD5 digest of the non-URL-encoded message attribute string. You can use
	// this attribute to verify that Amazon SQS received the message correctly.
	// Amazon SQS URL-decodes the message before creating the MD5 digest. For information
	// about MD5, see RFC1321 (https://www.ietf.org/rfc/rfc1321.txt).
	MD5OfMessageAttributes *string `type:"string"`

	// An MD5 digest of the non-URL-encoded message attribute string. You can use
	// this attribute to verify that Amazon SQS received the message correctly.
	// Amazon SQS URL-decodes the message before creating the MD5 digest. For information
	// about MD5, see RFC1321 (https://www.ietf.org/rfc/rfc1321.txt).
	//
	// MD5OfMessageBody is a required field
	MD5OfMessageBody *string `type:"string" required:"true"`

	// An MD5 digest of the non-URL-encoded message system attribute string. You
	// can use this attribute to verify that Amazon SQS received the message correctly.
	// Amazon SQS URL-decodes the message before creating the MD5 digest. For information
	// about MD5, see RFC1321 (https://www.ietf.org/rfc/rfc1321.txt).
	MD5OfMessageSystemAttributes *string `type:"string"`

	// An identifier for the message.
	//
	// MessageId is a required field
	MessageId *string `type:"string" required:"true"`

	// This parameter applies only to FIFO (first-in-first-out) queues.
	//
	// The large, non-consecutive number that Amazon SQS assigns to each message.
	//
	// The length of SequenceNumber is 128 bits. As SequenceNumber continues to
	// increase for a particular MessageGroupId.
	SequenceNumber *string `type:"string"`
}

// String returns the string representation
func (s SendMessageBatchResultEntry) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SendMessageBatchResultEntry) GoString() string {
	return s.String()
}

// SetId sets the Id field's value.
func (s *SendMessageBatchResultEntry) SetId(v string) *SendMessageBatchResultEntry {
	s.Id = &v
	return s
}

// SetMD5OfMessageAttributes sets the MD5OfMessageAttributes field's value.
func (s *SendMessageBatchResultEntry) SetMD5OfMessageAttributes(v string) *SendMessageBatchResultEntry {
	s.MD5OfMessageAttributes = &v
	return s
}

// SetMD5OfMessageBody sets the MD5OfMessageBody field's value.
func (s *SendMessageBatchResultEntry) SetMD5OfMessageBody(v string) *SendMessageBatchResultEntry {
	s.MD5OfMessageBody = &v
	return s
}

// SetMD5OfMessageSystemAttributes sets the MD5OfMessageSystemAttributes field's value.
func (s *SendMessageBatchResultEntry) SetMD5OfMessageSystemAttributes(v string) *SendMessageBatchResultEntry {
	s.MD5OfMessageSystemAttributes = &v
	return s
}

// SetMessageId sets the MessageId field's value.
func (s *SendMessageBatchResultEntry) SetMessageId(v string) *SendMessageBatchResultEntry {
	s.MessageId = &v
	return s
}

// SetSequenceNumber sets the SequenceNumber field's value.
func (s *SendMessageBatchResultEntry) SetSequenceNumber(v string) *SendMessageBatchResultEntry {
	s.SequenceNumber = &v
	return s
}

type SendMessageInput struct {
	_ struct{} `type:"structure"`

	// The length of time, in seconds, for which to delay a specific message. Valid
	// values: 0 to 900. Maximum: 15 minutes. Messages with a positive DelaySeconds
	// value become available for processing after the delay period is finished.
	// If you don't specify a value, the default value for the queue applies.
	//
	// When you set FifoQueue, you can't set DelaySeconds per message. You can set
	// this parameter only on a queue level.
	DelaySeconds *int64 `type:"integer"`

	// Each message attribute consists of a Name, Type, and Value. For more information,
	// see Amazon SQS Message Attributes (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-message-attributes.html)
	// in the Amazon Simple Queue Service Developer Guide.
	MessageAttributes map[string]*MessageAttributeValue `locationName:"MessageAttribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true"`

	// The message to send. The maximum string size is 256 KB.
	//
	// A message can include only XML, JSON, and unformatted text. The following
	// Unicode characters are allowed:
	//
	// #x9 | #xA | #xD | #x20 to #xD7FF | #xE000 to #xFFFD | #x10000 to #x10FFFF
	//
	// Any characters not included in this list will be rejected. For more information,
	// see the W3C specification for characters (http://www.w3.org/TR/REC-xml/#charsets).
	//
	// MessageBody is a required field
	MessageBody *string `type:"string" required:"true"`

	// This parameter applies only to FIFO (first-in-first-out) queues.
	//
	// The token used for deduplication of sent messages. If a message with a particular
	// MessageDeduplicationId is sent successfully, any messages sent with the same
	// MessageDeduplicationId are accepted successfully but aren't delivered during
	// the 5-minute deduplication interval. For more information, see Exactly-Once
	// Processing (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
	// in the Amazon Simple Queue Service Developer Guide.
	//
	//    * Every message must have a unique MessageDeduplicationId, You may provide
	//    a MessageDeduplicationId explicitly. If you aren't able to provide a MessageDeduplicationId
	//    and you enable ContentBasedDeduplication for your queue, Amazon SQS uses
	//    a SHA-256 hash to generate the MessageDeduplicationId using the body of
	//    the message (but not the attributes of the message). If you don't provide
	//    a MessageDeduplicationId and the queue doesn't have ContentBasedDeduplication
	//    set, the action fails with an error. If the queue has ContentBasedDeduplication
	//    set, your MessageDeduplicationId overrides the generated one.
	//
	//    * When ContentBasedDeduplication is in effect, messages with identical
	//    content sent within the deduplication interval are treated as duplicates
	//    and only one copy of the message is delivered.
	//
	//    * If you send one message with ContentBasedDeduplication enabled and then
	//    another message with a MessageDeduplicationId that is the same as the
	//    one generated for the first MessageDeduplicationId, the two messages are
	//    treated as duplicates and only one copy of the message is delivered.
	//
	// The MessageDeduplicationId is available to the consumer of the message (this
	// can be useful for troubleshooting delivery issues).
	//
	// If a message is sent successfully but the acknowledgement is lost and the
	// message is resent with the same MessageDeduplicationId after the deduplication
	// interval, Amazon SQS can't detect duplicate messages.
	//
	// Amazon SQS continues to keep track of the message deduplication ID even after
	// the message is received and deleted.
	//
	// The length of MessageDeduplicationId is 128 characters. MessageDeduplicationId
	// can contain alphanumeric characters (a-z, A-Z, 0-9) and punctuation (!"#$%&'()*+,-./:;<=>?@[\]^_`{|}~).
	//
	// For best practices of using MessageDeduplicationId, see Using the MessageDeduplicationId
	// Property (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/using-messagededuplicationid-property.html)
	// in the Amazon Simple Queue Service Developer Guide.
	MessageDeduplicationId *string `type:"string"`

	// This parameter applies only to FIFO (first-in-first-out) queues.
	//
	// The tag that specifies that a message belongs to a specific message group.
	// Messages that belong to the same message group are processed in a FIFO manner
	// (however, messages in different message groups might be processed out of
	// order). To interleave multiple ordered streams within a single queue, use
	// MessageGroupId values (for example, session data for multiple users). In
	// this scenario, multiple consumers can process the queue, but the session
	// data of each user is processed in a FIFO fashion.
	//
	//    * You must associate a non-empty MessageGroupId with a message. If you
	//    don't provide a MessageGroupId, the action fails.
	//
	//    * ReceiveMessage might return messages with multiple MessageGroupId values.
	//    For each MessageGroupId, the messages are sorted by time sent. The caller
	//    can't specify a MessageGroupId.
	//
	// The length of MessageGroupId is 128 characters. Valid values: alphanumeric
	// characters and punctuation (!"#$%&'()*+,-./:;<=>?@[\]^_`{|}~).
	//
	// For best practices of using MessageGroupId, see Using the MessageGroupId
	// Property (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/using-messagegroupid-property.html)
	// in the Amazon Simple Queue Service Developer Guide.
	//
	// MessageGroupId is required for FIFO queues. You can't use it for Standard
	// queues.
	MessageGroupId *string `type:"string"`

	// The message system attribute to send. Each message system attribute consists
	// of a Name, Type, and Value.
	//
	//    * Currently, the only supported message system attribute is AWSTraceHeader.
	//    Its type must be String and its value must be a correctly formatted AWS
	//    X-Ray trace string.
	//
	//    * The size of a message system attribute doesn't count towards the total
	//    size of a message.
	MessageSystemAttributes map[string]*MessageSystemAttributeValue `locationName:"MessageSystemAttribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true"`

	// The URL of the Amazon SQS queue to which a message is sent.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SendMessageInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SendMessageInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SendMessageInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SendMessageInput"}
	if s.MessageBody == nil {
		invalidParams.Add(request.NewErrParamRequired("MessageBody"))
	}
	if s.QueueUrl == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueUrl"))
	}
	if s.MessageAttributes != nil {
		for i, v := range s.MessageAttributes {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "MessageAttributes", i), err.(request.ErrInvalidParams))
			}
		}
	}
	if s.MessageSystemAttributes != nil {
		for i, v := range s.MessageSystemAttributes {
			if v == nil {
				continue
			}
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "MessageSystemAttributes", i), err.(request.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDelaySeconds sets the DelaySeconds field's value.
func (s *SendMessageInput) SetDelaySeconds(v int64) *SendMessageInput {
	s.DelaySeconds = &v
	return s
}

// SetMessageAttributes sets the MessageAttributes field's value.
func (s *SendMessageInput) SetMessageAttributes(v map[string]*MessageAttributeValue) *SendMessageInput {
	s.MessageAttributes = v
	return s
}

// SetMessageBody sets the MessageBody field's value.
func (s *SendMessageInput) SetMessageBody(v string) *SendMessageInput {
	s.MessageBody = &v
	return s
}

// SetMessageDeduplicationId sets the MessageDeduplicationId field's value.
func (s *SendMessageInput) SetMessageDeduplicationId(v string) *SendMessageInput {
	s.MessageDeduplicationId = &v
	return s
}

// SetMessageGroupId sets the MessageGroupId field's value.
func (s *SendMessageInput) SetMessageGroupId(v string) *SendMessageInput {
	s.MessageGroupId = &v
	return s
}

// SetMessageSystemAttributes sets the MessageSystemAttributes field's value.
func (s *SendMessageInput) SetMessageSystemAttributes(v map[string]*MessageSystemAttributeValue) *SendMessageInput {
	s.MessageSystemAttributes = v
	return s
}

// SetQueueUrl sets the QueueUrl field's value.
func (s *SendMessageInput) SetQueueUrl(v string) *SendMessageInput {
	s.QueueUrl = &v
	return s
}

// The MD5OfMessageBody and MessageId elements.
type SendMessageOutput struct {
	_ struct{} `type:"structure"`

	// An MD5 digest of the non-URL-encoded message attribute string. You can use
	// this attribute to verify that Amazon SQS received the message correctly.
	// Amazon SQS URL-decodes the message before creating the MD5 digest. For information
	// about MD5, see RFC1321 (https://www.ietf.org/rfc/rfc1321.txt).
	MD5OfMessageAttributes *string `type:"string"`

	// An MD5 digest of the non-URL-encoded message attribute string. You can use
	// this attribute to verify that Amazon SQS received the message correctly.
	// Amazon SQS URL-decodes the message before creating the MD5 digest. For information
	// about MD5, see RFC1321 (https://www.ietf.org/rfc/rfc1321.txt).
	MD5OfMessageBody *string `type:"string"`

	// An MD5 digest of the non-URL-encoded message system attribute string. You
	// can use this attribute to verify that Amazon SQS received the message correctly.
	// Amazon SQS URL-decodes the message before creating the MD5 digest.
	MD5OfMessageSystemAttributes *string `type:"string"`

	// An attribute containing the MessageId of the message sent to the queue. For
	// more information, see Queue and Message Identifiers (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-queue-message-identifiers.html)
	// in the Amazon Simple Queue Service Developer Guide.
	MessageId *string `type:"string"`

	// This parameter applies only to FIFO (first-in-first-out) queues.
	//
	// The large, non-consecutive number that Amazon SQS assigns to each message.
	//
	// The length of SequenceNumber is 128 bits. SequenceNumber continues to increase
	// for a particular MessageGroupId.
	SequenceNumber *string `type:"string"`
}

// String returns the string representation
func (s SendMessageOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SendMessageOutput) GoString() string {
	return s.String()
}

// SetMD5OfMessageAttributes sets the MD5OfMessageAttributes field's value.
func (s *SendMessageOutput) SetMD5OfMessageAttributes(v string) *SendMessageOutput {
	s.MD5OfMessageAttributes = &v
	return s
}

// SetMD5OfMessageBody sets the MD5OfMessageBody field's value.
func (s *SendMessageOutput) SetMD5OfMessageBody(v string) *SendMessageOutput {
	s.MD5OfMessageBody = &v
	return s
}

// SetMD5OfMessageSystemAttributes sets the MD5OfMessageSystemAttributes field's value.
func (s *SendMessageOutput) SetMD5OfMessageSystemAttributes(v string) *SendMessageOutput {
	s.MD5OfMessageSystemAttributes = &v
	return s
}

// SetMessageId sets the MessageId field's value.
func (s *SendMessageOutput) SetMessageId(v string) *SendMessageOutput {
	s.MessageId = &v
	return s
}

// SetSequenceNumber sets the SequenceNumber field's value.
func (s *SendMessageOutput) SetSequenceNumber(v string) *SendMessageOutput {
	s.SequenceNumber = &v
	return s
}

type SetQueueAttributesInput struct {
	_ struct{} `type:"structure"`

	// A map of attributes to set.
	//
	// The following lists the names, descriptions, and values of the special request
	// parameters that the SetQueueAttributes action uses:
	//
	//    * DelaySeconds - The length of time, in seconds, for which the delivery
	//    of all messages in the queue is delayed. Valid values: An integer from
	//    0 to 900 (15 minutes). Default: 0.
	//
	//    * MaximumMessageSize - The limit of how many bytes a message can contain
	//    before Amazon SQS rejects it. Valid values: An integer from 1,024 bytes
	//    (1 KiB) up to 262,144 bytes (256 KiB). Default: 262,144 (256 KiB).
	//
	//    * MessageRetentionPeriod - The length of time, in seconds, for which Amazon
	//    SQS retains a message. Valid values: An integer representing seconds,
	//    from 60 (1 minute) to 1,209,600 (14 days). Default: 345,600 (4 days).
	//
	//    * Policy - The queue's policy. A valid AWS policy. For more information
	//    about policy structure, see Overview of AWS IAM Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/PoliciesOverview.html)
	//    in the Amazon IAM User Guide.
	//
	//    * ReceiveMessageWaitTimeSeconds - The length of time, in seconds, for
	//    which a ReceiveMessage action waits for a message to arrive. Valid values:
	//    an integer from 0 to 20 (seconds). Default: 0.
	//
	//    * RedrivePolicy - The string that includes the parameters for the dead-letter
	//    queue functionality of the source queue. For more information about the
	//    redrive policy and dead-letter queues, see Using Amazon SQS Dead-Letter
	//    Queues (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-dead-letter-queues.html)
	//    in the Amazon Simple Queue Service Developer Guide. deadLetterTargetArn
	//    - The Amazon Resource Name (ARN) of the dead-letter queue to which Amazon
	//    SQS moves messages after the value of maxReceiveCount is exceeded. maxReceiveCount
	//    - The number of times a message is delivered to the source queue before
	//    being moved to the dead-letter queue. When the ReceiveCount for a message
	//    exceeds the maxReceiveCount for a queue, Amazon SQS moves the message
	//    to the dead-letter-queue. The dead-letter queue of a FIFO queue must also
	//    be a FIFO queue. Similarly, the dead-letter queue of a standard queue
	//    must also be a standard queue.
	//
	//    * VisibilityTimeout - The visibility timeout for the queue, in seconds.
	//    Valid values: an integer from 0 to 43,200 (12 hours). Default: 30. For
	//    more information about the visibility timeout, see Visibility Timeout
	//    (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-visibility-timeout.html)
	//    in the Amazon Simple Queue Service Developer Guide.
	//
	// The following attributes apply only to server-side-encryption (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html):
	//
	//    * KmsMasterKeyId - The ID of an AWS-managed customer master key (CMK)
	//    for Amazon SQS or a custom CMK. For more information, see Key Terms (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-sse-key-terms).
	//    While the alias of the AWS-managed CMK for Amazon SQS is always alias/aws/sqs,
	//    the alias of a custom CMK can, for example, be alias/MyAlias . For more
	//    examples, see KeyId (https://docs.aws.amazon.com/kms/latest/APIReference/API_DescribeKey.html#API_DescribeKey_RequestParameters)
	//    in the AWS Key Management Service API Reference.
	//
	//    * KmsDataKeyReusePeriodSeconds - The length of time, in seconds, for which
	//    Amazon SQS can reuse a data key (https://docs.aws.amazon.com/kms/latest/developerguide/concepts.html#data-keys)
	//    to encrypt or decrypt messages before calling AWS KMS again. An integer
	//    representing seconds, between 60 seconds (1 minute) and 86,400 seconds
	//    (24 hours). Default: 300 (5 minutes). A shorter time period provides better
	//    security but results in more calls to KMS which might incur charges after
	//    Free Tier. For more information, see How Does the Data Key Reuse Period
	//    Work? (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/sqs-server-side-encryption.html#sqs-how-does-the-data-key-reuse-period-work).
	//
	// The following attribute applies only to FIFO (first-in-first-out) queues
	// (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html):
	//
	//    * ContentBasedDeduplication - Enables content-based deduplication. For
	//    more information, see Exactly-Once Processing (https://docs.aws.amazon.com/AWSSimpleQueueService/latest/SQSDeveloperGuide/FIFO-queues.html#FIFO-queues-exactly-once-processing)
	//    in the Amazon Simple Queue Service Developer Guide. Every message must
	//    have a unique MessageDeduplicationId, You may provide a MessageDeduplicationId
	//    explicitly. If you aren't able to provide a MessageDeduplicationId and
	//    you enable ContentBasedDeduplication for your queue, Amazon SQS uses a
	//    SHA-256 hash to generate the MessageDeduplicationId using the body of
	//    the message (but not the attributes of the message). If you don't provide
	//    a MessageDeduplicationId and the queue doesn't have ContentBasedDeduplication
	//    set, the action fails with an error. If the queue has ContentBasedDeduplication
	//    set, your MessageDeduplicationId overrides the generated one. When ContentBasedDeduplication
	//    is in effect, messages with identical content sent within the deduplication
	//    interval are treated as duplicates and only one copy of the message is
	//    delivered. If you send one message with ContentBasedDeduplication enabled
	//    and then another message with a MessageDeduplicationId that is the same
	//    as the one generated for the first MessageDeduplicationId, the two messages
	//    are treated as duplicates and only one copy of the message is delivered.
	//
	// Attributes is a required field
	Attributes map[string]*string `locationName:"Attribute" locationNameKey:"Name" locationNameValue:"Value" type:"map" flattened:"true" required:"true"`

	// The URL of the Amazon SQS queue whose attributes are set.
	//
	// Queue URLs and names are case-sensitive.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SetQueueAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SetQueueAttributesInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetQueueAttributesInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "SetQueueAttributesInput"}
	if s.Attributes == nil {
		invalidParams.Add(request.NewErrParamRequired("Attributes"))
	}
	if s.QueueUrl == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueUrl"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetAttributes sets the Attributes field's value.
func (s *SetQueueAttributesInput) SetAttributes(v map[string]*string) *SetQueueAttributesInput {
	s.Attributes = v
	return s
}

// SetQueueUrl sets the QueueUrl field's value.
func (s *SetQueueAttributesInput) SetQueueUrl(v string) *SetQueueAttributesInput {
	s.QueueUrl = &v
	return s
}

type SetQueueAttributesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetQueueAttributesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s SetQueueAttributesOutput) GoString() string {
	return s.String()
}

type TagQueueInput struct {
	_ struct{} `type:"structure"`

	// The URL of the queue.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`

	// The list of tags to be added to the specified queue.
	//
	// Tags is a required field
	Tags map[string]*string `locationName:"Tag" locationNameKey:"Key" locationNameValue:"Value" type:"map" flattened:"true" required:"true"`
}

// String returns the string representation
func (s TagQueueInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s TagQueueInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *TagQueueInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "TagQueueInput"}
	if s.QueueUrl == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueUrl"))
	}
	if s.Tags == nil {
		invalidParams.Add(request.NewErrParamRequired("Tags"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetQueueUrl sets the QueueUrl field's value.
func (s *TagQueueInput) SetQueueUrl(v string) *TagQueueInput {
	s.QueueUrl = &v
	return s
}

// SetTags sets the Tags field's value.
func (s *TagQueueInput) SetTags(v map[string]*string) *TagQueueInput {
	s.Tags = v
	return s
}

type TagQueueOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s TagQueueOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s TagQueueOutput) GoString() string {
	return s.String()
}

type UntagQueueInput struct {
	_ struct{} `type:"structure"`

	// The URL of the queue.
	//
	// QueueUrl is a required field
	QueueUrl *string `type:"string" required:"true"`

	// The list of tags to be removed from the specified queue.
	//
	// TagKeys is a required field
	TagKeys []*string `locationNameList:"TagKey" type:"list" flattened:"true" required:"true"`
}

// String returns the string representation
func (s UntagQueueInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UntagQueueInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UntagQueueInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UntagQueueInput"}
	if s.QueueUrl == nil {
		invalidParams.Add(request.NewErrParamRequired("QueueUrl"))
	}
	if s.TagKeys == nil {
		invalidParams.Add(request.NewErrParamRequired("TagKeys"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetQueueUrl sets the QueueUrl field's value.
func (s *UntagQueueInput) SetQueueUrl(v string) *UntagQueueInput {
	s.QueueUrl = &v
	return s
}

// SetTagKeys sets the TagKeys field's value.
func (s *UntagQueueInput) SetTagKeys(v []*string) *UntagQueueInput {
	s.TagKeys = v
	return s
}

type UntagQueueOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UntagQueueOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UntagQueueOutput) GoString() string {
	return s.String()
}

const (
	// MessageSystemAttributeNameSenderId is a MessageSystemAttributeName enum value
	MessageSystemAttributeNameSenderId = "SenderId"

	// MessageSystemAttributeNameSentTimestamp is a MessageSystemAttributeName enum value
	MessageSystemAttributeNameSentTimestamp = "SentTimestamp"

	// MessageSystemAttributeNameApproximateReceiveCount is a MessageSystemAttributeName enum value
	MessageSystemAttributeNameApproximateReceiveCount = "ApproximateReceiveCount"

	// MessageSystemAttributeNameApproximateFirstReceiveTimestamp is a MessageSystemAttributeName enum value
	MessageSystemAttributeNameApproximateFirstReceiveTimestamp = "ApproximateFirstReceiveTimestamp"

	// MessageSystemAttributeNameSequenceNumber is a MessageSystemAttributeName enum value
	MessageSystemAttributeNameSequenceNumber = "SequenceNumber"

	// MessageSystemAttributeNameMessageDeduplicationId is a MessageSystemAttributeName enum value
	MessageSystemAttributeNameMessageDeduplicationId = "MessageDeduplicationId"

	// MessageSystemAttributeNameMessageGroupId is a MessageSystemAttributeName enum value
	MessageSystemAttributeNameMessageGroupId = "MessageGroupId"

	// MessageSystemAttributeNameAwstraceHeader is a MessageSystemAttributeName enum value
	MessageSystemAttributeNameAwstraceHeader = "AWSTraceHeader"
)

const (
	// MessageSystemAttributeNameForSendsAwstraceHeader is a MessageSystemAttributeNameForSends enum value
	MessageSystemAttributeNameForSendsAwstraceHeader = "AWSTraceHeader"
)

const (
	// QueueAttributeNameAll is a QueueAttributeName enum value
	QueueAttributeNameAll = "All"

	// QueueAttributeNamePolicy is a QueueAttributeName enum value
	QueueAttributeNamePolicy = "Policy"

	// QueueAttributeNameVisibilityTimeout is a QueueAttributeName enum value
	QueueAttributeNameVisibilityTimeout = "VisibilityTimeout"

	// QueueAttributeNameMaximumMessageSize is a QueueAttributeName enum value
	QueueAttributeNameMaximumMessageSize = "MaximumMessageSize"

	// QueueAttributeNameMessageRetentionPeriod is a QueueAttributeName enum value
	QueueAttributeNameMessageRetentionPeriod = "MessageRetentionPeriod"

	// QueueAttributeNameApproximateNumberOfMessages is a QueueAttributeName enum value
	QueueAttributeNameApproximateNumberOfMessages = "ApproximateNumberOfMessages"

	// QueueAttributeNameApproximateNumberOfMessagesNotVisible is a QueueAttributeName enum value
	QueueAttributeNameApproximateNumberOfMessagesNotVisible = "ApproximateNumberOfMessagesNotVisible"

	// QueueAttributeNameCreatedTimestamp is a QueueAttributeName enum value
	QueueAttributeNameCreatedTimestamp = "CreatedTimestamp"

	// QueueAttributeNameLastModifiedTimestamp is a QueueAttributeName enum value
	QueueAttributeNameLastModifiedTimestamp = "LastModifiedTimestamp"

	// QueueAttributeNameQueueArn is a QueueAttributeName enum value
	QueueAttributeNameQueueArn = "QueueArn"

	// QueueAttributeNameApproximateNumberOfMessagesDelayed is a QueueAttributeName enum value
	QueueAttributeNameApproximateNumberOfMessagesDelayed = "ApproximateNumberOfMessagesDelayed"

	// QueueAttributeNameDelaySeconds is a QueueAttributeName enum value
	QueueAttributeNameDelaySeconds = "DelaySeconds"

	// QueueAttributeNameReceiveMessageWaitTimeSeconds is a QueueAttributeName enum value
	QueueAttributeNameReceiveMessageWaitTimeSeconds = "ReceiveMessageWaitTimeSeconds"

	// QueueAttributeNameRedrivePolicy is a QueueAttributeName enum value
	QueueAttributeNameRedrivePolicy = "RedrivePolicy"

	// QueueAttributeNameFifoQueue is a QueueAttributeName enum value
	QueueAttributeNameFifoQueue = "FifoQueue"

	// QueueAttributeNameContentBasedDeduplication is a QueueAttributeName enum value
	QueueAttributeNameContentBasedDeduplication = "ContentBasedDeduplication"

	// QueueAttributeNameKmsMasterKeyId is a QueueAttributeName enum value
	QueueAttributeNameKmsMasterKeyId = "KmsMasterKeyId"

	// QueueAttributeNameKmsDataKeyReusePeriodSeconds is a QueueAttributeName enum value
	QueueAttributeNameKmsDataKeyReusePeriodSeconds = "KmsDataKeyReusePeriodSeconds"
)
