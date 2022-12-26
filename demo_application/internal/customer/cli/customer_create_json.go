package cli

type customerCreateJSON struct {
	FirstName  *string `json:"first_name"`
	MiddleName *string `json:"middle_name"`
	LastName   *string `json:"last_name"`
	Age        *int    `json:"age"`
}
