// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	settings_v1alpha1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/settings/v1alpha1"
	versioned "github.com/kubernetes-incubator/service-catalog/pkg/client/clientset/versioned"
	internalinterfaces "github.com/kubernetes-incubator/service-catalog/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/kubernetes-incubator/service-catalog/pkg/client/listers/settings/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PodPresetInformer provides access to a shared informer and lister for
// PodPresets.
type PodPresetInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PodPresetLister
}

type podPresetInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPodPresetInformer constructs a new informer for PodPreset type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPodPresetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPodPresetInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPodPresetInformer constructs a new informer for PodPreset type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPodPresetInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SettingsV1alpha1().PodPresets(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SettingsV1alpha1().PodPresets(namespace).Watch(options)
			},
		},
		&settings_v1alpha1.PodPreset{},
		resyncPeriod,
		indexers,
	)
}

func (f *podPresetInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPodPresetInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *podPresetInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&settings_v1alpha1.PodPreset{}, f.defaultInformer)
}

func (f *podPresetInformer) Lister() v1alpha1.PodPresetLister {
	return v1alpha1.NewPodPresetLister(f.Informer().GetIndexer())
}
