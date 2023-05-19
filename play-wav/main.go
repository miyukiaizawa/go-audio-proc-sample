package main

import (
	"fmt"
	"io"
	"os"

	"github.com/gordonklaus/portaudio"
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


	// PortAudioの初期化
	err = portaudio.Initialize()
	if err != nil {
		fmt.Printf("Failed to initialize PortAudio: %v", err)
		return
	}
	defer portaudio.Terminate()
	fmt.Println("portaudio.Initialized")

	// オーディオデータの読み取りバッファ作成
	bytesPerSample := int(header.BitsPerSample / 8)
	// バッファの時間長を指定 (20 ms)
	bufferSize := calculateBufferSize(header.SampleRate, 20.0)
	// バッファの作成 (モノラル再生するので１チャンネル)
	audioData := make([]int16, bufferSize*1)

	// 再生ストリームの作成
	stream, err := portaudio.OpenDefaultStream(0, int(header.NumChannels), float64(header.SampleRate), bufferSize, &audioData)
	if err != nil {
		fmt.Printf("Failed to open stream: %v", err)
		return
	}
	defer stream.Close()

	// 再生開始
	fmt.Println("Stream.Start")
	err = stream.Start()
	if err != nil {
		fmt.Printf("Failed to start stream: %v", err)
		return
	}
	defer stream.Stop()

	for {
		// バッファにオーディオデータを読み込む
		_, err := readWavData(file, audioData, bytesPerSample)
		if err != nil {
			if err == io.EOF {
				break // ファイルの終端に達したら再生終了
			} else {
				fmt.Printf("Failed to read audio data: %v", err)
				return
			}
		}

		// 再生バッファにオーディオデータを書き込む
		err = stream.Write()
		if err != nil {
			fmt.Printf("Failed to write to stream: %v", err)
			return
		}
	}	
}

// バッファーサイズを計算する
func calculateBufferSize(sampleRate uint32, bufferTimeMS float64) int {
	// サンプリングレートから秒単位の時間長に変換
	sampleTime := 1.0 / float64(sampleRate)

	// ミリ秒単位のバッファ時間を秒単位に変換
	bufferTimeInSeconds := bufferTimeMS / 1000.0

	// バッファのフレーム数を計算
	bufferSize := int(bufferTimeInSeconds / sampleTime)

	return bufferSize
}

// WAVデータを読み込む
func readWavData(file io.Reader, data []int16, bytesPerSample int) (int, error) {
	byteData := make([]byte, len(data)*bytesPerSample)

	// オーディオデータをバイト配列として読み込む
	numBytesRead, err := file.Read(byteData)
	if err != nil {
		return 0, err
	}

	// バイト配列をint16の配列に変換する
	for i := 0; i < numBytesRead; i += bytesPerSample {
		sample := int16(byteData[i]) | int16(byteData[i+1])<<8
		data[i/bytesPerSample] = sample
	}

	data = stereoToMono(data)


	return numBytesRead / bytesPerSample, nil
}


//データを抜き出す方法
func stereoToMono(inbuf []int16) []int16 {
	outbuf := make([]int16, len(inbuf)/2)

	for i := 0; i < len(inbuf); i += 2 {
		outbuf[i/2] = inbuf[i]
	}

	return outbuf
}
