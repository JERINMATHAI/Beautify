package request

import "time"

type ReqSalesReport struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
	//Pagination utils.Pagination `json:"pagination"`
}

type ApproveReturnRequest struct {
	ReturnID   uint   `json:"return_id"`
	OrderID    uint   `json:"order_id"`
	UserID     uint   `json:"user_id"`
	OrderTotal uint   `json:"-"`
	IsApproved bool   `json:"is_approved"`
	Comment    string `json:"comment"`
}
