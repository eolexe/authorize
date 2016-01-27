package authorize

import "encoding/json"

type Error struct {
	TransactionResponse *struct {
		Errors []struct {
			ErrorCode string
			ErrorText string
		}
	}
	Messages *Message `json:"messages"`
}

type Message struct {
	Code     string `json:"resultCode"`
	Messages []struct {
		Code string
		Text string
	} `json:"message"`
}

type Response struct {
	Messages       Message `json:"messages"`
	Raw            json.RawMessage
	ResponseStruct interface{}
	Err            error
}

func ParseResponse(r *Response, buff []byte) *Response {
	r.Raw = buff
	r.Err = json.Unmarshal(buff, r)

	if r.Err != nil {
		return r
	}

	if r.Messages.Code == "Error" {
		mErr := &Error{}
		r.Err = json.Unmarshal(r.Raw, mErr)
		if r.Err != nil {
			return r
		}

		r.Err = parseError(mErr)
		return r
	}

	r.Err = json.Unmarshal(buff, r.ResponseStruct)
	return r
}
