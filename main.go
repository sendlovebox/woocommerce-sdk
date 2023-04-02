package main

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"

	apiCalls "github.com/sendlovebox/woocommerce-sdk/api"
	"github.com/sendlovebox/woocommerce-sdk/model"
)

func main() {
	fmt.Println("Running Woocommerce SDK")
	_ = os.Setenv("TZ", "Africa/Lagos")
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	client := resty.New()
	logger.Info().Msg("app starting")
	defer logger.Info().Msg("stopped")
	api := apiCalls.New(&logger, client, model.BaseURL)
	api.RunInSandboxMode() // to ensure it is running in sandbox mode

	// we need to authenticate the application
	_ = api.Auth("", "")

	//products, _ := api.SearchProducts(context.Background(), model.SearchProductsRequest{
	//	Page:    1,
	//	PerPage: 10,
	//Order:          "",
	//OrderBy:        "",
	//Context:        "",
	//Search:         "",
	//After:          "",
	//Before:         "",
	//ModifiedAfter:  "",
	//ModifiedBefore: "",
	//DateAreGmt:     false,
	//Exclude: []string{"12", "34"},
	//Include:        nil,
	//Parent:         nil,
	//ParentExclude:  nil,
	//Slug:           "",
	//Status:         "",
	//Type:           "",
	//SKU:            "",
	//Featured:       "",
	//Category: "234",
	//Tag:            "",
	//ShippingClass:  "",
	//Attribute:      "",
	//AttributeTerm: "attribute",
	//TaxClass:    "",
	//OnSale:      "",
	//MinPrice:    "",
	//MaxPrice:    "",
	//StockStatus: "",
	//})
	//fmt.Println("this is the length of the products ---> ", len(products))

	//product, _ := api.GetProductByID(context.Background(), "13")
	//fmt.Println("product is --->", product)
}
