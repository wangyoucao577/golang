package ch6_intset

import "testing"

func TestHas(t *testing.T) {

	var s IntSet
	if s.Has(0) {
		t.Errorf("s: %s, want s.Has(0) false", s.String())
	}

	s.Add(0)
	if !s.Has(0) {
		t.Errorf("s: %s, want s.Has(0) true", s.String())
	}

}
