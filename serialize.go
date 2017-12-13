package redux

import (
	"encoding/json"
	"fmt"
)

// JSON serialize state of store to JSON format string
func (s *Store) JSON() (str string) {
	var formatString string
	str += "{\n"
	for k, v := range s.state {
		b, _ := json.Marshal(v)
		formatString = fmt.Sprintf("%s", b)
		str += "  \"" + k + "\":" + formatString + ",\n"
	}
	str = str[:len(str)-2] + "\n"
	str += "}"
	return
}
