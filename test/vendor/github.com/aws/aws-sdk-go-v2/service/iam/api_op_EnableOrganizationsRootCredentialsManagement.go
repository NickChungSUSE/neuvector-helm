// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables the management of privileged root user credentials across member
// accounts in your organization. When you enable root credentials management for [centralized root access]
// , the management account and the delegated admininstrator for IAM can manage
// root user credentials for member accounts in your organization.
//
// Before you enable centralized root access, you must have an account configured
// with the following settings:
//
//   - You must manage your Amazon Web Services accounts in [Organizations].
//
//   - Enable trusted access for Identity and Access Management in Organizations.
//     For details, see [IAM and Organizations]in the Organizations User Guide.
//
// [Organizations]: https://docs.aws.amazon.com/organizations/latest/userguide/orgs_introduction.html
// [centralized root access]: https://docs.aws.amazon.com/IAM/latest/UserGuide/id_root-user.html#id_root-user-access-management
// [IAM and Organizations]: https://docs.aws.amazon.com/organizations/latest/userguide/services-that-can-integrate-ra.html
func (c *Client) EnableOrganizationsRootCredentialsManagement(ctx context.Context, params *EnableOrganizationsRootCredentialsManagementInput, optFns ...func(*Options)) (*EnableOrganizationsRootCredentialsManagementOutput, error) {
	if params == nil {
		params = &EnableOrganizationsRootCredentialsManagementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "EnableOrganizationsRootCredentialsManagement", params, optFns, c.addOperationEnableOrganizationsRootCredentialsManagementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*EnableOrganizationsRootCredentialsManagementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableOrganizationsRootCredentialsManagementInput struct {
	noSmithyDocumentSerde
}

type EnableOrganizationsRootCredentialsManagementOutput struct {

	// The features you have enabled for centralized root access.
	EnabledFeatures []types.FeatureType

	// The unique identifier (ID) of an organization.
	OrganizationId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationEnableOrganizationsRootCredentialsManagementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpEnableOrganizationsRootCredentialsManagement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpEnableOrganizationsRootCredentialsManagement{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "EnableOrganizationsRootCredentialsManagement"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addSpanRetryLoop(stack, options); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opEnableOrganizationsRootCredentialsManagement(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	if err = addSpanInitializeStart(stack); err != nil {
		return err
	}
	if err = addSpanInitializeEnd(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestStart(stack); err != nil {
		return err
	}
	if err = addSpanBuildRequestEnd(stack); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opEnableOrganizationsRootCredentialsManagement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "EnableOrganizationsRootCredentialsManagement",
	}
}