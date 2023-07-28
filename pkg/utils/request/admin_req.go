package request

import "time"

type ReqSalesReport struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	//Pagination utils.Pagination `json:"pagination"`
}
