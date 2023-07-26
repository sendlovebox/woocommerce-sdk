package model

type (
	// Order schema as found in https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#order-properties
	Order struct {
		ID                 int         `json:"id"`
		ParentID           int         `json:"parent_id"`
		Number             string      `json:"number"`
		OrderKey           string      `json:"order_key"`
		CreatedVia         string      `json:"created_via"`
		Version            string      `json:"version"`
		Status             string      `json:"status"`
		Currency           string      `json:"currency"`
		DateCreated        string      `json:"date_created"`
		DateCreatedGmt     string      `json:"date_created_gmt"`
		DateModified       string      `json:"date_modified"`
		DateModifiedGmt    string      `json:"date_modified_gmt"`
		DiscountTotal      string      `json:"discount_total"`
		DiscountTax        string      `json:"discount_tax"`
		ShippingTotal      string      `json:"shipping_total"`
		ShippingTax        string      `json:"shipping_tax"`
		CartTax            string      `json:"cart_tax"`
		Total              string      `json:"total"`
		TotalTax           string      `json:"total_tax"`
		PricesIncludeTax   bool        `json:"prices_include_tax"`
		CustomerID         int         `json:"customer_id"`
		CustomerIPAddress  string      `json:"customer_ip_address"`
		CustomerUserAgent  string      `json:"customer_user_agent"`
		CustomerNote       string      `json:"customer_note"`
		Billing            Address     `json:"billing"`
		Shipping           Address     `json:"shipping"`
		PaymentMethod      string      `json:"payment_method"`
		PaymentMethodTitle string      `json:"payment_method_title"`
		TransactionID      string      `json:"transaction_id"`
		DatePaid           string      `json:"date_paid"`
		DatePaidGmt        string      `json:"date_paid_gmt"`
		DateCompleted      interface{} `json:"date_completed"`
		DateCompletedGmt   interface{} `json:"date_completed_gmt"`
		CartHash           string      `json:"cart_hash"`
		MetaData           []struct {
			ID    int    `json:"id"`
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"meta_data"`
		LineItems     []LineItem     `json:"line_items"`
		TaxLines      []TaxLine      `json:"tax_lines"`
		ShippingLines []ShippingLine `json:"shipping_lines"`
		FeeLines      []interface{}  `json:"fee_lines"`
		CouponLines   []interface{}  `json:"coupon_lines"`
		Refunds       []interface{}  `json:"refunds"`
		Links         Links          `json:"_links"`
	}

	// CreateOrderRequest request payload to create a new order. Reference in https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#create-an-order
	CreateOrderRequest struct {
		CustomerID         int                   `json:"customer_id,omitempty"`
		PaymentMethod      string                `json:"payment_method,omitempty"`
		PaymentMethodTitle string                `json:"payment_method_title,omitempty"`
		TransactionID      string                `json:"transaction_id,omitempty"`
		CustomerNote       string                `json:"customer_note,omitempty"`
		SetPaid            bool                  `json:"set_paid,omitempty"`
		Billing            Address               `json:"billing"`
		Shipping           Address               `json:"shipping"`
		FeeLines           []FeeLineRequest      `json:"fee_lines,omitempty"`
		LineItems          []LineItemRequest     `json:"line_items"`
		ShippingLines      []ShippingLineRequest `json:"shipping_lines"`
	}

	// UpdateOrderRequest request payload to update an existing order. Reference in https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#update-an-order
	UpdateOrderRequest struct {
		Status        string `json:"status,omitempty"`
		CustomerID    int    `json:"customer_id,omitempty"`
		CustomerNote  string `json:"customer_note,omitempty"`
		TransactionID string `json:"transaction_id,omitempty"`
		SetPaid       bool   `json:"set_paid,omitempty"`
	}

	// FeeLineRequest struct
	FeeLineRequest struct {
		Name      string `json:"name,omitempty"`
		TaxClass  string `json:"tax_class,omitempty"`
		TaxStatus string `json:"tax_status,omitempty"`
		Total     string `json:"total,omitempty"`
	}

	// LineItemRequest schema
	LineItemRequest struct {
		ProductID   int `json:"product_id"`
		Quantity    int `json:"quantity"`
		VariationID int `json:"variation_id,omitempty"`
	}

	// ShippingLineRequest schema
	ShippingLineRequest struct {
		MethodID    string `json:"method_id"`
		MethodTitle string `json:"method_title"`
		Total       string `json:"total"`
	}

	// Address schema
	Address struct {
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		Company   string `json:"company,omitempty"`
		Address1  string `json:"address_1,omitempty"`
		Address2  string `json:"address_2,omitempty"`
		City      string `json:"city,omitempty"`
		State     string `json:"state,omitempty"`
		Postcode  string `json:"postcode,omitempty"`
		Country   string `json:"country,omitempty"`
		Email     string `json:"email,omitempty"`
		Phone     string `json:"phone,omitempty"`
	}

	// ShippingLine details and schema for an order
	ShippingLine struct {
		ID          int           `json:"id"`
		MethodTitle string        `json:"method_title"`
		MethodID    string        `json:"method_id"`
		Total       string        `json:"total"`
		TotalTax    string        `json:"total_tax"`
		Taxes       []interface{} `json:"taxes"`
		MetaData    []interface{} `json:"meta_data"`
	}

	// TaxLine details and schema for an order
	TaxLine struct {
		ID               int           `json:"id"`
		RateCode         string        `json:"rate_code"`
		RateID           int           `json:"rate_id"`
		Label            string        `json:"label"`
		Compound         bool          `json:"compound"`
		TaxTotal         string        `json:"tax_total"`
		ShippingTaxTotal string        `json:"shipping_tax_total"`
		MetaData         []interface{} `json:"meta_data"`
	}

	// LineItem details and schema for an order i.e. what was actually ordered!
	LineItem struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		ProductID   int    `json:"product_id"`
		VariationID int    `json:"variation_id"`
		Quantity    int    `json:"quantity"`
		TaxClass    string `json:"tax_class"`
		Subtotal    string `json:"subtotal"`
		SubtotalTax string `json:"subtotal_tax"`
		Total       string `json:"total"`
		TotalTax    string `json:"total_tax"`
		Taxes       []struct {
			ID       int    `json:"id"`
			Total    string `json:"total"`
			Subtotal string `json:"subtotal"`
		} `json:"taxes"`
		MetaData []struct {
			ID    int    `json:"id"`
			Key   string `json:"key"`
			Value string `json:"value"`
		} `json:"meta_data"`
		SKU   string `json:"sku"`
		Price int    `json:"price"`
	}
)

const (
	// ShippingMethodFlatRateID string for id for flat rate shipping method type
	ShippingMethodFlatRateID = "flat_rate"
	// ShippingMethodFlatRateTitle string for title for flat rate shipping method type
	ShippingMethodFlatRateTitle = "Flat Rate"

	// PaymentMethodBankTransfer string representation of the bank transfer
	PaymentMethodBankTransfer = "bacs"
	// PaymentMethodTitleBankTransfer  string representation of the bank transfer title
	PaymentMethodTitleBankTransfer = "Direct Bank Transfer"
)
