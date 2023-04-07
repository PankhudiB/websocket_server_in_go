package model


type MessageType_A struct {
	Name  string `json:"name"`
	Place string `json:"place"`
}

type MessageType_B struct {
	Animal string `json:"animal"`
	Thing  string `json:"thing"`
}

type MessageWrapper struct {
	MessageType   string `json:"message_type"`
	MessageType_A `json:"content"`
}
