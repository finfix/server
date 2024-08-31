package testutil

import "bytes"

type Buffer struct {
	bytes.Buffer
}

func (b *Buffer) Write(p []byte) (n int, err error) {
	b.Buffer.Reset()
	return b.Buffer.Write(p)
}
