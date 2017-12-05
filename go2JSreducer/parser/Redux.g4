grammar Redux;

WS: [ \t\n\r]+ -> channel(HIDDEN) ;
Ident: StartLetter
    Letter*
    ;
fragment
StartLetter: [a-zA-Z]
    |   '_'
    |   '\u00C0'..'\u00D6'
    |   '\u00D8'..'\u00F6'
    |   '\u00F8'..'\u02FF'
    |   '\u0370'..'\u037D'
    |   '\u037F'..'\u1FFF'
    |   '\u200C'..'\u200D'
    |   '\u2070'..'\u218F'
    |   '\u2C00'..'\u2FEF'
    |   '\u3001'..'\uD7FF'
    |   '\uF900'..'\uFDCF'
    |   '\uFDF0'..'\uFFFD'
    ;
fragment
Letter: StartLetter 
    |   '0'..'9'
    |   '\u00B7'
    |   '\u0300'..'\u036F'
    |   '\u203F'..'\u2040'
    ;
Number: Digit+;
fragment
Digit: [0-9] ;

prog: stat+ ;

stat: reducer
    | defineGlobal
    ;

reducer:
    'func' Ident '(' Ident 'interface{}' ',' Ident 'Action' ')' 'interface{}' '{'
        Ident
    '}'
    ;

method:
    Ident '(' typeFlow (',' typeFlow)* ')' typeFlow
    ;

data: Ident typeFlow;
typeFlow: Ident 
    | 'struct' '{' data '}'
    | 'interface' '{' method* '}'
    | 
    ;
define: Ident typeFlow '=' expr;
defineGlobal: 'const' define
    | 'var' define
    | 'type' Ident typeFlow
    ;

expr: expr ('+'|'-'|'*'|'/') expr
    | '(' expr ')'
    | Number
    ;