package main

import (
	"aws-golang-terraform-colonies/functions/libs"
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/aws/aws-lambda-go/lambda"
)

// Colony metadata format: XXX-XXXXXXXXXX
var colonyRegExp = regexp.MustCompile(`[0-9A-Z]{6}\-[0-9A-Z]{13}`)
var planetRegExp = regexp.MustCompile(`[0-9A-Za-z]{3,20}`)
var corporationRegExp = regexp.MustCompile(`[0-9A-Z]{3}\-[0-9A-Z]{10}`)

/*
Handler is our Lambda handler invoked by the `lambda.Start` function call.
*/
func Handler(ctx context.Context, req libs.Request) (libs.Response, error) {
	var buffer bytes.Buffer
	colony := req.PathParameters["id"]

	// Validation for request metadata. Return 404 if fails.
	if !colonyRegExp.MatchString(colony) {
		log.Print("[REGEX FAILED]: ", req)
		return libs.WrapError(http.StatusNotFound, "Bad values in query and/or request")
	}

	// @todo: Insert data layer persistence logic here.

	clx := libs.Colony{
		ColonyID:    "ABCDEF-0123456789XYZ",
		Planet:      "SIRIUS",
		Corporation: "117-TETRISCORP",
		Coords:      "100.2342356234, -345.7823425345",
	}

	// Marshal response, send back through API.
	body, err := json.Marshal(map[string]interface{}{
		"message":      "Colony read successfully",
		"transmission": clx,
	})
	if err != nil {
		log.Print("[MARSHAL FAILED]: ", err)
		return libs.WrapError(http.StatusNotFound, "Bad JSON marshal task")
	}

	// Clean up JSON response before returning to consumer.
	json.HTMLEscape(&buffer, body)
	resp := libs.Response{
		StatusCode:      http.StatusOK,
		IsBase64Encoded: false,
		Body:            buffer.String(),
		Headers: map[string]string{
			"Content-Type":       "application/json",
			"X-Athena-API-Reply": "read-handler",
		},
	}

	return resp, nil
}

/*
Lambda handler entry point.
*/
func main() {
	lambda.Start(Handler)
}
