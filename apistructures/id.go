package apistructures

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
)

func (i *Item) ID() string {
	b, _ := json.Marshal(i)
	h := sha1.New()
	h.Write(b)
	str := string(h.Sum(nil))
	return fmt.Sprintf("%x", str)
}
