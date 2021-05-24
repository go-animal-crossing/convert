package load

import (
	"convert/apistructures"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

func Files(fs afero.Fs, directory string, pattern string) ([]string, error) {
	path := directory + pattern
	return afero.Glob(fs, path)
}

func UnmarshalFile(content []byte) []apistructures.Item {
	loaded := make([]apistructures.Item, 0)
	json.Unmarshal(content, &loaded)
	return loaded
}

func Load(fs afero.Fs, directory string) []apistructures.Item {
	items := make([]apistructures.Item, 0)
	files, _ := Files(fs, directory, "*.json")
	fmt.Printf("Found [%d] files\n", len(files))

	for _, file := range files {
		itemType := strings.ReplaceAll(filepath.Base(file), ".json", "")
		fmt.Printf("Loading items from file: [%s]\n", file)

		content, _ := afero.ReadFile(fs, file)
		loaded := UnmarshalFile(content)

		for _, item := range loaded {
			// apply any fixes to the data
			item = apistructures.DataFixes(item)
			item.Type = itemType
			items = append(items, item)
		}

	}
	fmt.Printf("Loaded [%d] items\n", len(items))

	return items
}
