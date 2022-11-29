package dpfm_api_output_formatter

type PaymentMethod struct {
	PaymentMethod     string            `json:"PaymentMethod"`
	PaymentMethodText PaymentMethodText `json:"PaymentMethodText"`
}

type PaymentMethodText struct {
	PaymentMethod     string `json:"PaymentMethod"`
	Language          string `json:"Language"`
	PaymentMethodName string `json:"PaymentMethodName"`
}
