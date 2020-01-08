/*
Copyright (c) 2019 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

// ResourceKind is the name of the type used to represent objects
// of type 'resource'.
const ResourceKind = "Resource"

// ResourceLinkKind is the name of the type used to represent links
// to objects of type 'resource'.
const ResourceLinkKind = "ResourceLink"

// ResourceNilKind is the name of the type used to nil references
// to objects of type 'resource'.
const ResourceNilKind = "ResourceNil"

// Resource represents the values of the 'resource' type.
//
// Identifies computing resources
type Resource struct {
	id           *string
	href         *string
	link         bool
	allowed      *int
	resourceName *string
	resourceType *string
}

// Kind returns the name of the type of the object.
func (o *Resource) Kind() string {
	if o == nil {
		return ResourceNilKind
	}
	if o.link {
		return ResourceLinkKind
	}
	return ResourceKind
}

// ID returns the identifier of the object.
func (o *Resource) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *Resource) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *Resource) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *Resource) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *Resource) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *Resource) Empty() bool {
	return o == nil || (o.id == nil &&
		o.allowed == nil &&
		o.resourceName == nil &&
		o.resourceType == nil &&
		true)
}

// Allowed returns the value of the 'allowed' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Number of allowed nodes
func (o *Resource) Allowed() int {
	if o != nil && o.allowed != nil {
		return *o.allowed
	}
	return 0
}

// GetAllowed returns the value of the 'allowed' attribute and
// a flag indicating if the attribute has a value.
//
// Number of allowed nodes
func (o *Resource) GetAllowed() (value int, ok bool) {
	ok = o != nil && o.allowed != nil
	if ok {
		value = *o.allowed
	}
	return
}

// ResourceName returns the value of the 'resource_name' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// platform-specific name, such as "M5.2Xlarge" for a type of EC2 node
func (o *Resource) ResourceName() string {
	if o != nil && o.resourceName != nil {
		return *o.resourceName
	}
	return ""
}

// GetResourceName returns the value of the 'resource_name' attribute and
// a flag indicating if the attribute has a value.
//
// platform-specific name, such as "M5.2Xlarge" for a type of EC2 node
func (o *Resource) GetResourceName() (value string, ok bool) {
	ok = o != nil && o.resourceName != nil
	if ok {
		value = *o.resourceName
	}
	return
}

// ResourceType returns the value of the 'resource_type' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
//
func (o *Resource) ResourceType() string {
	if o != nil && o.resourceType != nil {
		return *o.resourceType
	}
	return ""
}

// GetResourceType returns the value of the 'resource_type' attribute and
// a flag indicating if the attribute has a value.
//
//
func (o *Resource) GetResourceType() (value string, ok bool) {
	ok = o != nil && o.resourceType != nil
	if ok {
		value = *o.resourceType
	}
	return
}

// ResourceListKind is the name of the type used to represent list of objects of
// type 'resource'.
const ResourceListKind = "ResourceList"

// ResourceListLinkKind is the name of the type used to represent links to list
// of objects of type 'resource'.
const ResourceListLinkKind = "ResourceListLink"

// ResourceNilKind is the name of the type used to nil lists of objects of
// type 'resource'.
const ResourceListNilKind = "ResourceListNil"

// ResourceList is a list of values of the 'resource' type.
type ResourceList struct {
	href  *string
	link  bool
	items []*Resource
}

// Kind returns the name of the type of the object.
func (l *ResourceList) Kind() string {
	if l == nil {
		return ResourceListNilKind
	}
	if l.link {
		return ResourceListLinkKind
	}
	return ResourceListKind
}

// Link returns true iif this is a link.
func (l *ResourceList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *ResourceList) HREF() string {
	if l != nil && l.href != nil {
		return *l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *ResourceList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != nil
	if ok {
		value = *l.href
	}
	return
}

// Len returns the length of the list.
func (l *ResourceList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *ResourceList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *ResourceList) Get(i int) *Resource {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *ResourceList) Slice() []*Resource {
	var slice []*Resource
	if l == nil {
		slice = make([]*Resource, 0)
	} else {
		slice = make([]*Resource, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ResourceList) Each(f func(item *Resource) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *ResourceList) Range(f func(index int, item *Resource) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
