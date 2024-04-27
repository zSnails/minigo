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

declType: LEFTPARENTHESIS declType RIGHTPARENTHESIS #nestedType
        | IDENTIFIER #identifierDeclType
        | sliceDeclType #sliceType
        | arrayDeclType #arrayType
        | structDeclType #structType
        ;

sliceDeclType: LEFTBRACKET RIGHTBRACKET declType
             ;
arrayDeclType: LEFTBRACKET INTLITERAL RIGHTBRACKET declType
             ;

structDeclType: STRUCT LEFTCURLYBRACE structMemDecls? RIGHTCURLYBRACE
              ;

structMemDecls: singleVarDeclNoExps SEMICOLON (singleVarDeclNoExps SEMICOLON)*
              ;

identifierList: IDENTIFIER (COMMA IDENTIFIER)*
              ;

expression: left=expression (LESSTHAN | GREATERTHAN | LESSTHANEQUAL | GREATERTHANEQUAL | COMPARISON | NEGATION) right=expression #comparison
          | left=expression (AND | OR) right=expression #booleanOperation
          | left=expression (TIMES | DIV | MOD) right=expression #operationPrecedence1
          | left=expression (LEFTSHIFT | RIGHTSHIFT | AMPERSAND | AMPERSANDCARET | PLUS | MINUS | PIPE | CARET) right=expression #operationPrecedence2
          | primaryExpression #expressionPrimaryExpression
          // | PLUS expression #plusExpression
          // | MINUS expression #minusExpression
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

appendExpression: APPEND LEFTPARENTHESIS slice=expression COMMA value=expression RIGHTPARENTHESIS
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

simpleStatement: expression #expressionSimpleStatement
               | expression POSTINC #expressionPostInc
               | expression POSTDEC #expressionPostDec
               | assignmentStatement #assignmentSimpleStatement
               | left=expressionList WALRUS right=expressionList #walrusDeclaration
               ;

assignmentStatement: left=expressionList EQUALS right=expressionList #normalAssignment
                   | left=expression (IADD|IAND|ISUB|IOR|IMUL|IXOR|ILEFTSHIFT|IRIGHTSHIFT|IANDXOR|IMOD|IDIV) right=expression #inPlaceAssignment
                   ;

ifStatement: IF expression block #ifSingleExpression
           | IF expression block ELSE ifStatement #ifElseIf
           | IF expression block ELSE block #ifElseBlock
           | IF simpleStatement SEMICOLON expression block #ifSimpleNoElse
           | IF simpleStatement SEMICOLON expression block ELSE ifStatement #ifSimpleElseIf
           | IF simpleStatement SEMICOLON expression block ELSE block #ifSimpleElseBlock
           ;

loop: FOR block #infiniteFor
    | FOR expression block #whileFor
    | FOR first=simpleStatement SEMICOLON expression? SEMICOLON last=simpleStatement block #threePartFor
    ;

switch: SWITCH simpleStatement SEMICOLON expression LEFTCURLYBRACE expressionCaseClauseList RIGHTCURLYBRACE #simpleStatementSwitchExpression
      | SWITCH LEFTCURLYBRACE expressionCaseClauseList RIGHTCURLYBRACE #normalSwitch
      | SWITCH expression LEFTCURLYBRACE expressionCaseClauseList RIGHTCURLYBRACE #normalSwitchExpression
      | SWITCH simpleStatement SEMICOLON LEFTCURLYBRACE expressionCaseClauseList RIGHTCURLYBRACE #simpleStatementSwitch
      ;

expressionCaseClauseList: (expressionCaseClause expressionCaseClauseList)*
                        ;

expressionCaseClause: expressionSwitchCase COLON statementList
                    ;

expressionSwitchCase: CASE expressionList #switchCaseBranch
                    | DEFAULT #switchDefaultBranch
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

FLOATLITERAL: '-'?[0-9]+ '.' [0-9]+;
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
INTLITERAL: '-'?[0-9]+;
LEFTBRACKET: '[';
RIGHTBRACKET: ']';
COMMA: ',';
FUNC: 'func';
TYPE: 'type';
SEMICOLON: ';';
VAR: 'var';
EQUALS: '=';
PACKAGE: 'package';
IDENTIFIER: [a-zA-Z_][a-zA-Z0-9_]*;
//IDENTIFIER: '_'*?[a-zA-Z]('_'|[a-zA-Z0-9])*;
