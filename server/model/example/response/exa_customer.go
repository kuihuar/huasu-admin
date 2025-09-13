package response

import "github.com/kuihuar/huasu-admin/server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
