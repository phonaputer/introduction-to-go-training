package cli

import (
	"customer_app/internal/customer/logic/cerr"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strconv"
)

type CustomerCommandValidator interface {
	ValidateGet(args []string) (int, error)
	ValidateCreate(args []string) (*customerCreateJSON, error)
}

func NewCustomerCommandValidator() CustomerCommandValidator {
	return &customerCommandValidatorImpl{}
}

type customerCommandValidatorImpl struct{}

func (c *customerCommandValidatorImpl) ValidateGet(args []string) (int, error) {
	if len(args) != 1 {
		return 0, cerr.NewUserMsg("unexpected number of arguments for get customer")
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		return 0, cerr.SetUserMsg(err, "customer id is not a valid integer")
	}

	return id, nil
}

func (c *customerCommandValidatorImpl) ValidateCreate(args []string) (*customerCreateJSON, error) {
	if len(args) != 1 {
		return nil, cerr.NewUserMsg("unexpected number of arguments for create customer")
	}

	var err error
	var result customerCreateJSON

	file, err := validateFilePath(err, args[0])
	if file != nil {
		defer file.Close()
	}

	err = jsonReader(err, file, &result, "invalid customer JSON")
	err = notNull(err, result.FirstName, "first_name is required")
	err = notNull(err, result.LastName, "last_name is required")
	err = notNull(err, result.Age, "age is required")
	err = stringLength(err, result.FirstName, 1, 10, `first_name has invalid length`)
	err = stringLength(err, result.MiddleName, 1, 10, `middle_name has invalid length`)
	err = stringLength(err, result.LastName, 1, 10, `last_name has invalid length`)

	if err != nil {
		return nil, cerr.SetUserMsg(err, err.Error())
	}

	return &result, nil
}

func validateFilePath(err error, path string) (io.ReadCloser, error) {
	if err != nil {
		return nil, err
	}

	file, err := os.Open(path)
	if err != nil {
		logrus.WithError(err).Debug("error opening file")
		return nil, errors.New("could not open input file")
	}

	return file, nil
}

func jsonReader(err error, r io.Reader, result interface{}, msg string) error {
	if err != nil {
		return err
	}

	err = json.NewDecoder(r).Decode(result)
	if err != nil {
		logrus.WithError(err).Debug("error parsing JSON")
		return errors.New(msg)
	}

	return nil
}

func notNull[T any](err error, value *T, msg string) error {
	if err != nil {
		return err
	}

	if value == nil {
		return errors.New(msg)
	}

	return nil
}

func stringLength(err error, value *string, min int, max int, msg string) error {
	if err != nil || value == nil {
		return err
	}

	length := len(*value)
	if length < min || max < length {
		return errors.New(msg)
	}

	return nil
}
