package mapper

import (
	"intro_to_go_codealong/internal/domain"
	"intro_to_go_codealong/internal/view"
)

type Customer struct {}

func (c *Customer) CreateReqToDomain(v *view.CustomerCreateReq) *domain.Customer {
	int64Age, _ := v.Age.Int64()

	return &domain.Customer{
		FirstName:  *v.FirstName,
		MiddleName: v.MiddleName,
		LastName:   *v.LastName,
		Age:        int(int64Age),
	}
}
