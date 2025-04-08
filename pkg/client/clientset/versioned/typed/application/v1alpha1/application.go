// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	context "context"

	applicationv1alpha1 "github.com/argoproj/argo-cd/v3/pkg/apis/application/v1alpha1"
	scheme "github.com/argoproj/argo-cd/v3/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// ApplicationsGetter has a method to return a ApplicationInterface.
// A group's client should implement this interface.
type ApplicationsGetter interface {
	Applications(namespace string) ApplicationInterface
}

// ApplicationInterface has methods to work with Application resources.
type ApplicationInterface interface {
	Create(ctx context.Context, application *applicationv1alpha1.Application, opts v1.CreateOptions) (*applicationv1alpha1.Application, error)
	Update(ctx context.Context, application *applicationv1alpha1.Application, opts v1.UpdateOptions) (*applicationv1alpha1.Application, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*applicationv1alpha1.Application, error)
	List(ctx context.Context, opts v1.ListOptions) (*applicationv1alpha1.ApplicationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *applicationv1alpha1.Application, err error)
	ApplicationExpansion
}

// applications implements ApplicationInterface
type applications struct {
	*gentype.ClientWithList[*applicationv1alpha1.Application, *applicationv1alpha1.ApplicationList]
}

// newApplications returns a Applications
func newApplications(c *ArgoprojV1alpha1Client, namespace string) *applications {
	return &applications{
		gentype.NewClientWithList[*applicationv1alpha1.Application, *applicationv1alpha1.ApplicationList](
			"applications",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *applicationv1alpha1.Application { return &applicationv1alpha1.Application{} },
			func() *applicationv1alpha1.ApplicationList { return &applicationv1alpha1.ApplicationList{} },
		),
	}
}
