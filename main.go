package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"

	formatter "github.com/mdigger/goldmark-formatter"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

func ReadFile(path string) []byte {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
		return nil
	}
	return dat
}

func main() {
	md := goldmark.New(
		goldmark.WithRenderer(formatter.Markdown),
		goldmark.WithExtensions(extension.GFM),
	)

	var buf bytes.Buffer
	if err := md.Convert(ReadFile("/dev/stdin"), &buf); err != nil {
		panic(err)
	}
	fmt.Println(string(buf.Bytes()))
}
