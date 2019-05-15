// Copyright 2018 The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
<<<<<<< HEAD
	v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	scheme "github.com/coreos/prometheus-operator/pkg/client/versioned/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
=======
	"time"

	v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	scheme "github.com/coreos/prometheus-operator/pkg/client/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
>>>>>>> v0.0.4
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ServiceMonitorsGetter has a method to return a ServiceMonitorInterface.
// A group's client should implement this interface.
type ServiceMonitorsGetter interface {
	ServiceMonitors(namespace string) ServiceMonitorInterface
}

// ServiceMonitorInterface has methods to work with ServiceMonitor resources.
type ServiceMonitorInterface interface {
	Create(*v1.ServiceMonitor) (*v1.ServiceMonitor, error)
	Update(*v1.ServiceMonitor) (*v1.ServiceMonitor, error)
<<<<<<< HEAD
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.ServiceMonitor, error)
	List(opts meta_v1.ListOptions) (*v1.ServiceMonitorList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
=======
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.ServiceMonitor, error)
	List(opts metav1.ListOptions) (*v1.ServiceMonitorList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
>>>>>>> v0.0.4
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ServiceMonitor, err error)
	ServiceMonitorExpansion
}

// serviceMonitors implements ServiceMonitorInterface
type serviceMonitors struct {
	client rest.Interface
	ns     string
}

// newServiceMonitors returns a ServiceMonitors
func newServiceMonitors(c *MonitoringV1Client, namespace string) *serviceMonitors {
	return &serviceMonitors{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the serviceMonitor, and returns the corresponding serviceMonitor object, and an error if there is any.
<<<<<<< HEAD
func (c *serviceMonitors) Get(name string, options meta_v1.GetOptions) (result *v1.ServiceMonitor, err error) {
=======
func (c *serviceMonitors) Get(name string, options metav1.GetOptions) (result *v1.ServiceMonitor, err error) {
>>>>>>> v0.0.4
	result = &v1.ServiceMonitor{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicemonitors").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ServiceMonitors that match those selectors.
<<<<<<< HEAD
func (c *serviceMonitors) List(opts meta_v1.ListOptions) (result *v1.ServiceMonitorList, err error) {
=======
func (c *serviceMonitors) List(opts metav1.ListOptions) (result *v1.ServiceMonitorList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
>>>>>>> v0.0.4
	result = &v1.ServiceMonitorList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("servicemonitors").
		VersionedParams(&opts, scheme.ParameterCodec).
<<<<<<< HEAD
=======
		Timeout(timeout).
>>>>>>> v0.0.4
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested serviceMonitors.
<<<<<<< HEAD
func (c *serviceMonitors) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
=======
func (c *serviceMonitors) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
>>>>>>> v0.0.4
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("servicemonitors").
		VersionedParams(&opts, scheme.ParameterCodec).
<<<<<<< HEAD
=======
		Timeout(timeout).
>>>>>>> v0.0.4
		Watch()
}

// Create takes the representation of a serviceMonitor and creates it.  Returns the server's representation of the serviceMonitor, and an error, if there is any.
func (c *serviceMonitors) Create(serviceMonitor *v1.ServiceMonitor) (result *v1.ServiceMonitor, err error) {
	result = &v1.ServiceMonitor{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("servicemonitors").
		Body(serviceMonitor).
		Do().
		Into(result)
	return
}

// Update takes the representation of a serviceMonitor and updates it. Returns the server's representation of the serviceMonitor, and an error, if there is any.
func (c *serviceMonitors) Update(serviceMonitor *v1.ServiceMonitor) (result *v1.ServiceMonitor, err error) {
	result = &v1.ServiceMonitor{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("servicemonitors").
		Name(serviceMonitor.Name).
		Body(serviceMonitor).
		Do().
		Into(result)
	return
}

// Delete takes name of the serviceMonitor and deletes it. Returns an error if one occurs.
<<<<<<< HEAD
func (c *serviceMonitors) Delete(name string, options *meta_v1.DeleteOptions) error {
=======
func (c *serviceMonitors) Delete(name string, options *metav1.DeleteOptions) error {
>>>>>>> v0.0.4
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicemonitors").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
<<<<<<< HEAD
func (c *serviceMonitors) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
=======
func (c *serviceMonitors) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
>>>>>>> v0.0.4
	return c.client.Delete().
		Namespace(c.ns).
		Resource("servicemonitors").
		VersionedParams(&listOptions, scheme.ParameterCodec).
<<<<<<< HEAD
=======
		Timeout(timeout).
>>>>>>> v0.0.4
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched serviceMonitor.
func (c *serviceMonitors) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.ServiceMonitor, err error) {
	result = &v1.ServiceMonitor{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("servicemonitors").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
