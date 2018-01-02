# Contributing
You can feel free to joined into this project's development.<br>
Just few points have to note:<br>
- Overview issues, you can find what can you do
- You can create issue to show what do you want to do, or your needs
- Having basic Go knowledge can help you a lot
### Structure of redux
- `store/store.go` type store::Store & constructor function & it's core API method(Dispatch, Subscribe)
- `store/serialize.go` method of store::Store, about serialize
- `store/util.go` helper function for store
- `action/action.go` type action::Action & it's core API
