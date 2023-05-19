package main

// 音量の変更
func changeVol(inbuf []int16, volLeft float64, volRight float64) []int16 {
	outbuf := make([]int16, len(inbuf))

	for i := 0; i < len(inbuf); i += 2 {
		v1 := float64(inbuf[i])
		v2 := float64(inbuf[i+1])
		outbuf[i] = int16(v1 * volLeft)
		outbuf[i+1] = int16(v2 * volRight)
	}

	return outbuf
}
