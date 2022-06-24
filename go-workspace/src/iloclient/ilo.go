package iloclient

import (
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type ILOClient interface {
	Health() bool
	GetRandomUUID() (string, error)
}

type ILOStruct struct {
	rest *resty.Client
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
		fmt.Println(err)
		return false
	}

	if (resp.StatusCode() == 200 || resp.StatusCode() == 202) && output["status"] == "success" {
		return true
	}

	// not healthy
	return false
}

func (i ILOStruct) GetRandomUUID() (string, error) {

	type UUIDStruct struct {
		Id           int    `json:"id"`
		UUID         string `json:"uid"`
		Valid_us_ssn string `json:"valid_us_ssn"`
	}

	output := UUIDStruct{}
	// https://random-data-api.com/api/id_number/random_id_number
	resp, err := i.rest.R().SetResult(&output).Get("https://random-data-api.com/api/id_number/random_id_number")
	if err != nil {
		fmt.Println(err)
		return "", errors.New("error")
	}
	// if (resp.StatusCode()==200||resp.StatusCode()==202){
	// 	return output
	// }

	if (resp.StatusCode() == 200 || resp.StatusCode() == 202) && len(output.UUID) != 0 {
		return output.UUID, nil
	}

	return "", errors.New("error")
}
