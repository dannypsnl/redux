# Contributing
You can feel free to joined into this project's development.<br>
Just few points have to note:<br>
- Overview issues, you can find what can you do
- You can create issue to show what do you want to do, or your needs
- Having basic Go knowledge can help you a lot
### Structure of redux
- `store.go` type redux::Store & constructor function & it's core API method(Dispatch, Subscribe)
- `serialize.go` method of redux::Store, about serialize
- `util.go` helper function of redux(these are core inner function)
- `action.go` type redux::Action & it's helper function(these are core API)
