package request

type ComplaintOrder struct {
	Label           string `json:"label"`
	OrderNum        string `json:"order_num"`
	Contents        string `json:"contents"`
	ByComplainantId int    `json:"by_complainant_id"`
}
