package convertor

import (
	"convert/apistructures"
	"convert/targetstructures"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/gammazero/workerpool"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/afero"
)

var poolSize = 10

var typeMeta = map[string]targetstructures.TypeMeta{
	"bugs": {Title: "Bugs", Slug: "bugs"},
	"fish": {Title: "Fish", Slug: "fish"},
	"sea":  {Title: "Sea Creatures", Slug: "sea-creatures"},
}

func write(fs afero.Fs, directory string, out targetstructures.Output) {
	content, _ := json.Marshal(out)
	filename := directory + "out.json"

	fileDir := filepath.Dir(filename)
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		fs.MkdirAll(fileDir, os.ModePerm)
	}

	afero.WriteFile(fs, filename, content, os.FileMode(int(0777)))
}

func Convert(fs afero.Fs, directory string, items []apistructures.Item) {
	var mutex = &sync.Mutex{}
	wp := workerpool.New(poolSize)
	out := targetstructures.New()
	fmt.Printf("\nConverting\n")
	for _, item := range items {
		//d := directory
		i := item
		wp.Submit(func() {
			transformed := Transform(i)
			// lock and add to map
			mutex.Lock()
			out.Add(transformed)
			mutex.Unlock()

			fmt.Printf("  %-6v| %s\n", i.Type, transformed.Attributes.URIS.Slug)
		})
	}
	out.Sort()
	wp.StopWait()

	println("\n")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Section", "", "", "Count"})

	t.AppendSeparator()
	t.AppendFooter(table.Row{"Total", "", "", len(out.All)})
	t.Render()

	write(fs, directory, out)
}
