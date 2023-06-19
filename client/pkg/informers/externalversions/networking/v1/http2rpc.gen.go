// Copyright (c) 2022 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	networkingv1 "github.com/alibaba/higress/client/pkg/apis/networking/v1"
	versioned "github.com/alibaba/higress/client/pkg/clientset/versioned"
	internalinterfaces "github.com/alibaba/higress/client/pkg/informers/externalversions/internalinterfaces"
	v1 "github.com/alibaba/higress/client/pkg/listers/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// Http2RpcInformer provides access to a shared informer and lister for
// Http2Rpcs.
type Http2RpcInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.Http2RpcLister
}

type http2RpcInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHttp2RpcInformer constructs a new informer for Http2Rpc type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHttp2RpcInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHttp2RpcInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHttp2RpcInformer constructs a new informer for Http2Rpc type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHttp2RpcInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1().Http2Rpcs(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.NetworkingV1().Http2Rpcs(namespace).Watch(context.TODO(), options)
			},
		},
		&networkingv1.Http2Rpc{},
		resyncPeriod,
		indexers,
	)
}

func (f *http2RpcInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHttp2RpcInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *http2RpcInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&networkingv1.Http2Rpc{}, f.defaultInformer)
}

func (f *http2RpcInformer) Lister() v1.Http2RpcLister {
	return v1.NewHttp2RpcLister(f.Informer().GetIndexer())
}