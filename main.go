package main

import (
	"convert/convert"
	"convert/load"
	"fmt"

	"github.com/spf13/afero"
)

var directories = map[string]string{
	"data":      "./_source/data/",
	"images":    "./_source/images/",
	"converted": "./_source/converted/",
}

func main() {
	fmt.Printf("Starting conversion..\n")
	fs := afero.NewOsFs()
	items := load.Load(fs, directories["data"])
	convert.Convert(fs, directories["converted"], items)
	fmt.Printf("Ending conversion..\n")
}