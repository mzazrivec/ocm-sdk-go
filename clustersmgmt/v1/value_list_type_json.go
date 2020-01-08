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

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

import (
	"io"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// MarshalValueList writes a list of values of the 'value' type to
// the given writer.
func MarshalValueList(list []*Value, writer io.Writer) error {
	stream := helpers.NewStream(writer)
	writeValueList(list, stream)
	stream.Flush()
	return stream.Error
}

// writeValueList writes a list of value of the 'value' type to
// the given stream.
func writeValueList(list []*Value, stream *jsoniter.Stream) {
	stream.WriteArrayStart()
	for i, value := range list {
		if i > 0 {
			stream.WriteMore()
		}
		writeValue(value, stream)
	}
	stream.WriteArrayEnd()
}

// UnmarshalValueList reads a list of values of the 'value' type
// from the given source, which can be a slice of bytes, a string or a reader.
func UnmarshalValueList(source interface{}) (items []*Value, err error) {
	iterator, err := helpers.NewIterator(source)
	if err != nil {
		return
	}
	items = readValueList(iterator)
	err = iterator.Error
	return
}

// readValueList reads list of values of the ''value' type from
// the given iterator.
func readValueList(iterator *jsoniter.Iterator) []*Value {
	list := []*Value{}
	for iterator.ReadArray() {
		item := readValue(iterator)
		list = append(list, item)
	}
	return list
}
