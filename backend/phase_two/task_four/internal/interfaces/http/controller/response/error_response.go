package response

type ErrorResponse struct {
	Code    int         `json:"code"`              // 错误码，可自定义
	Message string      `json:"message"`           // 错误描述
	Details interface{} `json:"details,omitempty"` // 可选，具体错误详情，例如字段验证错误
}

func NewErrorResponse(code int, message string, details interface{}) *ErrorResponse {
	return &ErrorResponse{
		Code:    code,
		Message: message,
		Details: details,
	}

}
