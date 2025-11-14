// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"reflect"

	"github.com/google/uuid"
)

// InArray checks if a value exists in an array or slice.
func InArray(val interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				return true
			}
		}
	}

	return false
}

// GenerateUUID4 generates a new UUID v4 string.
func GenerateUUID4() string {
	u, err := uuid.NewRandom()
	if err != nil {
		return ""
	}
	return u.String()
}

// Unset removes an element at position i from a string slice.
func Unset(a []string, i int) []string {
	if i < 0 || i >= len(a) {
		return a
	}
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	return a[:len(a)-1]
}
