package common

type Error struct {
	Path   string
	Source []string
	Line   int
	Column int
	Msg    string
	Msg2   string
}

func NewError(filePath string, source []string, line, column int, msg, msg2 string) *Error {
	return &Error{
		Path:   filePath,
		Source: source,
		Line:   line,
		Column: column,
		Msg:    msg,
		Msg2:   msg2,
	}
}
