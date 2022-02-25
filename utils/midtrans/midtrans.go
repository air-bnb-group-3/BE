package midtrans

import (
	"github.com/labstack/gommon/log"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func InitConnection() coreapi.Client {
	var c = coreapi.Client{}
	c.New("SB-Mid-server-6xAOB7aeOdPQL9EEbmV9hnqE", midtrans.Sandbox)
	return c
}

func CreateTransaction(core coreapi.Client) *coreapi.ChargeResponse {
	req := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBCAKlikpay,
		BCAKlikPay: &coreapi.BCAKlikPayDetails{
			Desc: "test to payment process",
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "id booking", /*id booking*/
			GrossAmt: 200000,       /*total price di booking*/
		},
		Items: &[]midtrans.ItemDetails{
			{Name: "Purchase 1", Price: 100000, Qty: 2},
		},
	}
	apiRes, err := core.ChargeTransaction(req)
	if err != nil {
		log.Warn("Payment Error")
	}
	return apiRes
}
