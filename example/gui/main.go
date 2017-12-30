package main

import (
	"fmt"
	"github.com/andlabs/ui"
	"github.com/dannypsnl/redux/action"
	"github.com/dannypsnl/redux/store"
)

func counter(state interface{}, act action.Action) interface{} {
	if state == nil {
		return 0
	}
	switch act.Type {
	case "INC":
		return state.(int) + 1
	case "DEC":
		return state.(int) - 1
	default:
		return state
	}
}

func main() {
	store := store.New(counter)
	err := ui.Main(func() {
		buttonInc := ui.NewButton("+")
		buttonDec := ui.NewButton("-")
		number := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(buttonInc, false)
		box.Append(buttonDec, false)
		box.Append(number, false)
		window := ui.NewWindow("Hello", 200, 100, false)
		window.SetChild(box)
		buttonInc.OnClicked(func(*ui.Button) {
			store.Dispatch(action.New("INC"))
			s := fmt.Sprintf("Number is:%d", store.GetState("counter").(int))
			number.SetText(s)
		})
		buttonDec.OnClicked(func(*ui.Button) {
			store.Dispatch(action.New("DEC"))
			s := fmt.Sprintf("Number is:%d", store.GetState("counter").(int))
			number.SetText(s)
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}
