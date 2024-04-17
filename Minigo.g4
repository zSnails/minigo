grammar Minigo;

root: PACKAGE IDENTIFIER SEMICOLON topDeclarationList;

topDeclarationList: (variableDecl | typeDecl | funcDecl)*;

variableDecl: VAR singleVarDecl SEMICOLON                                   #variableDeclaration
            | VAR LEFTPARENTHESIS innerVarDecls RIGHTPARENTHESIS SEMICOLON  #multiVariableDeclaration
            | VAR LEFTPARENTHESIS RIGHTPARENTHESIS SEMICOLON                #emptyVariableDeclaration
            ;
innerVarDecls: singleVarDecl SEMICOLON (singleVarDecl SEMICOLON)*;

singleVarDecl: identifierList declType EQUALS expressionList #typedVarDecl
             | identifierList EQUALS expressionList #untypedVarDecl
             | singleVarDeclNoExps #singleVarDeclsNoExpsDecl
             ;

singleVarDeclNoExps: identifierList declType
                   ;

typeDecl: TYPE singleTypeDecl SEMICOLON #typeDeclaration
        | TYPE LEFTPARENTHESIS innerTypeDecls RIGHTPARENTHESIS SEMICOLON #multiTypeDeclaration
        | TYPE LEFTPARENTHESIS RIGHTPARENTHESIS SEMICOLON #emptyTypeDeclaration
        ;

innerTypeDecls: singleTypeDecl SEMICOLON (singleTypeDecl SEMICOLON)*
              ;

singleTypeDecl: IDENTIFIER declType
              ;

funcDecl: funcFrontDecl block SEMICOLON
        ;

funcFrontDecl: FUNC IDENTIFIER LEFTPARENTHESIS funcArgsDecls? RIGHTPARENTHESIS declType?
             ;

funcArgsDecls: singleVarDeclNoExps (COMMA singleVarDeclNoExps)*
             ;

declType: LEFTPARENTHESIS declType RIGHTPARENTHESIS
        | IDENTIFIER
        | sliceDeclType
        | arrayDeclType
        | structDeclType
        ;

sliceDeclType: LEFTBRACKET RIGHTBRACKET declType
             ;
arrayDeclType: LEFTBRACKET INTLITERAL RIGHTBRACKET declType
             ;

structDeclType: STRUCT LEFTCURLYBRACE structMemDecls* RIGHTCURLYBRACE
              ;

structMemDecls: singleVarDeclNoExps SEMICOLON (singleVarDeclNoExps SEMICOLON)*
              ;

identifierList: IDENTIFIER (COMMA IDENTIFIER)*
              ;

expression: primaryExpression #expressionPrimaryExpression
          | left=expression (TIMES | DIV | MOD | LEFTSHIFT | RIGHTSHIFT | AMPERSAND | AMPERSANDCARET | PLUS | MINUS | PIPE | CARET | COMPARISON | NEGATION | LESSTHAN | GREATERTHAN | LESSTHANEQUAL | GREATERTHANEQUAL | AND | OR) right=expression #operation
          | PLUS expression #plusExpression
          | MINUS expression #minusExpression
          | NOT expression #notExpression
          | CARET expression #caretExpression
          ;

expressionList: expression (COMMA expression)*
              ;

primaryExpression: operand #operandExpression
                 | primaryExpression selector #memberAccessor
                 | primaryExpression index #subIndex
                 | primaryExpression arguments #functionCall
                 | appendExpression #appendCall
                 | lengthExpression #lenCall
                 | capExpression #capCall
                 ;

operand: literal #literalOperand
       | IDENTIFIER #identifierOperand
       | LEFTPARENTHESIS expression RIGHTPARENTHESIS #expressionOperand
       ;

literal: INTLITERAL #intLiteral
       | FLOATLITERAL #floatLiteral
       | RUNELITERAL #runeLiteral
       | RAWSTRINGLITERAL #rawStringLiteral
       | INTERPRETEDSTRINGLITERAL #interpretedStringLiteral
       ;

index: LEFTBRACKET expression RIGHTBRACKET
     ;

arguments: LEFTPARENTHESIS expressionList* RIGHTPARENTHESIS
         ;

selector: DOT IDENTIFIER
        ;

appendExpression: APPEND LEFTPARENTHESIS expression COMMA expression RIGHTPARENTHESIS
                ;

lengthExpression: LEN LEFTPARENTHESIS expression RIGHTPARENTHESIS
                ;

capExpression: CAP LEFTPARENTHESIS expression RIGHTPARENTHESIS
             ;

statementList: statement*
             ;

block: LEFTCURLYBRACE statementList RIGHTCURLYBRACE
     ;

statement: PRINT LEFTPARENTHESIS expressionList? RIGHTPARENTHESIS SEMICOLON #printStatement
         | PRINTLN LEFTPARENTHESIS expressionList? RIGHTPARENTHESIS SEMICOLON #printlnStatement
         | RETURN expression? SEMICOLON #returnStatement
         | BREAK SEMICOLON #breakStatement
         | CONTINUE SEMICOLON #continueStatement
         | simpleStatement SEMICOLON #simpleStatementStatement
         | block SEMICOLON #blockStatement
         | switch SEMICOLON #switchStatement
         | ifStatement SEMICOLON #ifStatementStatement
         | loop SEMICOLON #loopStatement
         | typeDecl #typeDeclStatement
         | variableDecl #variableDeclStatement
         ;

simpleStatement: expression (POSTINC | POSTDEC)? #expressionSimpleStatement
               | assignmentStatement #assignmentSimpleStatement
               | left=expressionList WALRUS right=expressionList #walrusDeclaration
               ;

assignmentStatement: left=expressionList EQUALS right=expressionList #normalAssignment
                   | left=expression (IADD|IAND|ISUB|IOR|IMUL|IXOR|ILEFTSHIFT|IRIGHTSHIFT|IANDXOR|IMOD|IDIV) right=expression #inPlaceAssignment
                   ;

ifStatement: IF expression block
           | IF expression block ELSE (block|ifStatement)
           | IF simpleStatement SEMICOLON expression block
           | IF simpleStatement SEMICOLON expression block ELSE ifStatement
           | IF simpleStatement SEMICOLON expression block ELSE block
           ;

loop: FOR block
    | FOR expression block
    | FOR simpleStatement SEMICOLON expression SEMICOLON simpleStatement block
    | FOR simpleStatement SEMICOLON SEMICOLON simpleStatement block
    ;

switch: SWITCH simpleStatement SEMICOLON expression LEFTCURLYBRACE expressionCaseClauseList RIGHTCURLYBRACE
      | SWITCH expression LEFTCURLYBRACE expressionCaseClauseList RIGHTCURLYBRACE
      | SWITCH simpleStatement SEMICOLON LEFTCURLYBRACE expressionCaseClauseList RIGHTCURLYBRACE
      | SWITCH LEFTCURLYBRACE expressionCaseClauseList RIGHTCURLYBRACE
      ;

expressionCaseClauseList: (expressionCaseClause expressionCaseClauseList)*
                        ;

expressionCaseClause: expressionSwitchCase COLON statementList
                    ;

expressionSwitchCase: CASE expressionList
                    | DEFAULT
                    ;
// string_lit             = raw_string_lit | interpreted_string_lit .
// raw_string_lit         = "`" { unicode_char | newline } "`" .
// interpreted_string_lit = `"` { unicode_value | byte_value } `"` .

COMMENT: '//' ~('\n'|'\r')* '\r'? '\n' -> skip;
MULTILINE_COMMENT: '/*'.*?'*/' -> skip;
LEFTPARENTHESIS: '(';
RIGHTPARENTHESIS: ')';
DEFAULT: 'default';
CASE: 'case';
COLON: ':';
SWITCH: 'switch';
FOR: 'for';
IF: 'if';
ELSE: 'else';
IDIV: '/=';
IMOD: '%=';
IANDXOR: '&^=';
ILEFTSHIFT: '<<=';
IRIGHTSHIFT: '>>=';
IXOR: '^=';
IMUL: '*=';
IOR: '|=';
ISUB: '-=';
IAND: '&=';
IADD: '+=';
POSTINC: '++';
POSTDEC: '--';
CONTINUE: 'continue';
WALRUS: ':=';
BREAK: 'break';
RETURN: 'return';
PRINTLN: 'println';
PRINT: 'print';
CAP: 'cap';
LEN: 'len';
APPEND: 'append';
DOT: '.';
WHITESPACE: [ \r\n\t]+ -> skip;
NEWLINE: [\n];
// TODO: make these strings actually good.
INTERPRETEDSTRINGLITERAL: '"' ((.)*? | NEWLINE) '"';
RAWSTRINGLITERAL: '`'(.)*?'`';
ESCAPEDSEQUENCES: ('\\n'|'\\r'|'\\t'|'\\v');
RUNELITERAL: '\'' ([a-zA-Z]|ESCAPEDSEQUENCES) '\'';

FLOATLITERAL: [0-9]+ '.' [0-9]+;
NOT: '!';
OR: '||';
AND: '&&';
GREATERTHANEQUAL: '>=';
LESSTHANEQUAL: '<=';
GREATERTHAN: '>';
LESSTHAN: '<';
NEGATION: '!=';
COMPARISON: '==';
CARET: '^';
PIPE: '|';
MINUS: '-';
PLUS: '+';
AMPERSANDCARET: '&^';
AMPERSAND: '&';
RIGHTSHIFT: '>>';
LEFTSHIFT: '<<';
MOD: '%';
DIV: '/';
TIMES: '*';
LEFTCURLYBRACE: '{';
RIGHTCURLYBRACE: '}';
STRUCT: 'struct';
INTLITERAL: [0-9]+;
LEFTBRACKET: '[';
RIGHTBRACKET: ']';
COMMA: ',';
FUNC: 'func';
TYPE: 'type';
SEMICOLON: ';';
VAR: 'var';
EQUALS: '=';
PACKAGE: 'package';
IDENTIFIER: [a-zA-Z][a-zA-Z0-9]*;
