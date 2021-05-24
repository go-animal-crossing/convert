package apistructures

import (
	"encoding/json"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/spf13/afero"
)

func Samples() map[string][]Item {
	samples := map[string][]Item{}
	fs := afero.NewOsFs()

	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	dir := d + "/_samples/*.json"

	files, _ := afero.Glob(fs, dir)

	for _, file := range files {
		name := strings.ReplaceAll(filepath.Base(file), ".json", "")
		content, _ := afero.ReadFile(fs, file)
		loaded := make([]Item, 0)
		json.Unmarshal(content, &loaded)
		samples[name] = loaded
	}

	return samples
}
