// Code generated by "stringer -type Token"; DO NOT EDIT

package parser

import "fmt"

const _Token_name = "WhitespaceNewLineTabIntegerTextEOFSlashUnknownOpeningBracketClosingBracketOpeningParentheseClosingParentheseColonDotEqualGoroutinePanicRecoveredRuntimeErrorCreatedByRunningIOWaitChanReceiveSendSyscallPointerFrame"

var _Token_index = [...]uint8{0, 10, 17, 20, 27, 31, 34, 39, 46, 60, 74, 91, 108, 113, 116, 121, 130, 135, 144, 151, 156, 163, 165, 172, 178, 182, 189, 193, 200, 207, 212}

func (i Token) String() string {
	if i < 0 || i >= Token(len(_Token_index)-1) {
		return fmt.Sprintf("Token(%d)", i)
	}
	return _Token_name[_Token_index[i]:_Token_index[i+1]]
}
