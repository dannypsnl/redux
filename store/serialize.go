package store

import (
	"encoding/json"
	"fmt"
)

// Marshal serialize state of store to JSON format string
func (s *Store) Marshal() string {
	if b, err := json.Marshal(s.state); err != nil {
		panic(err)
	} else {
		return fmt.Sprintf("%s", b)
	}
}
