package main

//データを抜き出す方法
func stereoToMono1(inbuf []int16) []int16 {
	outbuf := make([]int16, len(inbuf)/2)

	for i := 0; i < len(inbuf); i += 2 {
		outbuf[i/2] = inbuf[i]
	}

	return outbuf
}

//左右のサンプルを平均化する方法
func stereoToMono2(inbuf []int16) []int16 {
	outbuf := make([]int16, len(inbuf)/2)

	for i := 0; i < len(inbuf); i += 2 {
		v1 := float64(inbuf[i]) 
		v2 := float64(inbuf[i+1])
		sample := int16((v1 + v2) / 2.0)
		outbuf[i/2] = sample
	}

	return outbuf
}