package dto

type OrderLineStruct struct {
	OrderId   int     `json:"order_id"`
	ProductId int     `json:"product_id"`
	Qty       float64 `json:"qty"`
	Price     float64 `json:"price"`
}

type OrderCreateDto struct {
	CustomerId    uint64            `json:"customer_id" form:"customer_id" binding:"required"`
	UserId        uint64            `json:"user_id" form:"user_id" binding:"required"`
	EmpId         uint64            `json:"car_date" form:"emp_id"`
	EmpId2        uint64            `json:"emp2_id" form:"emp2_id"`
	RouteId       uint64            `json:"route_id" form:"route_id"`
	CarId         uint64            `json:"car_id" form:"car_id"`
	PaymentTypeId uint64            `json:"payment_type_id" form:"payment_type_id"`
	CompanyId     uint64            `json:"company_id" form:"company_id"`
	BranchId      uint64            `json:"branch_id"`
	DataList      []OrderLineStruct `json:"data"`
	Discount      float64           `json:"discount"`
	RouteCode     string            `json:"route_code"`
	RunNo         string            `json:"runno"`
	IssueId       uint64            `json:"issue_id"`
}
