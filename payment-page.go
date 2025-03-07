package paymentpage

import (
	"strconv"
	"strings"
	"fmt"
	"net/url"
)

// Structure for build payment URL
type PaymentPage struct {
	// Signature handler for generate signature
	signatureHandler SignatureHandler
}

// Method builds payment URL
func (p *PaymentPage) GetUrl(baseUrl string, payment Payment) string {
	queryString := p.prepareQueryString(payment)

	return fmt.Sprintf("%s?%s", baseUrl, queryString)
}

// Method builds encrypted payment URL
func (p *PaymentPage) GetEncryptedUrl(baseUrl string, payment Payment, encryptionKey string) string {
	queryString := p.prepareQueryString(payment)
	parsedUrl, err := url.Parse(baseUrl)

	if err != nil {
		fmt.Println("Cant parse base url")
		return ""
	}

	pathWithQueryString := fmt.Sprintf("%s?%s", parsedUrl.Path, queryString)

	encryptedQueryString := Ase256(pathWithQueryString, encryptionKey)

	return fmt.Sprintf("%s://%s/%d/%s", parsedUrl.Scheme, parsedUrl.Host, payment.getProjectId(), encryptedQueryString)
}

func (p *PaymentPage) prepareQueryString(payment Payment) string {
	signature := p.signatureHandler.Sign(payment.GetParams())

	queryArray := []string{}

	for key, value := range payment.GetParams() {
		preparedValue := ""

		switch value := value.(type) {
		case string:
			preparedValue = value
		case int:
			preparedValue = strconv.Itoa(value)
		case bool:
			preparedValue = strconv.FormatBool(value)
		}

		queryArray = append(queryArray, concat(concat(key, "="), url.QueryEscape(preparedValue)))
	}

	queryString := strings.Join(queryArray, "&")
	queryString = concat(queryString, concat("&signature=", url.QueryEscape(signature)))

	return queryString
}

// Constructor for PaymentPage structure
func NewPaymentPage(signatureHandler SignatureHandler) *PaymentPage {
	paymentPage := PaymentPage{signatureHandler}

	return &paymentPage
}
