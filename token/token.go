package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"

	LET   = "LET"
	PRINT = "PRINT"
)

var keywords = map[string]Type{
	"let":   LET,
	"print": PRINT,
}

func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
