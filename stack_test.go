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
