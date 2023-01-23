/*
Copyright 2023 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/rancher/fleet/pkg/apis/fleet.cattle.io/v1alpha1"
	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/generic"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type GitRepoRestrictionHandler func(string, *v1alpha1.GitRepoRestriction) (*v1alpha1.GitRepoRestriction, error)

type GitRepoRestrictionController interface {
	generic.ControllerMeta
	GitRepoRestrictionClient

	OnChange(ctx context.Context, name string, sync GitRepoRestrictionHandler)
	OnRemove(ctx context.Context, name string, sync GitRepoRestrictionHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() GitRepoRestrictionCache
}

type GitRepoRestrictionClient interface {
	Create(*v1alpha1.GitRepoRestriction) (*v1alpha1.GitRepoRestriction, error)
	Update(*v1alpha1.GitRepoRestriction) (*v1alpha1.GitRepoRestriction, error)

	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1alpha1.GitRepoRestriction, error)
	List(namespace string, opts metav1.ListOptions) (*v1alpha1.GitRepoRestrictionList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GitRepoRestriction, err error)
}

type GitRepoRestrictionCache interface {
	Get(namespace, name string) (*v1alpha1.GitRepoRestriction, error)
	List(namespace string, selector labels.Selector) ([]*v1alpha1.GitRepoRestriction, error)

	AddIndexer(indexName string, indexer GitRepoRestrictionIndexer)
	GetByIndex(indexName, key string) ([]*v1alpha1.GitRepoRestriction, error)
}

type GitRepoRestrictionIndexer func(obj *v1alpha1.GitRepoRestriction) ([]string, error)

type gitRepoRestrictionController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewGitRepoRestrictionController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) GitRepoRestrictionController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &gitRepoRestrictionController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromGitRepoRestrictionHandlerToHandler(sync GitRepoRestrictionHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1alpha1.GitRepoRestriction
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1alpha1.GitRepoRestriction))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *gitRepoRestrictionController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1alpha1.GitRepoRestriction))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateGitRepoRestrictionDeepCopyOnChange(client GitRepoRestrictionClient, obj *v1alpha1.GitRepoRestriction, handler func(obj *v1alpha1.GitRepoRestriction) (*v1alpha1.GitRepoRestriction, error)) (*v1alpha1.GitRepoRestriction, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *gitRepoRestrictionController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *gitRepoRestrictionController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *gitRepoRestrictionController) OnChange(ctx context.Context, name string, sync GitRepoRestrictionHandler) {
	c.AddGenericHandler(ctx, name, FromGitRepoRestrictionHandlerToHandler(sync))
}

func (c *gitRepoRestrictionController) OnRemove(ctx context.Context, name string, sync GitRepoRestrictionHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromGitRepoRestrictionHandlerToHandler(sync)))
}

func (c *gitRepoRestrictionController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *gitRepoRestrictionController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *gitRepoRestrictionController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *gitRepoRestrictionController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *gitRepoRestrictionController) Cache() GitRepoRestrictionCache {
	return &gitRepoRestrictionCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *gitRepoRestrictionController) Create(obj *v1alpha1.GitRepoRestriction) (*v1alpha1.GitRepoRestriction, error) {
	result := &v1alpha1.GitRepoRestriction{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *gitRepoRestrictionController) Update(obj *v1alpha1.GitRepoRestriction) (*v1alpha1.GitRepoRestriction, error) {
	result := &v1alpha1.GitRepoRestriction{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *gitRepoRestrictionController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *gitRepoRestrictionController) Get(namespace, name string, options metav1.GetOptions) (*v1alpha1.GitRepoRestriction, error) {
	result := &v1alpha1.GitRepoRestriction{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *gitRepoRestrictionController) List(namespace string, opts metav1.ListOptions) (*v1alpha1.GitRepoRestrictionList, error) {
	result := &v1alpha1.GitRepoRestrictionList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *gitRepoRestrictionController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *gitRepoRestrictionController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1alpha1.GitRepoRestriction, error) {
	result := &v1alpha1.GitRepoRestriction{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type gitRepoRestrictionCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *gitRepoRestrictionCache) Get(namespace, name string) (*v1alpha1.GitRepoRestriction, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1alpha1.GitRepoRestriction), nil
}

func (c *gitRepoRestrictionCache) List(namespace string, selector labels.Selector) (ret []*v1alpha1.GitRepoRestriction, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.GitRepoRestriction))
	})

	return ret, err
}

func (c *gitRepoRestrictionCache) AddIndexer(indexName string, indexer GitRepoRestrictionIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1alpha1.GitRepoRestriction))
		},
	}))
}

func (c *gitRepoRestrictionCache) GetByIndex(indexName, key string) (result []*v1alpha1.GitRepoRestriction, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1alpha1.GitRepoRestriction, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1alpha1.GitRepoRestriction))
	}
	return result, nil
}
