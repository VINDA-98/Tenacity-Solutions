package response

import "github.com/VINDA-98/Tenacity-Solutions/server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
