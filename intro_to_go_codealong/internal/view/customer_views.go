package view

import "encoding/json"

type CustomerCreateReq struct {
	FirstName *string `json:"firstName"`
	MiddleName *string `json:"middleName"`
	LastName *string `json:"lastName"`
	Age *json.Number `json:"age"`
}
