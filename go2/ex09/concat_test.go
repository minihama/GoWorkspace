package ex09_test

import (
	"bytes"
	"ex09"
	"strings"
	"testing"
)

func TestConcat(t *testing.T) {
	s := ex09.Concat("a1", "b2")
	if s != "a1b2" {
		t.Errorf("Concat(a1, b2) = %s; want a1b2", s)
	}
}

func BenchmarkConcat(b *testing.B) {
	var str string
	for n := 0; n < b.N; n++ {
		str += "x"
	}
}

func BenchmarkBuffer(b *testing.B) {
	var buffer bytes.Buffer
	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")
	}
}

func BenchmarkCopy(b *testing.B) {
	bs := make([]byte, b.N)
	bl := 0
	for n := 0; n < b.N; n++ {
		bl += copy(bs[bl:], "x")
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	var strBuilder strings.Builder
	for n := 0; n < b.N; n++ {
		strBuilder.WriteString("x")
	}
}
