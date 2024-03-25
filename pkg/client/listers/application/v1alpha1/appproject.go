// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/dcoppa/argo-cd/v2/pkg/apis/application/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AppProjectLister helps list AppProjects.
// All objects returned here must be treated as read-only.
type AppProjectLister interface {
	// List lists all AppProjects in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AppProject, err error)
	// AppProjects returns an object that can list and get AppProjects.
	AppProjects(namespace string) AppProjectNamespaceLister
	AppProjectListerExpansion
}

// appProjectLister implements the AppProjectLister interface.
type appProjectLister struct {
	indexer cache.Indexer
}

// NewAppProjectLister returns a new AppProjectLister.
func NewAppProjectLister(indexer cache.Indexer) AppProjectLister {
	return &appProjectLister{indexer: indexer}
}

// List lists all AppProjects in the indexer.
func (s *appProjectLister) List(selector labels.Selector) (ret []*v1alpha1.AppProject, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppProject))
	})
	return ret, err
}

// AppProjects returns an object that can list and get AppProjects.
func (s *appProjectLister) AppProjects(namespace string) AppProjectNamespaceLister {
	return appProjectNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppProjectNamespaceLister helps list and get AppProjects.
// All objects returned here must be treated as read-only.
type AppProjectNamespaceLister interface {
	// List lists all AppProjects in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AppProject, err error)
	// Get retrieves the AppProject from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AppProject, error)
	AppProjectNamespaceListerExpansion
}

// appProjectNamespaceLister implements the AppProjectNamespaceLister
// interface.
type appProjectNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppProjects in the indexer for a given namespace.
func (s appProjectNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AppProject, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppProject))
	})
	return ret, err
}

// Get retrieves the AppProject from the indexer for a given namespace and name.
func (s appProjectNamespaceLister) Get(name string) (*v1alpha1.AppProject, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("appproject"), name)
	}
	return obj.(*v1alpha1.AppProject), nil
}
