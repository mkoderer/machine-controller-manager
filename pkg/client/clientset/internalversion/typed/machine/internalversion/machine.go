package internalversion

import (
	machine "github.com/gardener/node-controller-manager/pkg/apis/machine"
	scheme "github.com/gardener/node-controller-manager/pkg/client/clientset/internalversion/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MachinesGetter has a method to return a MachineInterface.
// A group's client should implement this interface.
type MachinesGetter interface {
	Machines() MachineInterface
}

// MachineInterface has methods to work with Machine resources.
type MachineInterface interface {
	Create(*machine.Machine) (*machine.Machine, error)
	Update(*machine.Machine) (*machine.Machine, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*machine.Machine, error)
	List(opts v1.ListOptions) (*machine.MachineList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *machine.Machine, err error)
	MachineExpansion
}

// machines implements MachineInterface
type machines struct {
	client rest.Interface
}

// newMachines returns a Machines
func newMachines(c *MachineClient) *machines {
	return &machines{
		client: c.RESTClient(),
	}
}

// Get takes name of the machine, and returns the corresponding machine object, and an error if there is any.
func (c *machines) Get(name string, options v1.GetOptions) (result *machine.Machine, err error) {
	result = &machine.Machine{}
	err = c.client.Get().
		Resource("machines").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Machines that match those selectors.
func (c *machines) List(opts v1.ListOptions) (result *machine.MachineList, err error) {
	result = &machine.MachineList{}
	err = c.client.Get().
		Resource("machines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested machines.
func (c *machines) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("machines").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a machine and creates it.  Returns the server's representation of the machine, and an error, if there is any.
func (c *machines) Create(machine *machine.Machine) (result *machine.Machine, err error) {
	result = &machine.Machine{}
	err = c.client.Post().
		Resource("machines").
		Body(machine).
		Do().
		Into(result)
	return
}

// Update takes the representation of a machine and updates it. Returns the server's representation of the machine, and an error, if there is any.
func (c *machines) Update(machine *machine.Machine) (result *machine.Machine, err error) {
	result = &machine.Machine{}
	err = c.client.Put().
		Resource("machines").
		Name(machine.Name).
		Body(machine).
		Do().
		Into(result)
	return
}

// Delete takes name of the machine and deletes it. Returns an error if one occurs.
func (c *machines) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("machines").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *machines) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Resource("machines").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched machine.
func (c *machines) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *machine.Machine, err error) {
	result = &machine.Machine{}
	err = c.client.Patch(pt).
		Resource("machines").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}