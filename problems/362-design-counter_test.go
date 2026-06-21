package problems

import "testing"

func TestCounter(t *testing.T) {
	counter := HitCounterConstructor()
	counter.Hit(1)
	counter.Hit(2)
	counter.Hit(3)

	// Test 1: Timestamp 4. Window [4-300, 4]. Hits 1, 2, 3 are all valid.
	count1 := counter.GetHits(4)
	if count1 != 3 {
		t.Errorf("Expected count to be 3, got %d", count1)
	}

	// Test 2: Timestamp 305. Window [5, 305]. Hits 1, 2, 3 are now EXPIRED.
	// Note: Your previous call to GetHits(4) might have cleaned up data depending on your implementation.
	// Assuming your implementation keeps data until it's > 300s old:
	// At timestamp 305, hit 1 (time=1) is expired (305-1 = 304 > 300).

	count2 := counter.GetHits(305)
	if count2 != 0 {
		t.Errorf("Expected count to be 0 at timestamp 305, got %d", count2)
	}

	// Better Test for "2":
	// Let's reset or use specific times.
	// If we want to test that hit 1 expires but 2 and 3 remain:
	// Current time needs to be > 301 (so 301 - 1 = 300, strictly greater? Problem says "last 300 seconds")
	// Usually "last 300 seconds" means (timestamp - 300, timestamp].
	// If timestamp is 301, window is (1, 301]. Hit at 1 is excluded.

	counter2 := HitCounterConstructor()
	counter2.Hit(1)
	counter2.Hit(2)
	counter2.Hit(3)

	count3 := counter2.GetHits(301) // Window (1, 301]. Hit 1 is out. Hits 2, 3 are in.
	if count3 != 2 {
		t.Errorf("Expected count to be 2 at timestamp 301, got %d", count3)
	}
}
