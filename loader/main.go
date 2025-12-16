package main

import (
	"fmt"
	"io"
	"os"
	"transfer_service/models"

	"ariga.io/atlas-provider-beego/beegoschema"
	"github.com/beego/beego/v2/client/orm"
)

func main() {
	stmts, err := beegoschema.New("mysql").Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to load beego schema: %v\n", err)
		os.Exit(1)
	}
	io.WriteString(os.Stdout, stmts)
}

// If your models are already registered in an init() function elsewhere, you can simply use
// a blank import to ensure that the init() function is called. Otherwise, you can register
// your models here.
func init() {
	orm.RegisterModel(new(models.Trx_transactions))
}
