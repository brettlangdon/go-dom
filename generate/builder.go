package main

import (
	"bytes"
	"fmt"
)

type Builder struct {
	bytes.Buffer
}

func NewBuilder() *Builder {
	return &Builder{}
}

func (b *Builder) WriteString(s string) (int, error) {
	return b.Buffer.WriteString(s + "\r\n")
}

func (b *Builder) WriteF(f string, args ...interface{}) (int, error) {
	return b.WriteString(fmt.Sprintf(f, args...))
}
