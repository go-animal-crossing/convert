package load

import (
	"convert/apistructures"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/spf13/afero"
)

func Load(fs afero.Fs, directory string) []apistructures.Item {
	items := make([]apistructures.Item, 0)

	pattern := directory + "*.json"
	files, _ := afero.Glob(fs, pattern)
	fmt.Printf("  > Found [%d] files\n", len(files))

	for _, file := range files {
		itemType := strings.ReplaceAll(filepath.Base(file), ".json", "")
		fmt.Printf("  > Loading items from file: [%s]\n", file)

		content, _ := afero.ReadFile(fs, file)
		loaded := make([]apistructures.Item, 0)
		json.Unmarshal(content, &loaded)

		for _, item := range loaded {
			item.Type = itemType
			items = append(items, item)
		}
	}
	fmt.Printf("  > Loaded [%d] items\n", len(items))

	return items
}
