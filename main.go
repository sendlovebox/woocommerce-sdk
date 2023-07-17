// Package main is the entry point of the application
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

	//products, _ := api.ListAllProducts(context.Background(), model.SearchProductsRequest{
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

	//product, _ := api.RetrieveAProduct(context.Background(), "13")
	//fmt.Println("product is --->", product)

	//tags, _ := api.ListAllProductTags(context.Background(), model.SearchTagsRequest{
	//	PaginationRequest: model.PaginationRequest{
	//		Page:    2,
	//		PerPage: 2,
	//		Offset:  0,
	//		Order:   "desc",
	//		OrderBy: "count",
	//	},
	//	Context:   "",
	//	Search:    "",
	//	Exclude:   nil,
	//	Include:   nil,
	//	HideEmpty: false,
	//	Product:   0,
	//	Slug:      "",
	//})
	//fmt.Println("product tags are --->", tags)

	//categories, _ := api.ListAllCategories(context.Background(), model.SearchCategoriesRequest{
	//	Page:    1,
	//	PerPage: 10,
	//	//Order:     "",
	//	//OrderBy:   "",
	//	//Context:   "",
	//	//Search:    "",
	//	//Exclude:   nil,
	//	//Include:   nil,
	//	//Parent:    0,
	//	//Product:   0,
	//	//HideEmpty: false,
	//	//Slug:      "",
	//})
	//fmt.Println("categories --->", categories)

	//category, _ := api.RetrieveACategory(context.Background(), "18")
	//fmt.Println("category is --->", category)

	//customer, _ := api.CreateACustomer(context.Background(), model.CreateCustomerRequest{
	//	Email:     "abcde@ninja.com",
	//	FirstName: "ABCDE",
	//	LastName:  "NINJA",
	//	Username:  "abcde-ninja",
	//	Billing: model.Address{
	//		//FirstName: "",
	//		//LastName:  "",
	//		//Company:   "",
	//		//Address1:  "",
	//		//Address2:  "",
	//		//City:      "",
	//		//State:     "",
	//		//Postcode:  "",
	//		//Country:   "",
	//		Email: "abcde@ninja.com",
	//		//Phone:     "",
	//	},
	//	Shipping: model.Address{
	//		//FirstName: "",
	//		//LastName:  "",
	//		//Company:   "",
	//		//Address1:  "",
	//		//Address2:  "",
	//		//City:      "",
	//		//State:     "",
	//		//Postcode:  "",
	//		//Country:   "",
	//		Email: "abcde@ninja.com",
	//		//Phone:     "",
	//	},
	//})
	//fmt.Println("customer is --->", customer)
}
