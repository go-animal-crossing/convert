package convert

import (
	"fmt"
	"strings"
)

func Safe(str string) string {
	str = strings.ReplaceAll(str, "'", "")
	str = strings.ReplaceAll(str, "\"", "")
	return str
}

func Title(str string) string {
	return strings.Title(str)
}

// Slugify converts string passed to a url friendly slug
func Slugify(str string) string {
	str = strings.ToLower(str)
	str = strings.ReplaceAll(str, " ", "-")
	str = strings.ReplaceAll(str, "'", "")
	return str
}

func URL(strs ...string) string {

	url := ""
	for _, str := range strs {
		url = fmt.Sprintf("%s/%s", url, Slugify(str))
	}
	return url
}
