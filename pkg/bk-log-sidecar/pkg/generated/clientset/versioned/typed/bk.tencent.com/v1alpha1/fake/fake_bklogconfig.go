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
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1alpha1 "github.com/TencentBlueKing/bk-log-sidecar/api/bk.tencent.com/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBkLogConfigs implements BkLogConfigInterface
type FakeBkLogConfigs struct {
	Fake *FakeBkV1alpha1
	ns   string
}

var bklogconfigsResource = schema.GroupVersionResource{Group: "bk.tencent.com", Version: "v1alpha1", Resource: "bklogconfigs"}

var bklogconfigsKind = schema.GroupVersionKind{Group: "bk.tencent.com", Version: "v1alpha1", Kind: "BkLogConfig"}

// Get takes name of the bkLogConfig, and returns the corresponding bkLogConfig object, and an error if there is any.
func (c *FakeBkLogConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.BkLogConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bklogconfigsResource, c.ns, name), &v1alpha1.BkLogConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BkLogConfig), err
}

// List takes label and field selectors, and returns the list of BkLogConfigs that match those selectors.
func (c *FakeBkLogConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.BkLogConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bklogconfigsResource, bklogconfigsKind, c.ns, opts), &v1alpha1.BkLogConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.BkLogConfigList{ListMeta: obj.(*v1alpha1.BkLogConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.BkLogConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bkLogConfigs.
func (c *FakeBkLogConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bklogconfigsResource, c.ns, opts))

}

// Create takes the representation of a bkLogConfig and creates it.  Returns the server's representation of the bkLogConfig, and an error, if there is any.
func (c *FakeBkLogConfigs) Create(ctx context.Context, bkLogConfig *v1alpha1.BkLogConfig, opts v1.CreateOptions) (result *v1alpha1.BkLogConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bklogconfigsResource, c.ns, bkLogConfig), &v1alpha1.BkLogConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BkLogConfig), err
}

// Update takes the representation of a bkLogConfig and updates it. Returns the server's representation of the bkLogConfig, and an error, if there is any.
func (c *FakeBkLogConfigs) Update(ctx context.Context, bkLogConfig *v1alpha1.BkLogConfig, opts v1.UpdateOptions) (result *v1alpha1.BkLogConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bklogconfigsResource, c.ns, bkLogConfig), &v1alpha1.BkLogConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BkLogConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBkLogConfigs) UpdateStatus(ctx context.Context, bkLogConfig *v1alpha1.BkLogConfig, opts v1.UpdateOptions) (*v1alpha1.BkLogConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(bklogconfigsResource, "status", c.ns, bkLogConfig), &v1alpha1.BkLogConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BkLogConfig), err
}

// Delete takes name of the bkLogConfig and deletes it. Returns an error if one occurs.
func (c *FakeBkLogConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(bklogconfigsResource, c.ns, name), &v1alpha1.BkLogConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBkLogConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bklogconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.BkLogConfigList{})
	return err
}

// Patch applies the patch and returns the patched bkLogConfig.
func (c *FakeBkLogConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.BkLogConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bklogconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.BkLogConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.BkLogConfig), err
}
