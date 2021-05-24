package convertor

import (
	"convert/apistructures"
	"convert/targetstructures"
	"fmt"
	"sync"

	"github.com/gammazero/workerpool"
	"github.com/spf13/afero"
)

var poolSize = 1

var typeMeta = map[string]targetstructures.TypeMeta{
	"bugs": {Title: "Bugs", Slug: "bugs"},
	"fish": {Title: "Fish", Slug: "fish"},
	"sea":  {Title: "Sea Creatures", Slug: "sea-creatures"},
}

func Convert(fs afero.Fs, directory string, items []apistructures.Item) map[string]targetstructures.Item {
	var mutex = &sync.Mutex{}

	wp := workerpool.New(poolSize)
	converted := map[string]targetstructures.Item{}

	for _, item := range items {
		//d := directory
		i := item
		wp.Submit(func() {
			item := transform(i)
			// lock and add to map
			mutex.Lock()
			converted[item.ID] = item
			mutex.Unlock()
			fmt.Printf("  >> [%s]:[%s] = [%s]\n", i.Type, i.Names.EuEn, item.Attributes.Titles.Safe)
		})
	}
	wp.StopWait()

	return converted
}
