package test

import (
	"github.com/VirgilZhao/paynow"
	"testing"
)

func TestGenerateWithCRCLengthLessThan4(t *testing.T) {
	val := paynow.GeneratePayNowString(paynow.Options{
		UEN:             "12345678",
		Editable:        false,
		Expiry:          "20260304",
		CompanyName:     "testcompany",
		Amount:          "0.10",
		ReferenceNumber: "testordernumber12345678",
	})
	t.Log(val)
}

func TestGenerateWithCRCLengthMeet4(t *testing.T) {
	val := paynow.GeneratePayNowString(paynow.Options{
		UEN:             "123456789",
		Editable:        false,
		Expiry:          "20260304",
		CompanyName:     "testcompany",
		Amount:          "0.99",
		ReferenceNumber: "testordernumber12345678",
	})
	t.Log(val)
}

func TestCompanyNameOutOfLength(t *testing.T) {
	val := paynow.GeneratePayNowString(paynow.Options{
		UEN:             "123456789",
		Editable:        false,
		Expiry:          "20260304",
		CompanyName:     "testcompanyoverlength1234567890",
		Amount:          "0.99",
		ReferenceNumber: "testordernumber12345678",
	})
	t.Log(val)
}

func TestUENWithSuffix(t *testing.T) {
	val := paynow.GeneratePayNowString(paynow.Options{
		UEN:             "123456789001",
		Editable:        false,
		Expiry:          "20260304",
		CompanyName:     "testcompany",
		Amount:          "0.99",
		ReferenceNumber: "testordernumber12345678",
	})
	t.Log(val)
}
