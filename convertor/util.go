package convertor

import (
	"fmt"
	"strings"
	"time"
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
	str = strings.ReplaceAll(str, `"`, "")
	return str
}

func URL(strs ...string) string {
	url := ""
	for _, str := range strs {
		url = fmt.Sprintf("%s/%s", url, Slugify(str))
	}
	return url
}

func ImagePath(itemType string, imageType string, filename string, extension string) string {
	return URL(
		itemType,
		imageType,
		fmt.Sprintf("%s.%s", filename, extension))
}

// MonthToTime creates a time from a month
func MonthToTime(month int) time.Time {
	now := time.Now()
	return time.Date(
		now.Year(),
		time.Month(month),
		1,
		0, 0, 0, 0, time.UTC)
}

func DateFormat() string {
	return "January"
}
