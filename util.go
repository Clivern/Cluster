// Copyright 2020 Clivern. All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package cluster

import (
	"reflect"
	"testing"

	"github.com/satori/go.uuid"
)

// InArray check if value is on array
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

// GenerateUUID4 create a UUID
func GenerateUUID4() string {
	u := uuid.Must(uuid.NewV4(), nil)
	return u.String()
}

// Unset remove element at position i
func Unset(a []string, i int) []string {
	a[i] = a[len(a)-1]
	a[len(a)-1] = ""
	return a[:len(a)-1]
}

// Expect compare two values for testing
func Expect(t *testing.T, got, want interface{}) {
	t.Logf(`Comparing values %v, %v`, got, want)

	if !reflect.DeepEqual(got, want) {
		t.Errorf(`got %v, want %v`, got, want)
	}
}
