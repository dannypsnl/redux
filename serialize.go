package redux

import (
	"encoding/json"
	"fmt"
)

// JSON serialize state of store to JSON format string
func (s *Store) JSON() (str string) {
	str += "{\n"
	for k, v := range s.state {
		// json.Marshal allow interface{}
		b, _ := json.Marshal(v)
		str += "  \"" + k + "\":" + fmt.Sprintf("%s", b) + ",\n"
	}
	str = str[:len(str)-2] + "\n"
	str += "}"
	return
}
