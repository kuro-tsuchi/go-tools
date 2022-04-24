package modules

import "time"

type User struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Age      int       `json:"age"`
	Birthday time.Time `json:"birthday"`
}
