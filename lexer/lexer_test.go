package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "a = 1"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IDENTIFIER, "a"},
		{token.ASSIGN, "="},
		{token.INT, "1"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tt.expectedType != tok.Type {
			t.Fatalf(
				"tests[%d] - tokentype wrong, expected=%q, got=%q",
				i, tt.expectedType, tok.Type,
			)
		}
		if tt.expectedLiteral != tok.Literal {
			t.Fatalf(
				"tests[%d] - literal wrong, expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal,
			)
		}
	}

}

func TestCodeBlock(t *testing.T) {
	input := `let five = 5;
let ten = 10;	
let add = fn(x, y) {
  x + y;
};
1 == 1
3 != 4
let result = add(five, ten);
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.INT, "1"},
		{token.EQ, "=="},
		{token.INT, "1"},
		{token.INT, "3"},
		{token.NEQ, "!="},
		{token.INT, "4"},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tt.expectedType != tok.Type {
			t.Fatalf(
				"tests[%d] - tokentype wrong, expected=%q, got=%q",
				i, tt.expectedType, tok.Type,
			)
		}
		if tt.expectedLiteral != tok.Literal {
			t.Fatalf(
				"tests[%d] - literal wrong, expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal,
			)
		}
	}

}

func TestIsLetter(t *testing.T) {
	res := isLetter('T')
	if res != true {
		t.Fatalf("T is not letter")
	}
}
