package benchregexp

import (
	"testing"
)

var expected = `"2016-06-13T21:07:40.774000000"`

func TestFindWRegExp(t *testing.T) {
	v, err := findWRegExp()
	if err != nil {
		t.Error(err)
	} else {
		if v != expected {
			t.Errorf("got %q, want %q", v, expected)
		}
	}
}

func BenchmarkFindWRegExp(b *testing.B) {
	var timestamp string
	for i := 0; i < b.N; i++ {
		timestamp, _ = findWRegExp()
	}
	_ = timestamp
}

func TestFindWRegExpPrecompiled(t *testing.T) {
	v, err := findWRegExpPrecompiled()
	if err != nil {
		t.Error(err)
	} else {
		if v != expected {
			t.Errorf("got %q, want %q", v, expected)
		}
	}
}

func BenchmarkFindWRegExpPrecompiled(b *testing.B) {
	var timestamp string
	for i := 0; i < b.N; i++ {
		timestamp, _ = findWRegExpPrecompiled()
	}
	_ = timestamp
}

func TestFindWIndex(t *testing.T) {
	v, err := findWIndex()
	if err != nil {
		t.Error(err)
	} else {
		if v != expected {
			t.Errorf("got %q, want %q", v, expected)
		}
	}
}

func BenchmarkFindWIndex(b *testing.B) {
	var timestamp string
	for i := 0; i < b.N; i++ {
		timestamp, _ = findWIndex()
	}
	_ = timestamp
}

func TestFindWIndexBuf(t *testing.T) {
	buf := make([]byte, 31)
	err := findWIndexBuf(buf)
	if err != nil {
		t.Error(err)
	} else {
		if string(buf) != expected {
			t.Errorf("got %q, want %q", string(buf), expected)
		}
	}
}

func BenchmarkFindWIndexBuf(b *testing.B) {
	buf := make([]byte, 31)
	for i := 0; i < b.N; i++ {
		findWIndexBuf(buf)
	}
	_ = buf
}
