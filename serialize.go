package redux

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// JSON serialize state of store to JSON format string
func (s *Store) JSON() (str string) {
	var formatString string
	str += "{\n"
	for k, v := range s.state {
		// TODO: fmt state is not useful, maybe use reflect get the type of state, then use json.Unmarshal jsonfy state.
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Struct:
			b, _ := json.Marshal(v)
			formatString = fmt.Sprintf("%s", b)
		case reflect.Int:
			formatString = fmt.Sprintf("%#v", v)
		case reflect.String:
			formatString = fmt.Sprintf("%#v", v)
		}
		str += "  \"" + k + "\":" + formatString + ",\n"
	}
	str = str[:len(str)-2] + "\n"
	str += "}"
	return
}
