package main

/*
	练习：rot13Reader

	有种常见的模式是一个io.Reader包装另一个io.Reader，然后通过某种方式修改其数据流。

	例如，gzip.NewReader函数接受一个io.Reader（已压缩的数据流）并返回一个同样实现了io.Reader
	的*gzip.Reader（解压后的数据流）。

	编写一个实现了io.Reader并从另一个io.Reader中读取数据的rot13Reader，
	通过应用rot13代换密码对数据流进行修改。

	rot13Reader类型已经提供。实现Read方法以满足io.Reader。
*/

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(b []byte) (int, error) {
	buffLength := len(b)
	buff := make([]byte, buffLength)
	for {
		_, err := rot13.r.Read(buff)

		if err == io.EOF {
			break
		}
	}

	for i := 0; i < len(buff); i++ {
		ch := buff[i]
		switch {
		case (ch > 97) && (ch < 122):
			b[i] = (ch-97+13)%26 + 97

		case (ch > 65) && (ch < 90):
			b[i] = (ch-65+13)%26 + 65

		default:
			b[i] = ch
		}
	}

	return len(buff), io.EOF
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
