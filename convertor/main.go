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
	all := map[string]targetstructures.Item{}
	bugs := map[string]targetstructures.Item{}
	fish := map[string]targetstructures.Item{}
	sea := map[string]targetstructures.Item{}

	for _, item := range items {
		//d := directory
		i := item
		wp.Submit(func() {
			item := transform(i)
			// lock and add to map
			mutex.Lock()
			id := item.ID
			all[id] = item
			// append
			switch t := item.Attributes.Type.Slug; t {
			case "bugs":
				bugs[id] = item
			case "fish":
				fish[id] = item
			case "sea-creatures":
				sea[id] = item
			}

			mutex.Unlock()
			fmt.Printf("  >> [%s]:[%s] = [%s]\n", i.Type, i.Names.EuEn, item.Attributes.Titles.Safe)
		})
	}
	wp.StopWait()

	out := targetstructures.Output{
		Bugs:         bugs,
		Fish:         fish,
		SeaCreatures: sea,
	}
	write(fs, directory, out)

}
