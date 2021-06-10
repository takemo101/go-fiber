package object

// SuggestInput input form to service
type SuggestInput struct {
	requestID   uint
	suggesterID uint
	message     string
}

func NewSuggestInput(
	requestID uint,
	message string,
) SuggestInput {
	return SuggestInput{
		requestID: requestID,
		message:   message,
	}
}

func (o SuggestInput) GetRequestID() uint {
	return o.requestID
}

func (o SuggestInput) GetMessage() string {
	return o.message
}
