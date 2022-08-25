
POST https://payment-service:60000/charge


{
	"amount": {
		"currency_code": "USD",
		"units": 100,
		"nanos": 50
	},
	"credit_card": {
		"credit_card_number": "4432-8015-6152-0454",
		"credit_card_cvv": 672,
		"credit_card_expiration_year": 2024,
		"credit_card_expiration_month": 2
	}
}
