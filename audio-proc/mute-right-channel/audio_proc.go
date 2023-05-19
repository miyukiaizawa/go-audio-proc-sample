package main

// 右チャンネルを無音にする
func muteRightChannel(inbuf []int16) []int16 {
	outbuf := inbuf

	for i := 1; i < len(inbuf); i += 2 {
		outbuf[i] = 0
	}

	return outbuf
}

// 左チャンネルを無音にする
func muteLeftChannel(inbuf []int16) []int16 {
	outbuf := inbuf

	for i := 0; i < len(inbuf); i += 2 {
		outbuf[i] = 0
	}

	return outbuf
}
