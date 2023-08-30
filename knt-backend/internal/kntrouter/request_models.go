package kntrouter

type ErrorModel struct {
	Err string `json:"error"`
}

type IdResponse struct {
	Id int64 `json:"id"`
}

type SpentFormat struct {
	Cost int `json:"moneySpent"`
}

type WebhookFormat struct {
	Balance   int    `json:"balance" validate:"required"`
	VunetId   string `json:"vunetid" validate:"required"`
	Reference string `json:"reference"`
}
