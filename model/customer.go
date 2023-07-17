package model

type (
	// Customer schema and properties as seen in https://woocommerce.github.io/woocommerce-rest-api-docs/?javascript#customer-properties
	Customer struct {
		ID               int           `json:"id"`
		DateCreated      string        `json:"date_created"`
		DateCreatedGmt   string        `json:"date_created_gmt"`
		DateModified     string        `json:"date_modified"`
		DateModifiedGmt  string        `json:"date_modified_gmt"`
		Email            string        `json:"email"`
		FirstName        string        `json:"first_name"`
		LastName         string        `json:"last_name"`
		Role             string        `json:"role"`
		Username         string        `json:"username"`
		Billing          Address       `json:"billing"`
		Shipping         Address       `json:"shipping"`
		IsPayingCustomer bool          `json:"is_paying_customer"`
		AvatarURL        string        `json:"avatar_url"`
		MetaData         []interface{} `json:"meta_data"`
		Links            Links         `json:"_links"`
	}

	// CreateCustomerRequest payload to create a new customer. Reference - https://woocommerce.github.io/woocommerce-rest-api-docs/?javascript#create-a-customer
	CreateCustomerRequest struct {
		Email     string  `json:"email"`
		FirstName string  `json:"first_name"`
		LastName  string  `json:"last_name"`
		Username  string  `json:"username"`
		Billing   Address `json:"billing,omitempty"`
		Shipping  Address `json:"shipping,omitempty"`
	}
)
