package tstore

// Msg is a message
type Msg struct {
	ID		string	`json:"id"`
	PID		string	`json:"pid"`
	From 	string	`json:"from"`
	To 		string	`json:"to"`
	Time 	int64	`json:"time"`
	Topic	string	`json:"topic"`
	Content string	`json:"content"`
}

// MsgStorer defines the Store(msg) method to store a Msg
type MsgStorer interface {
	Store(msg *Msg) error
}

// MsgFetcher defines methods for retreiving messages from a store
type MsgFetcher interface {
	Fetch(id *string) error
	FetchLatest(topic string, maxDepth int) error
}

// MsgSearcher defines the Search(str) method to search a message store 
type MsgSearcher interface {
	Search(str *string) []Msg
}
