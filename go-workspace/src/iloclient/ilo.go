package iloclient

import (
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
	//Package resty provides simple HTTP and REST client for Go inspired by Ruby rest-client.
	//Reference for resty: https://pkg.go.dev/github.com/go-resty/resty
)

// Interface are the custom type that is used to specify a set of one or more method signatures which are allowed to create a variable of an
// interface type and this variable can be assigned with a concrete type value that has the methods the interface requires.

type ILOClient interface {
	Health() bool
	//function to get the random UUID from the server out of all the available data in JSON format
	GetRandomUUID() (string, error)
	GetID() (int, error)
	GetSSN() (string, error)
}

type ILOStruct struct {
	rest *resty.Client //resty client object
}

func NewILOClient() ILOClient {
	return ILOStruct{
		rest: resty.New(),
	}
}

func (i ILOStruct) Health() bool {

	output := map[string]any{}
	var err error

	resp, err := i.rest.R().SetResult(&output).SetError(&err).Get("https://random-data-api.com/api/id_number/random_id_number")
	if err != nil {
		//not heathy
		fmt.Println(err)
		return false
	}

	if (resp.StatusCode() == 200 || resp.StatusCode() == 202) && output["status"] == "success" {
		return true
	}

	// healthy
	return true
}

func (i ILOStruct) GetRandomUUID() (string, error) {

	//creating a struct which contains different data members that is created for the corresponding data available on server
	type UUIDStruct struct {
		//available data format on server
		Id   int    `json:"id"`
		UUID string `json:"uid"`
		// Valid_us_ssn string `json:"valid_us_ssn"`
	}

	output := UUIDStruct{} //output is of type UUIDStruct
	// https://random-data-api.com/api/id_number/random_id_number
	resp, err := i.rest.R().SetResult(&output).Get("https://random-data-api.com/api/id_number/random_id_number") //random data api generator website link

	if err != nil {
		//if error was generated
		fmt.Println(err)
		return "", errors.New("error")
	}
	// if (resp.StatusCode()==200||resp.StatusCode()==202){
	// 	return output
	// }

	//if no error was generated then we check for the status code of the response to check
	//for the successful response

	if (resp.StatusCode() == 200 || resp.StatusCode() == 202) && len(output.UUID) != 0 {
		return output.UUID, nil
	}

	return "", errors.New("error")
}

// fuction to get the id
func (i ILOStruct) GetID() (int, error) {

	type UUIDStruct struct {
		//available data format on server
		Id           int    `json:"id"`
		UUID         string `json:"uid"`
		Valid_us_ssn string `json:"valid_us_ssn"`
	}

	output := UUIDStruct{} //output is of type UUIDStruct
	// https://random-data-api.com/api/id_number/random_id_number
	resp, err := i.rest.R().SetResult(&output).Get("https://random-data-api.com/api/id_number/random_id_number") //random data api generator website link

	if err != nil {
		//if error was generated
		fmt.Println(err)
		return -1, errors.New("error")
	}
	// if (resp.StatusCode()==200||resp.StatusCode()==202){
	// 	return output
	// }

	//if no error was generated then we check for the status code of the response to check
	//for the successful response

	if (resp.StatusCode() == 200 || resp.StatusCode() == 202) && len(output.UUID) != 0 {
		return output.Id, nil
	}

	return 0, errors.New("error")
}

func (i ILOStruct) GetSSN() (string, error) {

	//creating a struct which contains different data members that is created for the corresponding data available on server
	type SSNStruct struct {
		//available data format on server
		Id           int    `json:"id"`
		UUID         string `json:"uid"`
		Valid_us_ssn string `json:"valid_us_ssn"`
	}

	output := SSNStruct{} //output is of type UUIDStruct
	// https://random-data-api.com/api/id_number/random_id_number
	resp, err := i.rest.R().SetResult(&output).Get("https://random-data-api.com/api/id_number/random_id_number") //random data api generator website link

	if err != nil {
		//if error was generated
		fmt.Println(err)
		return "", errors.New("error")
	}
	// if (resp.StatusCode()==200||resp.StatusCode()==202){
	// 	return output
	// }

	//if no error was generated then we check for the status code of the response to check
	//for the successful response

	if (resp.StatusCode() == 200 || resp.StatusCode() == 202) && len(output.UUID) != 0 {
		return output.Valid_us_ssn, nil
	}

	return "", errors.New("error")
}
