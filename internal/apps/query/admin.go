package query

type AdminTokenView struct {
	Token     string `json:"token"`
	Timestamp int64  `json:"timestamp"`
	Exp       int64  `json:"exp"`
}
