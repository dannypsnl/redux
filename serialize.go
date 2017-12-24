package redux

import (
	"encoding/json"
	"fmt"
)

// JSON serialize state of store to JSON format string
func (s *Store) JSON() string {
	if b, err := json.Marshal(s.state); err != nil {
		panic(err)
	} else {
		return fmt.Sprintf("%s", b)
	}
}
