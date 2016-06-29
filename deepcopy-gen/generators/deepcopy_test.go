/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package generators

import "testing"

func Test_isRootedUnder(t *testing.T) {
	testCases := []struct {
		path   string
		roots  []string
		expect bool
	}{
		{
			path:   "/foo/bar",
			roots:  nil,
			expect: false,
		},
		{
			path:   "/foo/bar",
			roots:  []string{},
			expect: false,
		},
		{
			path: "/foo/bar",
			roots: []string{
				"/bad",
			},
			expect: false,
		},
		{
			path: "/foo/bar",
			roots: []string{
				"/foo",
			},
			expect: true,
		},
		{
			path: "/foo/bar",
			roots: []string{
				"/bad",
				"/foo",
			},
			expect: true,
		},
		{
			path: "/foo/bar/qux/zorb",
			roots: []string{
				"/foo/bar/qux",
			},
			expect: true,
		},
		{
			path: "/foo/bar",
			roots: []string{
				"/foo/bar",
			},
			expect: true,
		},
		{
			path: "/foo/barn",
			roots: []string{
				"/foo/bar",
			},
			expect: false,
		},
		{
			path: "/foo/bar",
			roots: []string{
				"/foo/barn",
			},
			expect: false,
		},
		{
			path: "/foo/bar",
			roots: []string{
				"",
			},
			expect: true,
		},
	}

	for i, tc := range testCases {
		r := isRootedUnder(tc.path, tc.roots)
		if r != tc.expect {
			t.Errorf("case[%d]: expected %t, got %t for %q in %q", i, tc.expect, r, tc.path, tc.roots)
		}
	}
}
