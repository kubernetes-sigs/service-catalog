// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/kubernetes-incubator/service-catalog/pkg/apis/settings/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakePodPresetBindings implements PodPresetBindingInterface
type FakePodPresetBindings struct {
	Fake *FakeSettingsV1alpha1
	ns   string
}

var podpresetbindingsResource = schema.GroupVersionResource{Group: "settings.servicecatalog.k8s.io", Version: "v1alpha1", Resource: "podpresetbindings"}

var podpresetbindingsKind = schema.GroupVersionKind{Group: "settings.servicecatalog.k8s.io", Version: "v1alpha1", Kind: "PodPresetBinding"}

// Get takes name of the podPresetBinding, and returns the corresponding podPresetBinding object, and an error if there is any.
func (c *FakePodPresetBindings) Get(name string, options v1.GetOptions) (result *v1alpha1.PodPresetBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(podpresetbindingsResource, c.ns, name), &v1alpha1.PodPresetBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodPresetBinding), err
}

// List takes label and field selectors, and returns the list of PodPresetBindings that match those selectors.
func (c *FakePodPresetBindings) List(opts v1.ListOptions) (result *v1alpha1.PodPresetBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(podpresetbindingsResource, podpresetbindingsKind, c.ns, opts), &v1alpha1.PodPresetBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.PodPresetBindingList{}
	for _, item := range obj.(*v1alpha1.PodPresetBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested podPresetBindings.
func (c *FakePodPresetBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(podpresetbindingsResource, c.ns, opts))

}

// Create takes the representation of a podPresetBinding and creates it.  Returns the server's representation of the podPresetBinding, and an error, if there is any.
func (c *FakePodPresetBindings) Create(podPresetBinding *v1alpha1.PodPresetBinding) (result *v1alpha1.PodPresetBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(podpresetbindingsResource, c.ns, podPresetBinding), &v1alpha1.PodPresetBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodPresetBinding), err
}

// Update takes the representation of a podPresetBinding and updates it. Returns the server's representation of the podPresetBinding, and an error, if there is any.
func (c *FakePodPresetBindings) Update(podPresetBinding *v1alpha1.PodPresetBinding) (result *v1alpha1.PodPresetBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(podpresetbindingsResource, c.ns, podPresetBinding), &v1alpha1.PodPresetBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodPresetBinding), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakePodPresetBindings) UpdateStatus(podPresetBinding *v1alpha1.PodPresetBinding) (*v1alpha1.PodPresetBinding, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(podpresetbindingsResource, "status", c.ns, podPresetBinding), &v1alpha1.PodPresetBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodPresetBinding), err
}

// Delete takes name of the podPresetBinding and deletes it. Returns an error if one occurs.
func (c *FakePodPresetBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(podpresetbindingsResource, c.ns, name), &v1alpha1.PodPresetBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakePodPresetBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(podpresetbindingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.PodPresetBindingList{})
	return err
}

// Patch applies the patch and returns the patched podPresetBinding.
func (c *FakePodPresetBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.PodPresetBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(podpresetbindingsResource, c.ns, name, data, subresources...), &v1alpha1.PodPresetBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.PodPresetBinding), err
}
