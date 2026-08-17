package main

import (
	"testing"
	"fmt"
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
	testBytes := []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.  Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.  Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.  Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum. Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.  Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")

	for b.Loop(){
		RotateBytes(testBytes)
	}
}
