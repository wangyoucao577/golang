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

func TestAddAll(t *testing.T) {
	var s IntSet

	s.AddAll(1, 2, 5)
	if s.Len() != 3 || !s.Has(1) || !s.Has(2) || !s.Has(5) {
		t.Errorf("s.AddAll(1, 2, 3), result s: %s", s.String())
	}

	s.AddAll(4, 10)
	if s.Len() != 5 || !s.Has(1) || !s.Has(2) || !s.Has(5) || !s.Has(4) || !s.Has(10) {
		t.Errorf("s.AddAll(4, 10), result s: %s", s.String())
	}

}
