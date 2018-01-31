package store

import (
	"encoding/json"
	"fmt"
)

// Marshal serialize state of store to JSON format string
func (s *Store) Marshal() string {
	tmpMap := make(map[string]interface{})
	s.state.Range(func(key, value interface{}) bool {
		tmpMap[key.(string)] = value
		return true
	})
	if b, err := json.Marshal(tmpMap); err != nil {
		panic(err)
	} else {
		return fmt.Sprintf("%s", b)
	}
}
