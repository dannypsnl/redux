package redux

import (
	"fmt"
)

func (s *Store) JSON() (str string) {
	str += "{\n"
	for k, v := range s.state {
		str += "  \"" + k + "\":" + fmt.Sprintf("%#v,\n", v)
	}
	str = str[:len(str)-2] + "\n"
	str += "}"
	return
}
