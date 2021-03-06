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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/foghornci/foghorn/pkg/apis/foghorn.jenkins.io/v1alpha1"
	scheme "github.com/foghornci/foghorn/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GitEventsGetter has a method to return a GitEventInterface.
// A group's client should implement this interface.
type GitEventsGetter interface {
	GitEvents(namespace string) GitEventInterface
}

// GitEventInterface has methods to work with GitEvent resources.
type GitEventInterface interface {
	Create(*v1alpha1.GitEvent) (*v1alpha1.GitEvent, error)
	Update(*v1alpha1.GitEvent) (*v1alpha1.GitEvent, error)
	UpdateStatus(*v1alpha1.GitEvent) (*v1alpha1.GitEvent, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GitEvent, error)
	List(opts v1.ListOptions) (*v1alpha1.GitEventList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GitEvent, err error)
	GitEventExpansion
}

// gitEvents implements GitEventInterface
type gitEvents struct {
	client rest.Interface
	ns     string
}

// newGitEvents returns a GitEvents
func newGitEvents(c *FoghornV1alpha1Client, namespace string) *gitEvents {
	return &gitEvents{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gitEvent, and returns the corresponding gitEvent object, and an error if there is any.
func (c *gitEvents) Get(name string, options v1.GetOptions) (result *v1alpha1.GitEvent, err error) {
	result = &v1alpha1.GitEvent{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gitevents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GitEvents that match those selectors.
func (c *gitEvents) List(opts v1.ListOptions) (result *v1alpha1.GitEventList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.GitEventList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gitevents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gitEvents.
func (c *gitEvents) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gitevents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a gitEvent and creates it.  Returns the server's representation of the gitEvent, and an error, if there is any.
func (c *gitEvents) Create(gitEvent *v1alpha1.GitEvent) (result *v1alpha1.GitEvent, err error) {
	result = &v1alpha1.GitEvent{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gitevents").
		Body(gitEvent).
		Do().
		Into(result)
	return
}

// Update takes the representation of a gitEvent and updates it. Returns the server's representation of the gitEvent, and an error, if there is any.
func (c *gitEvents) Update(gitEvent *v1alpha1.GitEvent) (result *v1alpha1.GitEvent, err error) {
	result = &v1alpha1.GitEvent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gitevents").
		Name(gitEvent.Name).
		Body(gitEvent).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *gitEvents) UpdateStatus(gitEvent *v1alpha1.GitEvent) (result *v1alpha1.GitEvent, err error) {
	result = &v1alpha1.GitEvent{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gitevents").
		Name(gitEvent.Name).
		SubResource("status").
		Body(gitEvent).
		Do().
		Into(result)
	return
}

// Delete takes name of the gitEvent and deletes it. Returns an error if one occurs.
func (c *gitEvents) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gitevents").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gitEvents) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gitevents").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched gitEvent.
func (c *gitEvents) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GitEvent, err error) {
	result = &v1alpha1.GitEvent{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gitevents").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
