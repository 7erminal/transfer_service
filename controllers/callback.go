package controllers

import (
	"encoding/json"

	"transfer_service/models"
	"transfer_service/structs/requests"
	"transfer_service/structs/responses"

	"github.com/beego/beego/v2/core/logs"

	beego "github.com/beego/beego/v2/server/web"
)

// TransferController operations for Transfer
type CallbackController struct {
	beego.Controller
}

// URLMapping ...
func (c *CallbackController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// Post ...
// @Title Create
// @Description create Transfer
// @Param	body		body 	requests.TransferCallbackRequest	true		"body for Transfer content"
// @Success 201 {object} models.Transfer
// @Failure 403 body is empty
// @router / [post]
func (c *CallbackController) Post() {
	var v requests.TransferCallbackRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	responseCode := 401
	responseMessage := "Error processing request"
	result := models.Trx_transactions{}

	if status, err := models.GetStatusByName(v.Status); err == nil {
		// Get Transaction by Transaction ID
		if trxTransaction, err := models.GetTrx_transactionsById(v.TransactionId); err == nil {
			trxTransaction.Status = status
			if status, err := models.GetStatusByName(v.Status); err == nil {
				trxTransaction.ResponseCode = v.ResponseCode
				trxTransaction.ResponseMessage = v.ResponseMessage
				trxTransaction.Status = status

				if err := models.UpdateTrx_transactionsById(trxTransaction); err == nil {
					logs.Info("Transaction Updated")
					responseCode = 200
					responseMessage = "Transaction updated successfully"
					result = *trxTransaction
				} else {
					logs.Error("Failed to update transaction ", err)
					responseCode = 500
					responseMessage = "Error updating transaction: " + err.Error()
				}
			} else {
				responseCode = 404
				responseMessage = "Transaction not found: " + err.Error()
			}
		} else {
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
