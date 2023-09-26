package models

type CommonResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type CommonError struct {
	Code      int      `json:"code"`
	ErrorData ApiError `json:"err_data"`
}

type ApiError struct {
	ErrorTitle   string `json:"err_title"`
	ErrorMessage string `json:"err_msg"`
}
