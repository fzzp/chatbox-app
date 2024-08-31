package socket

type ActionType int

const (
	HelloType ActionType = iota
)

type WsPayload struct {
	Action ActionType `json:"action"`
	Data   any        `json:"data"`
}
