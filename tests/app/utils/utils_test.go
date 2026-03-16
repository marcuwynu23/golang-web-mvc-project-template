package utils_test

import (
	"testing"

	"web_app/app/utils"
)

func TestPtrInt(t *testing.T) {
	v := 42
	p := utils.Ptr(v)

	if p == nil {
		t.Fatal("Ptr returned nil")
	}
	if *p != v {
		t.Fatalf("expected %d, got %d", v, *p)
	}
}

func TestPtrString(t *testing.T) {
	s := "hello"
	p := utils.Ptr(s)

	if p == nil {
		t.Fatal("Ptr returned nil")
	}
	if *p != s {
		t.Fatalf("expected %s, got %s", s, *p)
	}
}

