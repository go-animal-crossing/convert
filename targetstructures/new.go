package targetstructures

import "time"

func New() Output {
	return Output{
		Time: time.Now().UTC(),
		All:  make(map[string]Item),
	}
}
