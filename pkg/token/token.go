// Package token defines constants for lexer tokens.

package token

// TokenType represents the type of a token.
type TokenType string

const (
	// Special tokens
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers and literals
	IDENT TokenType = "IDENT" // variables, function names, etc.
	INT   TokenType = "INT"   // integers

	// Operators and delimiters
	ASSIGN    TokenType = "="
	PLUS      TokenType = "+"
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	// Parentheses and braces
	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
)
