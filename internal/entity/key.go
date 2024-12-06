package entity

type (
	Request struct {
		Text string `json:"text"`
	}

	Response struct {
		TextDecrypted string `json:"text_decrypted"`
		TextEncrypted string `json:"text_encrypted"`
	}
)

func NewResponse(textDecrypted string, textEncrypted string) Response {
	return Response{
		TextDecrypted: textDecrypted,
		TextEncrypted: textEncrypted,
	}
}
