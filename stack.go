// Copyright 2015 by Leipzig University Library, http://ub.uni-leipzig.de
//                   The Finc Authors, http://finc.info
//                   Martin Czygan, <martin.czygan@uni-leipzig.de>
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
//
//
// package xmlcutty implements support for the xmlcutty command line tool.
package xmlcutty

import "strings"

// StringStack implements LIFO. Not thread safe.
type StringStack struct {
	queue []string
}

// Push adds an element to the stack.
func (q *StringStack) Push(s string) {
	q.queue = append(q.queue, s)
}

// Top retrieves the last added element. Panics on an empty stack.
func (q *StringStack) Top() string {
	if len(q.queue) == 0 {
		panic("Top from empty queue")
	}
	return q.queue[len(q.queue)-1]
}

// Pop removes the last added element from the stack and returns it. Panics on
// an empty stack.
func (q *StringStack) Pop() string {
	if len(q.queue) == 0 {
		panic("Pop from empty queue")
	}
	r := q.Top()
	q.queue = q.queue[:len(q.queue)-1]
	return r
}

// String formats the stack in a path-like manner.
func (q *StringStack) String() string {
	return "/" + strings.Join(q.queue, "/")
}
