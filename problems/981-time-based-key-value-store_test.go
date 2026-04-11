package problems

import "testing"

func TestTimeMap(t *testing.T) {
	tm := ConstructTimeMap()
	tm.Set("foo", "bar", 1)
	val := tm.Get("foo", 1)
	if val != "bar" {
		t.Errorf("expected bar, got %s", val)
	}
	val = tm.Get("foo", 0)
	if val != "" {
		t.Errorf("expected empty, got %s", val)
	}
	tm.Set("foo", "baz", 2)
	val = tm.Get("foo", 2)
	if val != "baz" {
		t.Errorf("expected baz, got %s", val)
	}
	val = tm.Get("foo", 1)
	if val != "bar" {
		t.Errorf("expected bar, got %s", val)
	}
}

func BenchmarkTimeMap(b *testing.B) {
	// Option 1: BenchmarkTimeMap-10    	22377570	        66.39 ns/op	     131 B/op	       0 allocs/op
	// Option 2: BenchmarkTimeMap-10    	17314923	        76.77 ns/op	     135 B/op	       0 allocs/op
	tm := ConstructTimeMap()
	for b.Loop() {
		tm.Set("foo", "bar", 1)
		tm.Get("foo", 1)
	}
}
