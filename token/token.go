package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN      = "="
	PLUS        = "+"
	PLUS_PLUS   = "++"
	PLUS_EQ     = "+="
	MINUS       = "-"
	MINUS_MINUS = "--"
	MINUS_EQ    = "-="
	ASTERISK    = "*"
	ASTERISK_EQ = "*="
	SLASH       = "/"
	SLASH_EQ    = "/="
	BANG        = "!"

	LT    = "<"
	LT_EQ = "<="
	GT    = ">"
	GT_EQ = ">="

	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	FOR      = "FOR"
)

var keywords = map[string]TokenType{
	"func":   FUNCTION,
	"let":    LET,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"for":    FOR,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
