package entity

type (
	Request struct {
		Text string `json:"text"`
	}

	Response struct {
		Key           *string `json:"key,omitempty"`
		TextDecrypted *string `json:"text_decrypted,omitempty"`
		TextEncrypted *string `json:"text_encrypted,omitempty"`
	}
)

func NewKeyResponse(key *string) *Response {
	return &Response{
		Key: key,
	}
}

func NewResponse(textDecrypted, textEncrypted *string) Response {
	return Response{
		TextDecrypted: textDecrypted,
		TextEncrypted: textEncrypted,
	}
}
