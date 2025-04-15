package value

// Model is the description of the method for providing "Value for Value" payments
type Model struct {
	// Type is the payment type
	//
	// expected to be either value.PaymentLightning, value.PaymentHive, or value.PaymentWebMonetization
	Type string `json:"type"`
	// Method is the method for sending payment
	Method string `json:"method"`
	// Suggested is the suggested per second of playback to send. Unit is specific to the type; value may be empty.
	Suggested string `json:"suggested"`
}
