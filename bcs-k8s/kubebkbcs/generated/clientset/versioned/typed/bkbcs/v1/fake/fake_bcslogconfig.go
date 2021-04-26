/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	bkbcsv1 "github.com/Tencent/bk-bcs/bcs-k8s/kubebkbcs/apis/bkbcs/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBcsLogConfigs implements BcsLogConfigInterface
type FakeBcsLogConfigs struct {
	Fake *FakeBkbcsV1
	ns   string
}

var bcslogconfigsResource = schema.GroupVersionResource{Group: "bkbcs", Version: "v1", Resource: "bcslogconfigs"}

var bcslogconfigsKind = schema.GroupVersionKind{Group: "bkbcs", Version: "v1", Kind: "BcsLogConfig"}

// Get takes name of the bcsLogConfig, and returns the corresponding bcsLogConfig object, and an error if there is any.
func (c *FakeBcsLogConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *bkbcsv1.BcsLogConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(bcslogconfigsResource, c.ns, name), &bkbcsv1.BcsLogConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*bkbcsv1.BcsLogConfig), err
}

// List takes label and field selectors, and returns the list of BcsLogConfigs that match those selectors.
func (c *FakeBcsLogConfigs) List(ctx context.Context, opts v1.ListOptions) (result *bkbcsv1.BcsLogConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(bcslogconfigsResource, bcslogconfigsKind, c.ns, opts), &bkbcsv1.BcsLogConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &bkbcsv1.BcsLogConfigList{ListMeta: obj.(*bkbcsv1.BcsLogConfigList).ListMeta}
	for _, item := range obj.(*bkbcsv1.BcsLogConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested bcsLogConfigs.
func (c *FakeBcsLogConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(bcslogconfigsResource, c.ns, opts))

}

// Create takes the representation of a bcsLogConfig and creates it.  Returns the server's representation of the bcsLogConfig, and an error, if there is any.
func (c *FakeBcsLogConfigs) Create(ctx context.Context, bcsLogConfig *bkbcsv1.BcsLogConfig, opts v1.CreateOptions) (result *bkbcsv1.BcsLogConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(bcslogconfigsResource, c.ns, bcsLogConfig), &bkbcsv1.BcsLogConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*bkbcsv1.BcsLogConfig), err
}

// Update takes the representation of a bcsLogConfig and updates it. Returns the server's representation of the bcsLogConfig, and an error, if there is any.
func (c *FakeBcsLogConfigs) Update(ctx context.Context, bcsLogConfig *bkbcsv1.BcsLogConfig, opts v1.UpdateOptions) (result *bkbcsv1.BcsLogConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(bcslogconfigsResource, c.ns, bcsLogConfig), &bkbcsv1.BcsLogConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*bkbcsv1.BcsLogConfig), err
}

// Delete takes name of the bcsLogConfig and deletes it. Returns an error if one occurs.
func (c *FakeBcsLogConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(bcslogconfigsResource, c.ns, name), &bkbcsv1.BcsLogConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBcsLogConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(bcslogconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &bkbcsv1.BcsLogConfigList{})
	return err
}

// Patch applies the patch and returns the patched bcsLogConfig.
func (c *FakeBcsLogConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *bkbcsv1.BcsLogConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(bcslogconfigsResource, c.ns, name, pt, data, subresources...), &bkbcsv1.BcsLogConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*bkbcsv1.BcsLogConfig), err
}