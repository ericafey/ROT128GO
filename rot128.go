package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
Apply ROT128 on one byte: rotating one byte by 128, equal to adding 128
*/
func rotateByte(inputByte byte) byte {
	var rotateBy byte = 128
	return inputByte ^ rotateBy // XOR instead of adding to prevent overflow
}

// TODO optimize performance?
func rotateBytes(inputBytes []byte) []byte {
	var rotatedBytes []byte

	for _, inputByte := range inputBytes {
		rotatedBytes = append(rotatedBytes, rotateByte(inputByte))
	}

	return rotatedBytes
}

func main() {
	// read input from stdin
	var inputBytes []byte

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputBytes = scanner.Bytes()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	// apply rot128
	var rotatedBytes []byte = rotateBytes(inputBytes)

	// print result to stdout
	fmt.Printf("%s", rotatedBytes)
}
