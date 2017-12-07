package redux

import (
	_ "encoding/json"
	"fmt"
	"reflect"
)

// JSON serialize state of store to JSON format string
func (s *Store) JSON() (str string) {
	var b interface{}
	str += "{\n"
	for k, v := range s.state {
		// TODO: fmt state is not useful, maybe use reflect get the type of state, then use json.Unmarshal jsonfy state.
		rv := reflect.ValueOf(v)
		switch rv.Kind() {
		case reflect.Struct:
		case reflect.Interface:
			fmt.Println("everything is interface{}")
		case reflect.Int:
			//fmt.Println("int")
			b = v
		case reflect.String:
			b = v
		}
		//b, _ := json.Marshal(v)
		str += "  \"" + k + "\":" + fmt.Sprintf("%#v,\n", b)
		// fmt.Sprintf("%#v,\n", v)
	}
	str = str[:len(str)-2] + "\n"
	str += "}"
	return
}
