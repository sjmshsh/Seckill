package response

// Response 通用Response
type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}
