// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	applicationv1alpha1 "github.com/dcoppa/argo-cd/v2/pkg/apis/application/v1alpha1"
	versioned "github.com/dcoppa/argo-cd/v2/pkg/client/clientset/versioned"
	internalinterfaces "github.com/dcoppa/argo-cd/v2/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/dcoppa/argo-cd/v2/pkg/client/listers/application/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ApplicationInformer provides access to a shared informer and lister for
// Applications.
type ApplicationInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ApplicationLister
}

type applicationInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewApplicationInformer constructs a new informer for Application type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewApplicationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredApplicationInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredApplicationInformer constructs a new informer for Application type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredApplicationInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArgoprojV1alpha1().Applications(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ArgoprojV1alpha1().Applications(namespace).Watch(context.TODO(), options)
			},
		},
		&applicationv1alpha1.Application{},
		resyncPeriod,
		indexers,
	)
}

func (f *applicationInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredApplicationInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *applicationInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&applicationv1alpha1.Application{}, f.defaultInformer)
}

func (f *applicationInformer) Lister() v1alpha1.ApplicationLister {
	return v1alpha1.NewApplicationLister(f.Informer().GetIndexer())
}
