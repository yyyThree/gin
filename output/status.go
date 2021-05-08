package output

import (
	"github.com/yyyThree/gin/output/code"
)

type Status struct {
	Code code.Code `json:"code,omitempty"`
	// A developer-facing Statusor message, which should be in English. Any
	Message string `json:"message,omitempty"`
	// A list of messages that carry the Statusor details.  There is a common set of
	// message types for APIs to use.
	Details []interface{} `json:"details,omitempty"`
}

func Error(code code.Code) *Status {
	return &Status{Code: code}
}

func ErrorWithMessage(code code.Code, message string) *Status {
	return &Status{Code: code, Message: message}
}

func (e *Status) WithDetails(data interface{}) *Status {
	if data == nil {
		return e
	}
	switch v := data.(type) {
	case *Status:
		tmp := &Status{
			Code:    code.Code(v.GetCode()),
			Message: v.GetMessage(),
			Details: v.GetDetails(),
		}
		e.Details = append(e.Details, tmp)
	default:
		e.Details = append(e.Details, data)
	}
	return e
}

func (e *Status) GetCode() int {
	return int(e.Code)
}

func (e *Status) GetMessage() string {
	if e.Message == "" {
		return e.Code.String()
	}
	return e.Message
}

func (e *Status) GetDetails() []interface{} {
	if e == nil {
		return nil
	}
	return e.Details
}

func (e *Status) Error() string {
	return e.GetMessage()
}
