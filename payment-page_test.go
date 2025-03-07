package paymentpage

import (
	"net/url"
	"reflect"
	"testing"
	"strings"
)

func TestPaymentPage(t *testing.T) {
	t.Parallel()

	signatureHandler := NewSignatureHandler("qwerty")
	paymentPage := NewPaymentPage(*signatureHandler)

	comparePaymentHost := "test.test"
	comparePaymentPath := "/pay"
	comparePaymentQuery := map[string]interface{}{
		"project_id":             "11",
		"payment_id":             "test_payment_id",
		"payment_currency":       "EUR",
		"payment_amount":         "1000",
		"some_future_bool_param": "true",
		"interface_type":         `{"id": 20}`,
		"signature":              "+NGChjO/3L6vSJkEUSXBDPJBSUuEu4rXw4wtAoXiTDATSMerNixVYKdh9Cg2jTXSu1Ez9R+LxX/ioWr70Tlxew==",
	}

	payment := NewPayment(11, "test_payment_id")
	payment.SetParam("payment_currency", "EUR")
	payment.SetParam("payment_amount", 1000)
	payment.SetParam("some_future_bool_param", true)

	paymentUrl := paymentPage.GetUrl(*payment)
	parsedUrl, err := url.Parse(paymentUrl)
	query, _ := url.ParseQuery(parsedUrl.RawQuery)

	if err != nil {
		t.Error(
			"For", "GetUrl",
			"expected", "valid URL",
			"got", paymentUrl,
		)
	}

	if parsedUrl.Host != comparePaymentHost {
		t.Error(
			"For", "GetUrl", payment,
			"expected", comparePaymentHost,
			"got", parsedUrl.Host,
		)
	}

	if parsedUrl.Path != comparePaymentPath {
		t.Error(
			"For", "GetUrl", payment,
			"expected", comparePaymentPath,
			"got", parsedUrl.Path,
		)
	}

	realQuery := map[string]interface{}{}

	for key, value := range query {
		realQuery[key] = value[0]
	}

	equal := reflect.DeepEqual(comparePaymentQuery, realQuery)

	if !equal {
		t.Error(
			"For", "GetUrl",
			"expected", comparePaymentQuery,
			"got", realQuery,
		)
	}

	paymentUrlEncrypted := paymentPage.GetEncryptedUrl(*payment, "123")
	parsedUrlEncrypted, err := url.Parse(paymentUrlEncrypted)

	queryFromEncrypted, _ := url.ParseQuery(parsedUrlEncrypted.RawQuery)

	if len(queryFromEncrypted) > 0 {
		t.Error(
			"For", "GetEncryptedUrl",
			"expected", "empty",
			"got", queryFromEncrypted,
		)
	}

	pathItems := strings.Split(parsedUrlEncrypted.Path, "/")

	if pathItems[1] != "11" {
		t.Error(
			"For", "GetEncryptedUrl",
			"expected", "11",
			"got", pathItems[1],
		)
	}
}
