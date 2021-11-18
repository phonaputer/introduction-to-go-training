package validator

import (
	"fmt"
	"intro_to_go_codealong/internal/itgerr"
	"intro_to_go_codealong/internal/view"
	"io"
	"net/http"
	"encoding/json"
)

type Customer struct {}

func (c *Customer) CreateRequest(r *http.Request) (*view.CustomerCreateReq, error) {
	var res view.CustomerCreateReq

	err := decodeJSON(nil, r.Body, &res)
	err = requiredString(err, res.FirstName, "firstName")
	err = requiredString(err, res.LastName, "lastName")
	err = requiredJSONNumber(err, res.Age, "age")
	err = jsonNumberToInt(err, res.Age, "age")

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func jsonNumberToInt(err error, value *json.Number, fieldName string) error {
	if err != nil {
		return err
	}

	_, err = value.Int64()
	if err != nil {
		return itgerr.WithKind(err, itgerr.KindInvalidInput, fmt.Sprintf("%s is not a valid integer", fieldName))
	}

	return nil
}

func requiredJSONNumber(err error, value *json.Number, fieldName string) error {
	if err != nil {
		return err
	}

	if value == nil {
		return itgerr.WithKind(nil, itgerr.KindInvalidInput, fmt.Sprintf("%s is required", fieldName))
	}

	return nil
}

func requiredString(err error, value *string, fieldName string) error {
	if err != nil {
		return err
	}

	if value == nil {
		return itgerr.WithKind(nil, itgerr.KindInvalidInput, fmt.Sprintf("%s is required", fieldName))
	}

	return nil
}

func decodeJSON(inputErr error, jsonReader io.Reader, res interface{}) error {
	if inputErr != nil {
		return inputErr
	}

	err := json.NewDecoder(jsonReader).Decode(&res)
	if err != nil {
		return itgerr.WithKind(err, itgerr.KindInvalidInput, "json body not valid")
	}

	return nil
}

