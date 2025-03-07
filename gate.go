package paymentpage

// Structure for communicate with our
type Gate struct {
	// Instance for check signature
	signatureHandler SignatureHandler

	// Instance for build payment URL
	paymentPage PaymentPage
}

// Method build payment URL
func (g *Gate) GetPaymentPageUrl(baseUrl string, payment Payment) string {
	return g.paymentPage.GetUrl(baseUrl, payment)
}

// Method builds encrypted payment URL
func (g *Gate) GetEncryptedPaymentPageUrl(baseUrl string, payment Payment, encryptionKey string) string {
	return g.paymentPage.GetEncryptedUrl(baseUrl, payment, encryptionKey)
}

// Method for handling callback
func (g *Gate) HandleCallback(callbackData string) (*Callback, error) {
	callback, callbackError := NewCallback(g.signatureHandler, callbackData)

	return callback, callbackError
}

// Constructor for Gate structure
func NewGate(secret string) *Gate {
	signatureHandler := NewSignatureHandler(secret)
	paymentPage := NewPaymentPage(*signatureHandler)
	gate := Gate{*signatureHandler, *paymentPage}

	return &gate
}
