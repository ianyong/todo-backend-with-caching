package api

const (
	Error       = MessageType(0)
	Warning     = MessageType(1)
	Information = MessageType(2)
	Success     = MessageType(3)
)

type MessageType uint

type StatusMessage struct {
	Content string      `json:"content"`
	Type    MessageType `json:"type"`
}

type StatusMessages []StatusMessage

func ErrorMessage(content string) StatusMessage {
	return StatusMessage{
		Content: content,
		Type:    Error,
	}
}

func WarningMessage(content string) StatusMessage {
	return StatusMessage{
		Content: content,
		Type:    Warning,
	}
}

func InformationMessage(content string) StatusMessage {
	return StatusMessage{
		Content: content,
		Type:    Information,
	}
}

func SuccessMessage(content string) StatusMessage {
	return StatusMessage{
		Content: content,
		Type:    Success,
	}
}
