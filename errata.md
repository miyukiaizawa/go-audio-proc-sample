## 正誤表

下記の通り、本文で一部誤記がありました。お詫びして訂正いたします。 

31ページ
2.6.1のリスト2.10 無音に変換

訂正前
```go
func muteRightChannel(inbuf []int16) []int16 {
	outbuf := make([]int16, len(inbuf))

	for i := 1; i < len(inbuf); i += 2 {
		outbuf[i] = 0
	}

	return outbuf
}
```

訂正後
```go
func muteRightChannel(inbuf []int16) []int16 {
	outbuf := inbuf

	for i := 1; i < len(inbuf); i += 2 {
		outbuf[i] = 0
	}

	return outbuf
}
```