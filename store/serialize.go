package store

import (
	"encoding/json"
	"fmt"
)

// Marshal serialize state of store to JSON format string
func (s *Store) Marshal() string {
	b, err := json.Marshal(s.state)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s", b)
}
