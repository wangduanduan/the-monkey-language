package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"let": LET,
	"fn":  FUNCTION,
}

func LookupIdent(Literal string) TokenType {

	if _, ok := keywords[Literal]; ok {
		return keywords[Literal]
	}

	return IDENT
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN = "ASSIGN"
	PLUS   = "PLUS"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
