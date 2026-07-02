package utils

import (
	"regexp"
	"testing"
)

func TestOrderID(t *testing.T) {
	re := regexp.MustCompile(`^TCC-[A-Z0-9]{6}-\d+$`)
	seen := map[string]bool{}
	for i := 0; i < 1000; i++ {
		id := OrderID()
		if !re.MatchString(id) {
			t.Fatalf("bad order id format: %q", id)
		}
		seen[id] = true
	}
	if len(seen) < 990 {
		t.Fatalf("too many collisions: %d unique of 1000", len(seen))
	}
}
