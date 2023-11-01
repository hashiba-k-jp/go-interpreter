// token/token.go

package token

type TokenType string

type Token struct{
    Type    TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF     = "EOF"

    // identifier + literal
    IDENT   = "IDENT" // add, foobar, ...
    INT     = "INT" // 1, 2, 3, ...
    STRING  = "STRING"

    // operand
    ASSIGN      = "="
    PLUS        = "+"
    MINUS       = "-"
    BANG        = "!"
    ASTERISK    = "*"
    SLASH       = "/"
    LT          = "<"
    GT          = ">"

    // delimiter
    COMMA       = "."
    SEMICOLON   = ";"
    LPAREN      = "("
    RPAREN      = ")"
    LBRACE      = "{"
    RBRACE      = "}"
    LBRACKET    = "[" // added
    RBRACKET    = "]" // added

    // keywords
    FUNCTION    = "FUNCTION"
    LET         = "LET"
    TRUE        = "TRUE"
    FALSE       = "FALSE"
    IF          = "IF"
    ELSE        = "ELSE"
    RETURN      = "RETURN"

    EQ          = "=="
    NOT_EQ      = "!="
)

// map[<TYPE>]<NAME>
// Dose it seem to be similar the dict type of python ...?
var keywords = map[string]TokenType{
    "fn"    : FUNCTION,
    "let"   : LET,
    "true"  : TRUE,
    "false" : FALSE,
    "if"    : IF,
    "else"  : ELSE,
    "return": RETURN,
}

func LookupIdent(ident string) TokenType{
    if tok, ok := keywords[ident]; ok{
        return tok
    }
    return IDENT
}
