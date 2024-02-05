package computational_complexity

import "testing"

func TestSpaceQuadratic(t *testing.T) {
	spaceQuadratic(10)
	t.Errorf("wrong!")
}
