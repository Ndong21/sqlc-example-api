package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// This function initiates a payment request
func RequestPayment(amount, currency, number, description, reference string) (string, error) {

	//load the api-key from the .env file, handle errors if they occur
	key := os.Getenv("key")

	baseUrl := os.Getenv("baseUrl")

	type Data struct {
		Amount             string `json:"amount,omitempty"`
		Currency           string `json:"currency,omitempty"`
		From               string `json:"from,omitempty"`
		Description        string `json:"description,omitempty"`
		External_reference string `json:"external_reference,omitempty"`
		Reference          string `json:"reference,omitempty"`
		Status             string `json:"status,omitempty"`
		Operator           string `json:"operator,omitempty"`
		Code               string `json:"code,omitempty"`
		Operator_reference string `json:"operator_reference,omitempty"`
	}

	data := Data{
		Amount:             amount,
		Currency:           currency,
		From:               number,
		Description:        description,
		External_reference: reference,
	}

	//convert the payment details into json
	//because the body of the request expects json data, catch any error that
	//may occur
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("error: %s", err.Error())
	} else {
		fmt.Println(string(jsonData))
	}

	//initialize an http client.
	// http.Client{} is a struct in Go's net/http package
	client := &http.Client{}

	//create an http request using http.NewRequest function from net/http

	url := fmt.Sprintf("%vcollect/", baseUrl)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("error occured:", err)
	}

	//adding custom headers to http.newrequest()
	//authorization and content type
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", key)

	//send an http request and return the response
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error occured", err)
	}

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	var result Data
	// decode the body of the response to a readable format using the sresult tructs
	error := json.NewDecoder(resp.Body).Decode(&result)
	if error != nil {
		log.Fatal(error)
	}

	//obtain the transaction reference from the response body
	//this reference will be used to check the status of the transaction
	return result.Reference, err
}
