/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	zbmv1alpha1 "k8s.io/vk8scontroller/pkg/apis/zbm/v1alpha1"
	versioned "k8s.io/vk8scontroller/pkg/generated/clientset/versioned"
	internalinterfaces "k8s.io/vk8scontroller/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "k8s.io/vk8scontroller/pkg/generated/listers/zbm/v1alpha1"
)

// ZbmInformer provides access to a shared informer and lister for
// Zbms.
type ZbmInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ZbmLister
}

type zbmInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewZbmInformer constructs a new informer for Zbm type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewZbmInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredZbmInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredZbmInformer constructs a new informer for Zbm type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredZbmInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IbmV1alpha1().Zbms(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.IbmV1alpha1().Zbms(namespace).Watch(context.TODO(), options)
			},
		},
		&zbmv1alpha1.Zbm{},
		resyncPeriod,
		indexers,
	)
}

func (f *zbmInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredZbmInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *zbmInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&zbmv1alpha1.Zbm{}, f.defaultInformer)
}

func (f *zbmInformer) Lister() v1alpha1.ZbmLister {
	return v1alpha1.NewZbmLister(f.Informer().GetIndexer())
}
