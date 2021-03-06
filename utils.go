package parser

func isWhitespace(ch rune) bool {
	return ch == ' '
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '·'
}

func isInteger(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isHexa(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || isInteger(ch)
}
