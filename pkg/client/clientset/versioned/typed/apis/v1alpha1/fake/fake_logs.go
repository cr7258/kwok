/*
Copyright 2023 The Kubernetes Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "sigs.k8s.io/kwok/pkg/apis/v1alpha1"
)

// FakeLogs implements LogsInterface
type FakeLogs struct {
	Fake *FakeKwokV1alpha1
	ns   string
}

var logsResource = v1alpha1.SchemeGroupVersion.WithResource("logs")

var logsKind = v1alpha1.SchemeGroupVersion.WithKind("Logs")

// Get takes name of the logs, and returns the corresponding logs object, and an error if there is any.
func (c *FakeLogs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.Logs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(logsResource, c.ns, name), &v1alpha1.Logs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Logs), err
}

// List takes label and field selectors, and returns the list of Logs that match those selectors.
func (c *FakeLogs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.LogsList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(logsResource, logsKind, c.ns, opts), &v1alpha1.LogsList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.LogsList{ListMeta: obj.(*v1alpha1.LogsList).ListMeta}
	for _, item := range obj.(*v1alpha1.LogsList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested logs.
func (c *FakeLogs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(logsResource, c.ns, opts))

}

// Create takes the representation of a logs and creates it.  Returns the server's representation of the logs, and an error, if there is any.
func (c *FakeLogs) Create(ctx context.Context, logs *v1alpha1.Logs, opts v1.CreateOptions) (result *v1alpha1.Logs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(logsResource, c.ns, logs), &v1alpha1.Logs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Logs), err
}

// Update takes the representation of a logs and updates it. Returns the server's representation of the logs, and an error, if there is any.
func (c *FakeLogs) Update(ctx context.Context, logs *v1alpha1.Logs, opts v1.UpdateOptions) (result *v1alpha1.Logs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(logsResource, c.ns, logs), &v1alpha1.Logs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Logs), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeLogs) UpdateStatus(ctx context.Context, logs *v1alpha1.Logs, opts v1.UpdateOptions) (*v1alpha1.Logs, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(logsResource, "status", c.ns, logs), &v1alpha1.Logs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Logs), err
}

// Delete takes name of the logs and deletes it. Returns an error if one occurs.
func (c *FakeLogs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(logsResource, c.ns, name, opts), &v1alpha1.Logs{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeLogs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(logsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.LogsList{})
	return err
}

// Patch applies the patch and returns the patched logs.
func (c *FakeLogs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.Logs, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(logsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Logs{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Logs), err
}
