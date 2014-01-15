// Copyright 2013 Alvaro J. Genial. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package form

import (
	"reflect"
	"testing"
)

type foo int
type bar interface {
	void()
}
type qux struct{}
type zee []bar

func TestCanIndex(t *testing.T) {

	for _, c := range []struct {
		x interface{}
		b bool
	}{
		{int(0), false},
		{foo(0), false},
		{qux{}, false},
		{(*int)(nil), false},
		{(*foo)(nil), false},
		{(*bar)(nil), false},
		{(*qux)(nil), false},
		{[]qux{}, true},
		{[5]qux{}, true},
		{&[]foo{}, true},
		{&[5]foo{}, true},
		{zee{}, true},
		{&zee{}, true},
		{map[int]foo{}, false},
		{map[interface{}]bar{}, false},
		{(chan<- int)(nil), false},
		{(chan bar)(nil), false},
		{(<-chan foo)(nil), false},
	} {
		v := reflect.ValueOf(c.x)
		if b := canIndex(v); b != c.b {
			t.Errorf("canIndex(%v)\nwant (%v)\nhave (%v)", v, c.b, b)
		}
	}
}