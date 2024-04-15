package utils

import (
	"testing"
)

func TestNewToken(t *testing.T) {
	m := make(map[string]bool)
	for x := 1; x < 100; x++ {
		s := NewToken()

		if s == "" {
			t.Errorf("NewToken returned empty %s", s)
		}

		if m[s] {
			t.Errorf("NewToken returned duplicated token %s", s)
		}

		m[s] = true
	}
}

func TestNewTokenV2(t *testing.T) {
	m := make(map[string]bool)
	for x := 1; x < 100; x++ {
		s := NewTokenV2()

		if s == "" {
			t.Errorf("NewTokenV2 returned empty %s", s)
		}

		if m[s] {
			t.Errorf("NewTokenV2 returned duplicated token %s", s)
		}

		m[s] = true
	}
}

func BenchmarkNewToken(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t := NewToken()
			if t == "" {
				b.Fatal("NewToken returned empty")
			}
		}
	})
}

func BenchmarkNewTokenV2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			t := NewTokenV2()
			if t == "" {
				b.Fatal("NewTokenV2 returned empty")
			}
		}
	})
}
