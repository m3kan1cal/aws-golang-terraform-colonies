package libs

import (
	"encoding/json"
	"log"
)

/*
WrapError is helper for returning friendly error responses to clients.
*/
func WrapError(status int, details string) (Response, error) {
	body := &ClientError{Error: details}
	b, err := json.Marshal(body)
	if err != nil {
		log.Print("[MARSHAL FAILED]: ", err)
	}

	return Response{
		StatusCode: status,
		Body:       string(b),
	}, nil
}
