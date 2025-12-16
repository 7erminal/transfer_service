package controllers

import (
	"encoding/json"
	"strconv"
	"time"

	apiservices "transfer_service/controllers/api_services"
	"transfer_service/models"
	"transfer_service/structs/requests"
	"transfer_service/structs/responses"

	"github.com/beego/beego/v2/core/logs"

	beego "github.com/beego/beego/v2/server/web"
)

// TransferController operations for Transfer
type TransferController struct {
	beego.Controller
}

// URLMapping ...
func (c *TransferController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("TransferCommission", c.TransferCommission)
}

// Post ...
// @Title Create
// @Description create Transfer
// @Param	body		body 	requests.TransferRequest	true		"body for Transfer content"
// @Success 201 {object} models.Transfer
// @Failure 403 body is empty
// @router / [post]
func (c *TransferController) Post() {
	var v requests.TransferRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	responseCode := 401
	responseMessage := "Error processing request"
	result := models.Trx_transactions{}

	if status, err := models.GetStatusByName(v.Status); err == nil {
		// Generate a unique transaction ID (you can customize this as needed)
		transactionId := "TXN" + strconv.FormatInt(time.Now().Unix(), 10) + "." + v.RequestId

		trx_transaction := models.Trx_transactions{
			TransactionId:          transactionId,
			Amount:                 v.Amount,
			TotalDebitAmount:       v.TotalDebitAmount,
			Charge:                 v.Charge,
			Commission:             v.Commission,
			SenderAccountNumber:    v.SenderAccountNumber,
			RecipientAccountNumber: v.RecipientAccountNumber,
			TransferCode:           v.TransferCode,
			ResponseCode:           "",
			ResponseMessage:        "",
			CreatedBy:              1,
			ModifiedBy:             1,
			Active:                 1,
			Status:                 status,
		}

		if _, err := models.AddTrx_transactions(&trx_transaction); err == nil {
			trx_details_id := "TRXDTL" + strconv.FormatInt(time.Now().Unix(), 10) + "." + v.RequestId
			trx_transactionDetails := models.Trx_transaction_details{
				Trx_transactionDetailsId: trx_details_id,
				TransactionDescription:   v.Description,
				TransactionId:            &trx_transaction,
				Amount:                   v.Amount,
				Charge:                   v.Charge,
				Commission:               v.Commission,
				Sender:                   v.SenderAccountNumber,
				Recipient:                v.RecipientAccountNumber,
				Status:                   status,
				ResponseCode:             "",
				RecipientName:            v.RecipientName,
				DateCreated:              time.Now(),
				DateModified:             time.Now(),
				CreatedBy:                1,
				ModifiedBy:               1,
				Active:                   1,
			}
			responseCode = 200
			responseMessage = "Transfer created successfully"
			result = trx_transaction

			if _, err := models.AddTrx_transaction_details(&trx_transactionDetails); err == nil {
				// Successfully added transaction details
				logs.Info("Transaction details added successfully")
				responseMessage = "Transaction processed successfully"
			} else {
				// Handle error adding transaction details
				responseMessage = "Transfer partially processed: " + err.Error()
				logs.Error("Error adding transaction details: %v", err)
			}

			// Send commission to commission wallet
		} else {
			responseCode = 500
			responseMessage = "Error creating transfer: " + err.Error()
		}
	}

	resp := responses.TransferResponseDTO{
		StatusCode: responseCode,
		Result:     &result,
		StatusDesc: responseMessage,
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

// TransferCommission ...
// @Title Transfer Commission
// @Description Transfer Commission
// @Param	body		body 	requests.TransferCommissionRequest	true		"body for Transfer content"
// @Success 201 {object} models.Transfer
// @Failure 403 body is empty
// @router /transfer-commission/ [post]
func (c *TransferController) TransferCommission() {
	var v requests.TransferCommissionRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	responseCode := 401
	responseMessage := "Error processing request"
	result := models.Trx_transactions{}

	if status, err := models.GetStatusByName(v.Status); err == nil {
		// Generate a unique transaction ID (you can customize this as needed)
		if trx_transaction, err := models.GetTrx_transactionsById(v.TransactionId); err == nil {
			trx_transaction.Commission = v.Commission

			if err := models.UpdateTrx_transactionsById(trx_transaction); err == nil {
				trx_details_id := "TRXDTL" + strconv.FormatInt(time.Now().Unix(), 10) + "." + v.RequestId
				trx_transactionDetails := models.Trx_transaction_details{
					Trx_transactionDetailsId: trx_details_id,
					TransactionDescription:   "Commission Transfer for Transaction ID " + v.TransactionId,
					TransactionId:            trx_transaction,
					Amount:                   v.Amount,
					Charge:                   v.Charge,
					Commission:               v.Commission,
					Sender:                   v.SenderAccountNumber,
					Recipient:                v.RecipientAccountNumber,
					Status:                   status,
					ResponseCode:             "",
					RecipientName:            "Commission Wallet",
					DateCreated:              time.Now(),
					DateModified:             time.Now(),
					CreatedBy:                1,
					ModifiedBy:               1,
					Active:                   1,
				}

				if _, err := models.AddTrx_transaction_details(&trx_transactionDetails); err == nil {
					// Successfully added transaction details
					logs.Info("Commission transaction details added successfully")

					if callbackUrl, err := models.GetApplication_propertyByCode("TRANSFER_TO_ACCOUNT_CALLBACK_URL"); err == nil {
						req := requests.TransferApiRequest{
							ClientRefernce:           trx_transaction.TransactionId,
							Amount:                   v.Amount,
							Description:              v.Description,
							DestinationAccountNumber: v.RecipientAccountNumber,
							CallbackUrl:              callbackUrl.PropertyValue,
						}
						if resp, err := apiservices.HubtelTransferToAccount(&c.Controller, req); err == nil {
							logs.Info("Commission transferred to commission wallet successfully: %v", resp)

							responseCode = 200
							responseMessage = "Commission transferred successfully"
							result = *trx_transaction
						} else {
							// Handle error transferring commission
							logs.Error("Error transferring commission to commission wallet: %v", err)
						}

					} else {
						// Handle error transferring commission
						responseCode = 500
						responseMessage = "Error transferring commission: " + err.Error()
						logs.Error("Error transferring commission to commission wallet: %v", err)
					}
				} else {
					// Handle error adding transaction details
					responseCode = 500
					responseMessage = "Error transferring commission: " + err.Error()
					logs.Error("Error adding commission transaction details: %v", err)
				}
			} else {
				responseCode = 500
				responseMessage = "Error updating transfer: " + err.Error()
			}
		} else {
			logs.Error("Transaction not found: %v", err)
			responseCode = 404
			responseMessage = "Transaction not found: " + err.Error()
		}
	}

	resp := responses.TransferResponseDTO{
		StatusCode: responseCode,
		Result:     &result,
		StatusDesc: responseMessage,
	}

	c.Data["json"] = resp
	c.ServeJSON()
}

// GetOne ...
// @Title GetOne
// @Description get Transfer by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Transfer
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TransferController) GetOne() {

}

// GetAll ...
// @Title GetAll
// @Description get Transfer
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Transfer
// @Failure 403
// @router / [get]
func (c *TransferController) GetAll() {

}

// Put ...
// @Title Put
// @Description update the Transfer
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Transfer	true		"body for Transfer content"
// @Success 200 {object} models.Transfer
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TransferController) Put() {

}

// Delete ...
// @Title Delete
// @Description delete the Transfer
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TransferController) Delete() {

}
