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
	wp.StopWait()

	println("\n")
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Section", "", "", "Count"})
	t.AppendRows([]table.Row{
		{"Bugs", "", "", len(out.Bugs)},
		{"Fish", "", "", len(out.Fish)},
		{"Sea", "", "", len(out.Sea)},

		{"New", "", "", len(out.New.All)},
		{"New", "North", "", len(out.New.Northern.All)},
		{"New", "North", "Bugs", len(out.New.Northern.Bugs)},
		{"New", "North", "Fish", len(out.New.Northern.Fish)},
		{"New", "North", "Sea", len(out.New.Northern.Sea)},
		{"New", "South", "", len(out.New.Southern.All)},
		{"New", "South", "Bugs", len(out.New.Southern.Bugs)},
		{"New", "South", "Fish", len(out.New.Southern.Fish)},
		{"New", "South", "Sea", len(out.New.Southern.Sea)},

		{"Leaving", "", "", len(out.Leaving.All)},
		{"Leaving", "North", "", len(out.Leaving.Northern.All)},
		{"Leaving", "North", "Bugs", len(out.Leaving.Northern.Bugs)},
		{"Leaving", "North", "Fish", len(out.Leaving.Northern.Fish)},
		{"Leaving", "North", "Sea", len(out.Leaving.Northern.Sea)},
		{"Leaving", "South", "", len(out.Leaving.Southern.All)},
		{"Leaving", "South", "Bugs", len(out.Leaving.Southern.Bugs)},
		{"Leaving", "South", "Fish", len(out.Leaving.Southern.Fish)},
		{"Leaving", "South", "Sea", len(out.Leaving.Southern.Sea)},

		{"Available", "", "", len(out.Available.All)},
		{"Available", "North", "", len(out.Available.Northern.All)},
		{"Available", "North", "Bugs", len(out.Available.Northern.Bugs)},
		{"Available", "North", "Fish", len(out.Available.Northern.Fish)},
		{"Available", "North", "Sea", len(out.Available.Northern.Sea)},
		{"Available", "South", "", len(out.Available.Southern.All)},
		{"Available", "South", "Bugs", len(out.Available.Southern.Bugs)},
		{"Available", "South", "Fish", len(out.Available.Southern.Fish)},
		{"Available", "South", "Sea", len(out.Available.Southern.Sea)},
	})

	t.AppendSeparator()
	t.AppendFooter(table.Row{"Total", "", "", len(out.All)})
	t.Render()

	write(fs, directory, out)
}
