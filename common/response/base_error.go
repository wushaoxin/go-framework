package response

const defaultCode = -1

type CodeError struct {
    Code int    `json:"code"`
    Msg  string `json:"msg"`
    Data any    `json:"data"`
}

func NewCodeError(code int, msg string) *CodeError {
    return &CodeError{Code: code, Msg: msg}
}

func NewDefaultError(msg string) *CodeError {
    return NewCodeError(defaultCode, msg)
}

func NewServerError() *CodeError {
    return NewCodeError(defaultCode, "server error")
}

func (e *CodeError) Error() string {
    return e.Msg
}

func (e *CodeError) Body() *Body {
    return &Body{
        Success: false,
        Code:    e.Code,
        Msg:     e.Msg,
        Data:    e.Data,
    }
}

func (e *CodeError) WithData(data any) *CodeError {
    e.Data = data
    return e
}
