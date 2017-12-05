grammar Redux;

Ident: StratLetter
    Letter*
    ;
    fragment
StratLetter: [a-zA-Z]
    ;
fragment
Letter: [a-zA-Z] ;
fragment
Digit: [0-9] ;

reducer:
    'func' Ident '(' Ident 'interface{}' ',' Ident 'Action' ')' 'interface{}' '{'
        Ident
    '}'
    ;