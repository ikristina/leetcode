package problems

import "testing"

func TestRandomizedSet(t *testing.T) {
	rs := Constructor()

	// Test insert
	if got := rs.Insert(1); got != true {
		t.Errorf("Insert(1) = %v, want true", got)
	}
	if got := rs.Insert(2); got != true {
		t.Errorf("Insert(2) = %v, want true", got)
	}
	if got := rs.Insert(1); got != false {
		t.Errorf("Insert(1) = %v, want false", got)
	}

	// Test getRandom
	for i := 0; i < 10; i++ {
		got := rs.GetRandom()
		if got != 1 && got != 2 {
			t.Errorf("GetRandom() = %v, want 1 or 2", got)
		}
	}

	// Test remove
	if got := rs.Remove(1); got != true {
		t.Errorf("Remove(1) = %v, want true", got)
	}
	if got := rs.Remove(1); got != false {
		t.Errorf("Remove(1) = %v, want false", got)
	}

	// After removing 1, only 2 should remain
	for i := 0; i < 5; i++ {
		if got := rs.GetRandom(); got != 2 {
			t.Errorf("GetRandom() = %v, want 2", got)
		}
	}
}
