package request

type Issue struct {
	UserId  uint64 `json:"user_id"`
	Content string `json:"content"`
}
