// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package managedgrafanaiface provides an interface to enable mocking the Amazon Managed Grafana service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package managedgrafanaiface

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/aws/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/aws/aws-sdk-go/service/managedgrafana"
)

// ManagedGrafanaAPI provides an interface to enable mocking the
// managedgrafana.ManagedGrafana service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Managed Grafana.
//    func myFunc(svc managedgrafanaiface.ManagedGrafanaAPI) bool {
//        // Make svc.AssociateLicense request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := managedgrafana.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockManagedGrafanaClient struct {
//        managedgrafanaiface.ManagedGrafanaAPI
//    }
//    func (m *mockManagedGrafanaClient) AssociateLicense(input *managedgrafana.AssociateLicenseInput) (*managedgrafana.AssociateLicenseOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockManagedGrafanaClient{}
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
type ManagedGrafanaAPI interface {
	AssociateLicense(*managedgrafana.AssociateLicenseInput) (*managedgrafana.AssociateLicenseOutput, error)
	AssociateLicenseWithContext(aws.Context, *managedgrafana.AssociateLicenseInput, ...request.Option) (*managedgrafana.AssociateLicenseOutput, error)
	AssociateLicenseRequest(*managedgrafana.AssociateLicenseInput) (*request.Request, *managedgrafana.AssociateLicenseOutput)

	CreateWorkspace(*managedgrafana.CreateWorkspaceInput) (*managedgrafana.CreateWorkspaceOutput, error)
	CreateWorkspaceWithContext(aws.Context, *managedgrafana.CreateWorkspaceInput, ...request.Option) (*managedgrafana.CreateWorkspaceOutput, error)
	CreateWorkspaceRequest(*managedgrafana.CreateWorkspaceInput) (*request.Request, *managedgrafana.CreateWorkspaceOutput)

	CreateWorkspaceApiKey(*managedgrafana.CreateWorkspaceApiKeyInput) (*managedgrafana.CreateWorkspaceApiKeyOutput, error)
	CreateWorkspaceApiKeyWithContext(aws.Context, *managedgrafana.CreateWorkspaceApiKeyInput, ...request.Option) (*managedgrafana.CreateWorkspaceApiKeyOutput, error)
	CreateWorkspaceApiKeyRequest(*managedgrafana.CreateWorkspaceApiKeyInput) (*request.Request, *managedgrafana.CreateWorkspaceApiKeyOutput)

	DeleteWorkspace(*managedgrafana.DeleteWorkspaceInput) (*managedgrafana.DeleteWorkspaceOutput, error)
	DeleteWorkspaceWithContext(aws.Context, *managedgrafana.DeleteWorkspaceInput, ...request.Option) (*managedgrafana.DeleteWorkspaceOutput, error)
	DeleteWorkspaceRequest(*managedgrafana.DeleteWorkspaceInput) (*request.Request, *managedgrafana.DeleteWorkspaceOutput)

	DeleteWorkspaceApiKey(*managedgrafana.DeleteWorkspaceApiKeyInput) (*managedgrafana.DeleteWorkspaceApiKeyOutput, error)
	DeleteWorkspaceApiKeyWithContext(aws.Context, *managedgrafana.DeleteWorkspaceApiKeyInput, ...request.Option) (*managedgrafana.DeleteWorkspaceApiKeyOutput, error)
	DeleteWorkspaceApiKeyRequest(*managedgrafana.DeleteWorkspaceApiKeyInput) (*request.Request, *managedgrafana.DeleteWorkspaceApiKeyOutput)

	DescribeWorkspace(*managedgrafana.DescribeWorkspaceInput) (*managedgrafana.DescribeWorkspaceOutput, error)
	DescribeWorkspaceWithContext(aws.Context, *managedgrafana.DescribeWorkspaceInput, ...request.Option) (*managedgrafana.DescribeWorkspaceOutput, error)
	DescribeWorkspaceRequest(*managedgrafana.DescribeWorkspaceInput) (*request.Request, *managedgrafana.DescribeWorkspaceOutput)

	DescribeWorkspaceAuthentication(*managedgrafana.DescribeWorkspaceAuthenticationInput) (*managedgrafana.DescribeWorkspaceAuthenticationOutput, error)
	DescribeWorkspaceAuthenticationWithContext(aws.Context, *managedgrafana.DescribeWorkspaceAuthenticationInput, ...request.Option) (*managedgrafana.DescribeWorkspaceAuthenticationOutput, error)
	DescribeWorkspaceAuthenticationRequest(*managedgrafana.DescribeWorkspaceAuthenticationInput) (*request.Request, *managedgrafana.DescribeWorkspaceAuthenticationOutput)

	DisassociateLicense(*managedgrafana.DisassociateLicenseInput) (*managedgrafana.DisassociateLicenseOutput, error)
	DisassociateLicenseWithContext(aws.Context, *managedgrafana.DisassociateLicenseInput, ...request.Option) (*managedgrafana.DisassociateLicenseOutput, error)
	DisassociateLicenseRequest(*managedgrafana.DisassociateLicenseInput) (*request.Request, *managedgrafana.DisassociateLicenseOutput)

	ListPermissions(*managedgrafana.ListPermissionsInput) (*managedgrafana.ListPermissionsOutput, error)
	ListPermissionsWithContext(aws.Context, *managedgrafana.ListPermissionsInput, ...request.Option) (*managedgrafana.ListPermissionsOutput, error)
	ListPermissionsRequest(*managedgrafana.ListPermissionsInput) (*request.Request, *managedgrafana.ListPermissionsOutput)

	ListPermissionsPages(*managedgrafana.ListPermissionsInput, func(*managedgrafana.ListPermissionsOutput, bool) bool) error
	ListPermissionsPagesWithContext(aws.Context, *managedgrafana.ListPermissionsInput, func(*managedgrafana.ListPermissionsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*managedgrafana.ListTagsForResourceInput) (*managedgrafana.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *managedgrafana.ListTagsForResourceInput, ...request.Option) (*managedgrafana.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*managedgrafana.ListTagsForResourceInput) (*request.Request, *managedgrafana.ListTagsForResourceOutput)

	ListWorkspaces(*managedgrafana.ListWorkspacesInput) (*managedgrafana.ListWorkspacesOutput, error)
	ListWorkspacesWithContext(aws.Context, *managedgrafana.ListWorkspacesInput, ...request.Option) (*managedgrafana.ListWorkspacesOutput, error)
	ListWorkspacesRequest(*managedgrafana.ListWorkspacesInput) (*request.Request, *managedgrafana.ListWorkspacesOutput)

	ListWorkspacesPages(*managedgrafana.ListWorkspacesInput, func(*managedgrafana.ListWorkspacesOutput, bool) bool) error
	ListWorkspacesPagesWithContext(aws.Context, *managedgrafana.ListWorkspacesInput, func(*managedgrafana.ListWorkspacesOutput, bool) bool, ...request.Option) error

	TagResource(*managedgrafana.TagResourceInput) (*managedgrafana.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *managedgrafana.TagResourceInput, ...request.Option) (*managedgrafana.TagResourceOutput, error)
	TagResourceRequest(*managedgrafana.TagResourceInput) (*request.Request, *managedgrafana.TagResourceOutput)

	UntagResource(*managedgrafana.UntagResourceInput) (*managedgrafana.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *managedgrafana.UntagResourceInput, ...request.Option) (*managedgrafana.UntagResourceOutput, error)
	UntagResourceRequest(*managedgrafana.UntagResourceInput) (*request.Request, *managedgrafana.UntagResourceOutput)

	UpdatePermissions(*managedgrafana.UpdatePermissionsInput) (*managedgrafana.UpdatePermissionsOutput, error)
	UpdatePermissionsWithContext(aws.Context, *managedgrafana.UpdatePermissionsInput, ...request.Option) (*managedgrafana.UpdatePermissionsOutput, error)
	UpdatePermissionsRequest(*managedgrafana.UpdatePermissionsInput) (*request.Request, *managedgrafana.UpdatePermissionsOutput)

	UpdateWorkspace(*managedgrafana.UpdateWorkspaceInput) (*managedgrafana.UpdateWorkspaceOutput, error)
	UpdateWorkspaceWithContext(aws.Context, *managedgrafana.UpdateWorkspaceInput, ...request.Option) (*managedgrafana.UpdateWorkspaceOutput, error)
	UpdateWorkspaceRequest(*managedgrafana.UpdateWorkspaceInput) (*request.Request, *managedgrafana.UpdateWorkspaceOutput)

	UpdateWorkspaceAuthentication(*managedgrafana.UpdateWorkspaceAuthenticationInput) (*managedgrafana.UpdateWorkspaceAuthenticationOutput, error)
	UpdateWorkspaceAuthenticationWithContext(aws.Context, *managedgrafana.UpdateWorkspaceAuthenticationInput, ...request.Option) (*managedgrafana.UpdateWorkspaceAuthenticationOutput, error)
	UpdateWorkspaceAuthenticationRequest(*managedgrafana.UpdateWorkspaceAuthenticationInput) (*request.Request, *managedgrafana.UpdateWorkspaceAuthenticationOutput)
}

var _ ManagedGrafanaAPI = (*managedgrafana.ManagedGrafana)(nil)
