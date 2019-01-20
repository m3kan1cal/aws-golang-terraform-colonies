package main

import (
	"aws-golang-terraform-colonies/functions/libs"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

	// Validation for request metadata. Return 400 if fails.
	if req.Headers["Content-Type"] != "application/json" {
		return libs.WrapError(http.StatusBadRequest, "Bad content type")
	}

	clx := new(libs.Colony)
	err := json.Unmarshal([]byte(req.Body), clx)
	if err != nil {
		log.Print("[ENTITY FAILED]: ", req, err)
		return libs.WrapError(http.StatusUnprocessableEntity, "Bad entity not processed")
	}

	// Validation for colony metadata. Return 400 if fails.
	if !colonyRegExp.MatchString(clx.Name) || !planetRegExp.MatchString(clx.Planet) ||
		!corporationRegExp.MatchString(clx.Corporation) {
		log.Print("[REGEX FAILED]: ", req)
		return libs.WrapError(http.StatusBadRequest, "Bad values in query and/or request")
	}
	if clx.Name == "" || clx.Planet == "" ||
		clx.Corporation == "" || clx.Coords == "" {
		log.Print("[VALUES FAILED]: ", req)
		return libs.WrapError(http.StatusBadRequest, "Bad values in query and/or request")
	}

	// @todo: Insert data layer persistence logic here.

	// Marshal response, send back through API.
	body, err := json.Marshal(map[string]interface{}{
		"message":      "Colony created successfully",
		"transmission": clx,
	})
	if err != nil {
		log.Print("[MARSHAL FAILED]: ", err)
		return libs.WrapError(http.StatusNotFound, "Bad JSON marshal task")
	}

	json.HTMLEscape(&buffer, body)
	resp := libs.Response{
		StatusCode:      http.StatusCreated,
		IsBase64Encoded: false,
		Body:            buffer.String(),
		Headers: map[string]string{
			"Content-Type":       "application/json",
			"X-Athena-API-Reply": "create-handler",
			"Location":           fmt.Sprintf("/colonies/%s", clx.Name),
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
