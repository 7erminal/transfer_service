package requests

type TransferApiRequest struct {
	ClientRefernce           string
	Amount                   float64
	Description              string
	DestinationAccountNumber string
	CallbackUrl              string
}

type TransferRequest struct {
	RequestId              string
	Amount                 float64
	Charge                 float64
	Commission             float64
	TotalDebitAmount       float64
	SenderAccountNumber    string
	RecipientAccountNumber string
	TransferCode           string
	Description            string
	RecipientName          string
	Status                 string
}

type TransferCommissionRequest struct {
	TransactionId          string
	RequestId              string
	Amount                 float64
	Charge                 float64
	Commission             float64
	TotalDebitAmount       float64
	SenderAccountNumber    string
	RecipientAccountNumber string
	TransferCode           string
	Description            string
	RecipientName          string
	Status                 string
}

type TransferCallbackRequest struct {
	TransactionId   string
	ResponseCode    string
	ResponseMessage string
	Status          string
	Charge          string
	RecipientName   string
}
