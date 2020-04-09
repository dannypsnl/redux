pub trait ZeroValue {
    fn zero_value() -> Self;
}

pub struct Store<State: ZeroValue, Action> {
    state: State,
    reducer: fn(&mut State, Action) -> State,
}

impl<State: ZeroValue, Action> Store<State, Action> {
    pub fn new(reducer: fn(&mut State, Action) -> State) -> Store<State, Action> {
        Store {
            state: State::zero_value().into(),
            reducer,
        }
    }

    pub fn dispatch(&mut self, action: Action) -> () {
        let reducer = self.reducer;
        let new_state = reducer(&mut self.state, action);
        self.state = new_state.into();
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn store_new() {
        let mut store = Store::new(counter);
        store.dispatch("INC");
        assert_eq!(store.state, 1);
    }

    impl ZeroValue for i64 {
        fn zero_value() -> Self {
            0
        }
    }
    fn counter(state: &mut i64, action: impl ToString) -> i64 {
        match action.to_string().as_str() {
            "INC" => *state + 1,
            "DEC" => *state - 1,
            _ => *state,
        }
    }
}
