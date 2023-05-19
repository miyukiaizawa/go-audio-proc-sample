package main

import (
	"encoding/binary"
	"fmt"
	"io"
)

type WAVHeader struct {
	ChunkID [4]byte `json:"ChunkID"`
	ChunkSize uint32 `json:"ChunkSize"`
	Format [4]byte `json:"Format"`
	Subchunk1ID [4]byte `json:"Subchunk1ID"`
	Subchunk1Size uint32 `json:"Subchunk1Size"`
	AudioFormat uint16 `json:"AudioFormat"`
	NumChannels uint16 `json:"NumChannels"`
	SampleRate uint32 `json:"SampleRate"`
	ByteRate uint32 `json:"ByteRate"`
	BlockAlign uint16 `json:"BlockAlign"`
	BitsPerSample uint16 `json:"BitsPerSample"`
	Subchunk2ID [4]byte `json:"Subchunk2ID"`
	Subchunk2Size uint32 `json:"Subchunk2Size"`
}

func LoadHeader(file io.Reader) (*WAVHeader, error) {
	h := &WAVHeader{}

	// チャンクIDのチェック
	if err := binary.Read(file, binary.LittleEndian, &h.ChunkID); err != nil {
		fmt.Printf("Failed to read chunkID: %v", err)
		return nil, err
	}
	if string(h.ChunkID[:]) != "RIFF" {
		fmt.Println("Invalid WAV file format")
		return nil, fmt.Errorf("Invalid WAV file format")
	}

	// チャンクサイズの読み取り（使用しない）
	if err := binary.Read(file, binary.LittleEndian, &h.ChunkSize); err != nil {
		fmt.Printf("Failed to read chunkSize: %v", err)
		return nil, err
	}

	// フォーマットのチェック
	if err := binary.Read(file, binary.LittleEndian, &h.Format); err != nil {
		fmt.Printf("Failed to read format: %v", err)
		return nil, err
	}
	if string(h.Format[:]) != "WAVE" {
		fmt.Println("Invalid WAV file format")
		return nil, fmt.Errorf("Invalid WAV file format")
	}

	// fmtチャンクのチェック
	if err := binary.Read(file, binary.LittleEndian, &h.Subchunk1ID); err != nil {
		fmt.Printf("Failed to read subchunk1ID: %v", err)
		return nil, err
	}
	if string(h.Subchunk1ID[:]) != "fmt " {
		fmt.Println("Invalid WAV file format")
		return nil, fmt.Errorf("Invalid WAV file format")
	}

	// fmtチャンクサイズの読み取り（使用しない）
	if err := binary.Read(file, binary.LittleEndian, &h.Subchunk1Size); err != nil {
		fmt.Printf("Failed to read subchunk1Size: %v", err)
		return nil, err
	}

	// オーディオフォーマットの読み取り
	if err := binary.Read(file, binary.LittleEndian, &h.AudioFormat); err != nil {
		fmt.Printf("Failed to read audioFormat: %v", err)
		return nil, err
	}

	// チャンネル数の読み取り
	if err := binary.Read(file, binary.LittleEndian, &h.NumChannels); err != nil {
		fmt.Printf("Failed to read numChannels: %v", err)
		return nil, err
	}
	// サンプリングレートの読み取り
	if err := binary.Read(file, binary.LittleEndian, &h.SampleRate); err != nil {
		fmt.Printf("Failed to read sampleRate: %v", err)
		return nil, err
	}

	// バイトレートの読み取り（使用しない）
	if err := binary.Read(file, binary.LittleEndian, &h.ByteRate); err != nil {
		fmt.Printf("Failed to read byteRate: %v", err)
		return nil, err
	}

	// ブロックアラインの読み取り（使用しない）
	if err := binary.Read(file, binary.LittleEndian, &h.BlockAlign); err != nil {
		fmt.Printf("Failed to read blockAlign: %v", err)
		return nil, err
	}

	// サンプルあたりのビット数の読み取り
	if err := binary.Read(file, binary.LittleEndian, &h.BitsPerSample); err != nil {
		fmt.Printf("Failed to read bitsPerSample: %v", err)
		return nil, err
	}

	// dataチャンクのチェック
	if err := binary.Read(file, binary.LittleEndian, &h.Subchunk2ID); err != nil {
		fmt.Printf("Failed to read subchunk2ID: %v", err)
		return nil, err
	}
	if string(h.Subchunk2ID[:]) != "data" {
		fmt.Println("Invalid WAV file format")
		return nil, fmt.Errorf("Invalid WAV file format")
	}

	// オーディオデータサイズの読み取り
	if err := binary.Read(file, binary.LittleEndian, &h.Subchunk2Size); err != nil {
		fmt.Printf("Failed to read subchunk2Size: %v", err)
		return nil, err
	}
	return h, nil
}