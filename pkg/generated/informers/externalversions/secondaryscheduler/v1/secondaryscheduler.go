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

package v1

import (
	"context"
	time "time"

	secondaryschedulerv1 "github.com/openshift/secondary-scheduler-operator/pkg/apis/secondaryscheduler/v1"
	versioned "github.com/openshift/secondary-scheduler-operator/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/openshift/secondary-scheduler-operator/pkg/generated/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/secondary-scheduler-operator/pkg/generated/listers/secondaryscheduler/v1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// SecondarySchedulerInformer provides access to a shared informer and lister for
// SecondarySchedulers.
type SecondarySchedulerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.SecondarySchedulerLister
}

type secondarySchedulerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewSecondarySchedulerInformer constructs a new informer for SecondaryScheduler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewSecondarySchedulerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredSecondarySchedulerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredSecondarySchedulerInformer constructs a new informer for SecondaryScheduler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredSecondarySchedulerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecondaryschedulersV1().SecondarySchedulers(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.SecondaryschedulersV1().SecondarySchedulers(namespace).Watch(context.TODO(), options)
			},
		},
		&secondaryschedulerv1.SecondaryScheduler{},
		resyncPeriod,
		indexers,
	)
}

func (f *secondarySchedulerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredSecondarySchedulerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *secondarySchedulerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&secondaryschedulerv1.SecondaryScheduler{}, f.defaultInformer)
}

func (f *secondarySchedulerInformer) Lister() v1.SecondarySchedulerLister {
	return v1.NewSecondarySchedulerLister(f.Informer().GetIndexer())
}
