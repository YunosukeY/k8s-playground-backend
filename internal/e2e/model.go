package e2e

type RawMessage struct {
	From string
	To   []string
	Data string
}
type Message struct {
	Raw RawMessage
}
type Messages struct {
	Items []Message `json:"items"`
}
