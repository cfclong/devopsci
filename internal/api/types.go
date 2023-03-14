package api

// Result ..
type Result interface{}

type SuccessResult struct {
	IsSuccess bool        `json:"IsSuccess"`
	Data      interface{} `json:"Data,omitempty"`
	Message   string      `json:"Message,omitempty"`
	ErrMsg    string      `json:"ErrMsg,omitempty"`
}

type ErrorResult struct {
	IsSuccess bool   `json:"IsSuccess"`
	ErrCode   string `json:"ErrCode,omitempty"`
	ErrMsg    string `json:"ErrMsg,omitempty"`
	ErrDetail string `json:"ErrDetail,omitempty"`
}

func NewResult(isSuccess bool, data interface{}, errMsg string) Result {
	return &SuccessResult{IsSuccess: isSuccess, Data: data, ErrMsg: errMsg}
}

func NewErrorResult(errCode, errMsg, errDetail string) Result {
	return &ErrorResult{
		IsSuccess: false,
		ErrCode:   errCode,
		ErrMsg:    errMsg,
		ErrDetail: errDetail,
	}
}

func NewSuccessResult(data ...interface{}) Result {
	var result interface{}
	if len(data) > 0 {
		result = data[0]
	}
	return &SuccessResult{
		IsSuccess: true,
		Data:      result,
	}
}
