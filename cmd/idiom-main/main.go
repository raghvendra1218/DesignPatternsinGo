package main

import (
	im "github.com/DesignPatternsinGO/idiom"
)

func main() {
	err := im.NewFile("/Users/raghvendradixit/go/src/github.com/DesignPatternsinGo/file2.txt", im.Permissions(0666), im.Contents("Lorem Ipsum Dolor Amet"))
	if err != nil {
		panic(err)
	}
}
