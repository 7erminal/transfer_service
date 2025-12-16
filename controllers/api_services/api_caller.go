package apiservices

import (
	"bytes"
	"encoding/json"
	"io"
	"transfer_service/api"
	"transfer_service/structs/requests"
	"transfer_service/structs/responses"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func HubtelTransferToAccount(c *beego.Controller, req requests.TransferApiRequest) (responses.TransferApiResponse, error) {
	host, _ := beego.AppConfig.String("hubtelTransferBaseUrl")
	salesId, _ := beego.AppConfig.String("hubtelSalesID")
	authorizationKey, _ := beego.AppConfig.String("authorizationKeySales")

	// serviceId, _ := helpers.GetServiceId(req.Network)
	logs.Info("Sending transfer to account request for ", req.DestinationAccountNumber)
	logs.Info("Callback URL is ", req.CallbackUrl)
	logs.Info("Amount is ", req.Amount)
	logs.Info("Client Reference is ", req.ClientRefernce)

	reqText, _ := json.Marshal(req)

	logs.Info("Request to process transfer to account: ", string(reqText))

	logs.Info("URL:: ", host, "/", salesId)

	request := api.NewRequest(
		host,
		"/"+salesId,
		api.POST)
	request.HeaderField["Authorization"] = "Basic " + authorizationKey

	request.InterfaceParams["Description"] = req.Description
	request.InterfaceParams["Amount"] = req.Amount
	request.InterfaceParams["ClientReference"] = req.ClientRefernce
	request.InterfaceParams["DestinationAccountNumber"] = req.DestinationAccountNumber
	request.InterfaceParams["PrimaryCallbackUrl"] = req.CallbackUrl

	// request.Params = {"UserId": strconv.Itoa(int(userid))}
	client := api.Client{
		Request: request,
		Type_:   "body",
	}
	res, err := client.SendRequest()
	if err != nil {
		logs.Error("client.Error: %v", err)
		c.Data["json"] = err.Error()
	}
	defer res.Body.Close()
	read, err := io.ReadAll(res.Body)
	if err != nil {
		c.Data["json"] = err.Error()
	}

	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, read, "", "  "); err != nil {
		logs.Info("Raw response received is ", string(read))
	} else {
		logs.Info("Raw response received is \n", prettyJSON.String())
	}
	// data := map[string]interface{}{}
	// var dataOri responses.UserOriResponseDTO
	var data responses.TransferApiResponse
	json.Unmarshal(read, &data)
	c.Data["json"] = data

	logs.Info("Resp is ", data)
	// logs.Info("Resp is ", data.User.Branch.Country.DefaultCurrency)

	return data, nil
}
