package parser

func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func isInteger(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isHexa(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || isInteger(ch)
}
