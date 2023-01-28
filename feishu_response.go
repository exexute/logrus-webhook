package logrusWebhook

type FeiShuResponse struct {
	Extra         interface{} `json:"Extra"`
	StatusCode    int         `json:"StatusCode"`
	StatusMessage string      `json:"StatusMessage"`

	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
