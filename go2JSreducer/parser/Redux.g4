grammar Redux;

WS: [ \t\n\r]+ -> channel(HIDDEN) ;
STRING: '"' .*? '"'
    | '`' .*? '`'
    ;
ID: StartLetter
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
    | function
    | defineGlobal
    | packageStat
    | importStat
    ;

function:
    'func' ID '(' (ID typeFlow (',' ID typeFlow)*)? ')' '{'
        stat
    '}'
    ;

reducer:
    'func' ID '(' ID 'interface{}' ',' ID 'redux.Action' ')' 'interface{}' '{'
        'if' ID '==' 'nil' '{'
            'return' initialState=expr
        '}'
        'switch' ID '.' 'Type' '{'
        cases
        'default' ':'
            'return' expr
        '}'
    '}'
    ;
cases:
    ('case' expr ':'
        'return' expr)+
    ;
method:
    ID '(' typeFlow (',' typeFlow)* ')' typeFlow
    ;

data: ID typeFlow;
typeFlow: ID 
    | 'struct' '{' data '}'
    | 'interface' '{' method* '}'
    | 
    ;
define: ID typeFlow '=' expr;
defineGlobal: 'const' define
    | 'var' define
    | 'type' ID typeFlow
    ;

expr: expr ('+'|'-'|'*'|'/') expr
    | '(' expr ')'
    | expr '.' '(' ID ')' // cast
    | Number
    | STRING
    | ID
    ;

packageStat: 'package' ID ;
importStat: 'import' '(' STRING* ')'
    | 'import' STRING
    ;