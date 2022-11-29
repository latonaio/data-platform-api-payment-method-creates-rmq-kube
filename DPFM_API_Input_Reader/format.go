package dpfm_api_input_reader

import (
	"data-platform-api-payment-method-creates-rmq-kube/DPFM_API_Caller/requests"
)

func (sdc *SDC) ConvertToPaymentMethod() *requests.PaymentMethod {
	data := sdc.PaymentMethod
	return &requests.PaymentMethod{
		PaymentMethod: data.PaymentMethod,
	}
}

func (sdc *SDC) ConvertToPaymentMethodText() *requests.PaymentMethodText {
	dataPaymentMethod := sdc.PaymentMethod
	data := sdc.PaymentMethod.PaymentMethodText
	return &requests.PaymentMethodText{
		PaymentMethod:     dataPaymentMethod.PaymentMethod,
		Language:          data.Language,
		PaymentMethodName: data.PaymentMethodName,
	}
}
