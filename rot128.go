package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
ROT128: rotating one byte by 128, equal to flipping the 8th bit of a byte
*/
func RotateByte(inputByte byte) byte {
	var rotateBy byte = 128
	return inputByte ^ rotateBy // XOR instead of adding to prevent overflow
}

func RotateBytes(inputBytes []byte) []byte {
	var rotatedBytes []byte

	for _, inputByte := range inputBytes {
		rotatedBytes = append(rotatedBytes, RotateByte(inputByte))
	}

	return rotatedBytes
}

func ReadInputBytes() []byte {
	var inputBytes []byte

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inputBytes = scanner.Bytes()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	return inputBytes
}

func main() {
	inputBytes := ReadInputBytes()

	rotatedBytes := RotateBytes(inputBytes)

	fmt.Printf("%s", rotatedBytes) 
}
