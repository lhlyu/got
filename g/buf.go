package g

import (
	"bytes"
	"fmt"
)

type bufer struct {
	buf *bytes.Buffer
}

func NewBufer() *bufer{
	return &bufer{
		buf: &bytes.Buffer{},
	}
}

func (b *bufer) String() string{
	return b.buf.String()
}

func (b *bufer) Add(param...interface{}) *bufer{
	s := fmt.Sprint(param...)
	b.buf.WriteString(s)
	return b
}

func (b *bufer) Addf(format string, param...interface{}) *bufer{
	s := fmt.Sprintf(format,param...)
	b.buf.WriteString(s)
	return b
}
