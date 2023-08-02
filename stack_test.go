// Copyright 2015 by Leipzig University Library, http://ub.uni-leipzig.de
//
//	The Finc Authors, http://finc.info
//	Martin Czygan, <martin.czygan@uni-leipzig.de>
//
// This file is part of some open source application.
//
// Some open source application is free software: you can redistribute
// it and/or modify it under the terms of the GNU General Public
// License as published by the Free Software Foundation, either
// version 3 of the License, or (at your option) any later version.
//
// Some open source application is distributed in the hope that it will
// be useful, but WITHOUT ANY WARRANTY; without even the implied warranty
// of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Foobar.  If not, see <http://www.gnu.org/licenses/>.
//
// @license GPL-3.0+ <http://spdx.org/licenses/GPL-3.0+>
package xmlcutty

import "testing"

func TestStringStack(t *testing.T) {
	q := StringStack{}

	want := "Hello"
	q.Push(want)
	got := q.Pop()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	want = "Hello"
	q.Push(want)
	got = q.Pop()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}

	want = "Third"
	q.Push(want)
	q.Push(want)
	q.Push(want)
	got = q.Pop()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
	got = q.Pop()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
	got = q.Pop()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}

func TestString(t *testing.T) {
	q := StringStack{}
	want := "/"
	got := q.String()
	if q.String() != want {
		t.Errorf("got %s, want %s", got, want)
	}

	q = StringStack{}
	q.Push("a")
	q.Push("b")
	q.Push("c")
	want = "/a/b/c"
	got = q.String()
	if q.String() != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
