package responses

import "transfer_service/models"

type TransferResponseApiData struct {
	ClientRefernce string  `json:"clientReference"`
	Amount         float64 `json:"amount"`
	Charges        float64 `json:"charges"`
	Description    string  `json:"description"`
	RecipientName  string  `json:"recipientName"`
	Meta           string  `json:"meta"`
}

type TransferApiResponse struct {
	ResponseCode string                  `json:"responseCode"`
	Message      string                  `json:"message"`
	Data         TransferResponseApiData `json:"data"`
}

type TransferResponseData struct {
	ClientRefernce string  `json:"clientReference"`
	Amount         float64 `json:"amount"`
	Charges        float64 `json:"charges"`
	Description    string  `json:"description"`
	RecipientName  string  `json:"recipientName"`
	Meta           string  `json:"meta"`
}

type TransferResponseDTO struct {
	StatusCode int                      `json:"success"`
	StatusDesc string                   `json:"statusDesc"`
	Result     *models.Trx_transactions `json:"result"`
}
