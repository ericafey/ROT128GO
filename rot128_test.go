package main

import (
	"fmt"
	"testing"
	"bytes"
)

func TestRotateByteHappyPath(t *testing.T) {
	var tests = []struct {
		testByte, want byte
	}{
		{0, 128},
		{1, 129},
		{127,255},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.testByte)
		t.Run(testname, func(t *testing.T) {
			got := RotateByte(tt.testByte)
			if got != tt.want {
				t.Errorf("got %d, want %d", got, tt.want)
			}
		})
	}
}

func TestRotateByteOverflow(t *testing.T){
	var testByte, want  byte = 128, 0

	got := RotateByte(testByte)
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestRotateBytesHelloWorld(t *testing.T) {
	testBytes := []byte("Hello World!")

	got := RotateBytes(RotateBytes(testBytes))
	if !bytes.Equal(got,testBytes) {
		t.Errorf("got %d, want %d", got, testBytes)
	}
}

func TestRotateBytesNewline(t *testing.T) {
	testBytes := []byte("Hello World!\n Bye World.")

	got := RotateBytes(RotateBytes(testBytes))
	if !bytes.Equal(got,testBytes) {
		
		t.Errorf("got %d, want %d", got, testBytes)
	}
}

func TestRotateBytesEmpty(t *testing.T) {
	testBytes := []byte("")

	got := RotateBytes(RotateBytes(testBytes))
	if !bytes.Equal(got,testBytes) {
		t.Errorf("got %d, want %d", got, testBytes)
	}
}


func BenchmarkRotateByte(b *testing.B) {
	var testByte byte = 0
	for b.Loop(){
		RotateByte(testByte)
	}
}

func BenchmarkRotateBytes(b *testing.B) {
	testBytes := []byte("Hello World!")

	for b.Loop(){
		RotateBytes(testBytes)
	}
}
