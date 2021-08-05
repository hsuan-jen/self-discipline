package request

type CheckSignature struct {
	Signature string `json:"signature" form:"signature"`
	Timestamp int16  `json:"timestamp" form:"timestamp"`
	Nonce     string `json:"nonce" form:"nonce"`
	Echostr   string `json:"echostr" form:"echostr"`
}
