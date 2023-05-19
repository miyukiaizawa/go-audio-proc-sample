package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "./bird.wav" // 再生するWAVファイルのパス

	// WAVファイルをオープン
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %v", err)
		return
	}
	defer file.Close()
	fmt.Println("Open wav")

	
	header, err := LoadHeader(file)
	if err != nil {
		fmt.Printf("Failed to Open: %v", err)
	}
	fmt.Printf("%+v\n", header)
}