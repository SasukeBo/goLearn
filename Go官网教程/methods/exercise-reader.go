package main

/*
	练习：Reader
	实现一个Reader类型，它产生一个ASCII字符 'A' 的无限流。
*/

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: 给 MyReader 添加一个 Read([]byte) (int, error) 方法
// Finish

func (r *MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

func main() {
	reader.Validate(&MyReader{})
}
