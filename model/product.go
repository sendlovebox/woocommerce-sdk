package model

type (
	// SearchProductsRequest schema to search or filter products
	SearchProductsRequest struct {
		PaginationRequest

		Context        string      `url:"context,omitempty" json:"context"`
		Search         string      `url:"search,omitempty" json:"search"`
		After          string      `url:"after,omitempty" json:"after"`
		Before         string      `url:"before,omitempty" json:"before"`
		ModifiedAfter  string      `url:"modified_after,omitempty" json:"modified_after"`
		ModifiedBefore string      `url:"modified_before,omitempty" json:"modified_before"`
		DateAreGmt     bool        `url:"date_are_gmt,omitempty" json:"date_are_gmt"`
		Featured       bool        `url:"featured,omitempty" json:"featured"`
		OnSale         bool        `url:"on_sale,omitempty" json:"on_sale"`
		Exclude        []string    `url:"exclude,omitempty" json:"exclude"`
		Include        []string    `url:"include,omitempty" json:"include"`
		Parent         []string    `url:"parent,omitempty" json:"parent"`
		ParentExclude  []string    `url:"parent_exclude,omitempty" json:"parent_exclude"`
		Status         Status      `url:"status,omitempty" json:"status"`
		Type           ProductType `url:"type,omitempty" json:"type"`
		SKU            string      `url:"sku,omitempty" json:"sku"`
		Category       string      `url:"category,omitempty" json:"category"`
		Tag            string      `url:"tag,omitempty" json:"tag"`
		ShippingClass  string      `url:"shipping_class,omitempty" json:"shipping_class"`
		Attribute      string      `url:"attribute,omitempty" json:"attribute"`
		AttributeTerm  string      `url:"attribute_term,omitempty" json:"attribute_term"`
		TaxClass       string      `url:"tax_class,omitempty" json:"tax_class"`
		MinPrice       string      `url:"min_price,omitempty" json:"min_price"`
		MaxPrice       string      `url:"max_price,omitempty" json:"max_price"`
		StockStatus    string      `url:"stock_status,omitempty" json:"stock_status"`
		Slug           string      `url:"slug,omitempty" json:"slug"`
	}

	// Product schema
	Product struct {
		ID                int           `mapstructure:"id" json:"id"`
		Name              string        `mapstructure:"name" json:"name"`
		Slug              string        `mapstructure:"slug" json:"slug"`
		Permalink         string        `mapstructure:"permalink" json:"permalink"`
		DateCreated       string        `mapstructure:"date_created" json:"date_created"`
		DateCreatedGmt    string        `mapstructure:"date_created_gmt" json:"date_created_gmt"`
		DateModified      string        `mapstructure:"date_modified" json:"date_modified"`
		DateModifiedGmt   string        `mapstructure:"date_modified_gmt" json:"date_modified_gmt"`
		Type              string        `mapstructure:"type" json:"type"`
		Status            string        `mapstructure:"status" json:"status"`
		Featured          bool          `mapstructure:"featured" json:"featured"`
		CatalogVisibility string        `mapstructure:"catalog_visibility" json:"catalog_visibility"`
		Description       string        `mapstructure:"description" json:"description"`
		ShortDescription  string        `mapstructure:"short_description" json:"short_description"`
		SKU               string        `mapstructure:"sku" json:"sku"`
		Price             string        `mapstructure:"price" json:"price"`
		RegularPrice      string        `mapstructure:"regular_price" json:"regular_price"`
		SalePrice         string        `mapstructure:"sale_price" json:"sale_price"`
		DateOnSaleFrom    string        `mapstructure:"date_on_sale_from" json:"date_on_sale_from"`
		DateOnSaleFromGmt string        `mapstructure:"date_on_sale_from_gmt" json:"date_on_sale_from_gmt"`
		DateOnSaleTo      string        `mapstructure:"date_on_sale_to" json:"date_on_sale_to"`
		DateOnSaleToGmt   string        `mapstructure:"date_on_sale_to_gmt" json:"date_on_sale_to_gmt"`
		PriceHTML         string        `mapstructure:"price_html" json:"price_html"`
		OnSale            bool          `mapstructure:"on_sale" json:"on_sale"`
		Purchasable       bool          `mapstructure:"purchasable" json:"purchasable"`
		TotalSales        int           `mapstructure:"total_sales" json:"total_sales"`
		Virtual           bool          `mapstructure:"virtual" json:"virtual"`
		Downloadable      bool          `mapstructure:"downloadable" json:"downloadable"`
		Downloads         []interface{} `mapstructure:"downloads" json:"downloads"`
		DownloadLimit     int           `mapstructure:"download_limit" json:"download_limit"`
		DownloadExpiry    int           `mapstructure:"download_expiry" json:"download_expiry"`
		ExternalURL       string        `mapstructure:"external_url" json:"external_url"`
		ButtonText        string        `mapstructure:"button_text" json:"button_text"`
		TaxStatus         string        `mapstructure:"tax_status" json:"tax_status"`
		TaxClass          string        `mapstructure:"tax_class" json:"tax_class"`
		ManageStock       bool          `mapstructure:"manage_stock" json:"manage_stock"`
		StockQuantity     interface{}   `mapstructure:"stock_quantity" json:"stock_quantity"`
		StockStatus       string        `mapstructure:"stock_status" json:"stock_status"`
		BackOrders        string        `mapstructure:"backorders" json:"backorders"`
		BackordersAllowed bool          `mapstructure:"backorders_allowed" json:"backorders_allowed"`
		BackOrdered       bool          `mapstructure:"backordered" json:"backordered"`
		SoldIndividually  bool          `mapstructure:"sold_individually" json:"sold_individually"`
		Weight            string        `mapstructure:"weight" json:"weight"`
		Dimensions        struct {
			Length string `mapstructure:"length" json:"length"`
			Width  string `mapstructure:"width" json:"width"`
			Height string `mapstructure:"height" json:"height"`
		} `mapstructure:"dimensions" json:"dimensions"`
		ShippingRequired bool   `mapstructure:"shipping_required" json:"shipping_required"`
		ShippingTaxable  bool   `mapstructure:"shipping_taxable" json:"shipping_taxable"`
		ShippingClass    string `mapstructure:"shipping_class" json:"shipping_class"`
		ShippingClassID  int    `mapstructure:"shipping_class_id" json:"shipping_class_id"`
		ReviewsAllowed   bool   `mapstructure:"reviews_allowed" json:"reviews_allowed"`
		AverageRating    string `mapstructure:"average_rating" json:"average_rating"`
		RatingCount      int    `mapstructure:"rating_count" json:"rating_count"`
		RelatedIDs       []int  `mapstructure:"related_ids" json:"related_ids"`
		UpsellIDs        []int  `mapstructure:"upsell_ids" json:"upsell_ids"`
		CrossSellIDs     []int  `mapstructure:"cross_sell_ids" json:"cross_sell_ids"`
		ParentID         int    `mapstructure:"parent_id" json:"parent_id"`
		PurchaseNote     string `mapstructure:"purchase_note" json:"purchase_note"`
		Categories       []struct {
			ID   int    `mapstructure:"id" json:"id"`
			Name string `mapstructure:"name" json:"name"`
			Slug string `mapstructure:"slug" json:"slug"`
		} `mapstructure:"categories" json:"categories"`
		Tags []struct {
			ID   int    `mapstructure:"id" json:"id"`
			Name string `mapstructure:"name" json:"name"`
			Slug string `mapstructure:"slug" json:"slug"`
		} `mapstructure:"tags" json:"tags"`
		Images            []Image            `mapstructure:"images" json:"images"`
		Attributes        []ProductAttribute `mapstructure:"attributes" json:"attributes"`
		DefaultAttributes []interface{}      `mapstructure:"default_attributes" json:"default_attributes"`
		Variations        []int              `mapstructure:"variations" json:"variations"`
		GroupedProducts   []interface{}      `mapstructure:"grouped_products" json:"grouped_products"`
		MenuOrder         int                `mapstructure:"menu_order" json:"menu_order"`
		MetaData          []Metadata         `mapstructure:"meta_data" json:"meta_data"`
		Links             struct {
			Self []struct {
				Href string `mapstructure:"href" json:"href"`
			} `mapstructure:"self" json:"self"`
			Collection []struct {
				Href string `mapstructure:"href" json:"href"`
			} `mapstructure:"collection" json:"collection"`
		} `mapstructure:"_links" json:"_links"`
	}

	// Metadata is the schema for the metadata
	Metadata struct {
		ID    int         `mapstructure:"id" json:"id"`
		Key   string      `mapstructure:"key" json:"key"`
		Value interface{} `mapstructure:"value" json:"value"`
	}

	// ProductVariation schema
	ProductVariation struct {
		ID                int           `json:"id"`
		DateCreated       string        `json:"date_created"`
		DateCreatedGmt    string        `json:"date_created_gmt"`
		DateModified      string        `json:"date_modified"`
		DateModifiedGmt   string        `json:"date_modified_gmt"`
		Description       string        `json:"description"`
		Permalink         string        `json:"permalink"`
		SKU               string        `json:"sku"`
		Price             string        `json:"price"`
		RegularPrice      string        `json:"regular_price"`
		SalePrice         string        `json:"sale_price"`
		DateOnSaleFrom    string        `json:"date_on_sale_from"`
		DateOnSaleFromGmt string        `json:"date_on_sale_from_gmt"`
		DateOnSaleTo      string        `json:"date_on_sale_to"`
		DateOnSaleToGmt   string        `json:"date_on_sale_to_gmt"`
		OnSale            bool          `json:"on_sale"`
		Status            string        `json:"status"`
		Purchasable       bool          `json:"purchasable"`
		Virtual           bool          `json:"virtual"`
		Downloadable      bool          `json:"downloadable"`
		Downloads         []interface{} `json:"downloads"`
		DownloadLimit     int           `json:"download_limit"`
		DownloadExpiry    int           `json:"download_expiry"`
		TaxStatus         string        `json:"tax_status"`
		TaxClass          string        `json:"tax_class"`
		ManageStock       bool          `json:"manage_stock"`
		StockQuantity     interface{}   `json:"stock_quantity"`
		StockStatus       string        `json:"stock_status"`
		Backorders        string        `json:"backorders"`
		BackordersAllowed bool          `json:"backorders_allowed"`
		BackOrdered       bool          `json:"backordered"`
		Weight            string        `json:"weight"`
		Dimensions        struct {
			Length string `json:"length"`
			Width  string `json:"width"`
			Height string `json:"height"`
		} `json:"dimensions"`
		ShippingClass   string               `json:"shipping_class"`
		ShippingClassID int                  `json:"shipping_class_id"`
		Image           Image                `json:"image"`
		Attributes      []VariationAttribute `json:"attributes"`
		MenuOrder       int                  `json:"menu_order"`
		MetaData        []interface{}        `json:"meta_data"`
		Links           struct {
			Self []struct {
				Href string `json:"href"`
			} `json:"self"`
			Collection []struct {
				Href string `json:"href"`
			} `json:"collection"`
			Up []struct {
				Href string `json:"href"`
			} `json:"up"`
		} `json:"_links"`
	}

	// Image schema
	Image struct {
		ID              int    `mapstructure:"id" json:"id"`
		DateCreated     string `mapstructure:"date_created" json:"date_created"`
		DateCreatedGmt  string `mapstructure:"date_created_gmt" json:"date_created_gmt"`
		DateModified    string `mapstructure:"date_modified" json:"date_modified"`
		DateModifiedGmt string `mapstructure:"date_modified_gmt" json:"date_modified_gmt"`
		Src             string `mapstructure:"src" json:"src"`
		Name            string `mapstructure:"name" json:"name"`
		Alt             string `mapstructure:"alt" json:"alt"`
	}

	// SearchTagsRequest schema for request to search for tags
	SearchTagsRequest struct {
		PaginationRequest
		Context   string   `url:"context,omitempty" json:"context"`
		Search    string   `url:"search,omitempty" json:"search"`
		Exclude   []string `url:"exclude,omitempty" json:"exclude"`
		Include   []string `url:"include,omitempty" json:"include"`
		HideEmpty bool     `url:"hide_empty,omitempty" json:"hide_empty"`
		Product   int      `url:"product,omitempty" json:"product"`
		Slug      string   `url:"slug,omitempty" json:"slug"`
	}

	// ProductTag schema for product tags
	ProductTag struct {
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Slug        string `json:"slug"`
		Description string `json:"description"`
		Count       int    `json:"count"`
		Links       struct {
			Self []struct {
				Href string `json:"href"`
			} `json:"self"`
			Collection []struct {
				Href string `json:"href"`
			} `json:"collection"`
		} `json:"_links"`
	}

	// PaginationRequest schema for all requests
	PaginationRequest struct {
		Page    int       `url:"page,omitempty" json:"page,omitempty"`
		PerPage int       `url:"per_page,omitempty" json:"per_page,omitempty"`
		Offset  int       `url:"offset,omitempty" json:"offset,omitempty"`
		Order   OrderType `url:"order,omitempty" json:"order,omitempty"`
		OrderBy OrderBy   `url:"orderby,omitempty" json:"orderby,omitempty"`
	}

	// ProductAttribute schema
	ProductAttribute struct {
		ID        int      `mapstructure:"id" json:"id"`
		Name      string   `mapstructure:"name" json:"name"`
		Position  int      `mapstructure:"position" json:"position"`
		Visible   bool     `mapstructure:"visible" json:"visible"`
		Variation bool     `mapstructure:"variation" json:"variation"`
		Options   []string `mapstructure:"options" json:"options"`
	}

	// VariationAttribute schema
	VariationAttribute struct {
		ID     int    `mapstructure:"id" json:"id"`
		Name   string `mapstructure:"name" json:"name"`
		Option string `mapstructure:"option" json:"option"`
	}

	// SearchProductVariationsRequest schema to search for product variations
	SearchProductVariationsRequest struct {
		PaginationRequest

		Context       string   `url:"context,omitempty" json:"context"`
		Search        string   `url:"search,omitempty" json:"search"`
		After         string   `url:"after,omitempty" json:"after"`
		Before        string   `url:"before,omitempty" json:"before"`
		OnSale        bool     `url:"on_sale,omitempty" json:"on_sale"`
		Exclude       []string `url:"exclude,omitempty" json:"exclude"`
		Include       []string `url:"include,omitempty" json:"include"`
		Parent        []string `url:"parent,omitempty" json:"parent"`
		ParentExclude []string `url:"parent_exclude,omitempty" json:"parent_exclude"`
		Status        Status   `url:"status,omitempty" json:"status"`
		SKU           string   `url:"sku,omitempty" json:"sku"`
		TaxClass      string   `url:"tax_class,omitempty" json:"tax_class"`
		MinPrice      string   `url:"min_price,omitempty" json:"min_price"`
		MaxPrice      string   `url:"max_price,omitempty" json:"max_price"`
		StockStatus   string   `url:"stock_status,omitempty" json:"stock_status"`
		Slug          string   `url:"slug,omitempty" json:"slug"`
	}

	// ReviewProductRequest schema to create a product review - https://woocommerce.github.io/woocommerce-rest-api-docs/?javascript#create-a-product-review
	ReviewProductRequest struct {
		ProductID     int    `json:"product_id,omitempty"`
		Review        string `json:"review,omitempty"`
		Reviewer      string `json:"reviewer,omitempty"`
		ReviewerEmail string `json:"reviewer_email,omitempty"`
		Rating        int    `json:"rating,omitempty"`
	}

	// ProductReview schema as seen in https://woocommerce.github.io/woocommerce-rest-api-docs/?javascript#product-review-properties
	ProductReview struct {
		ID                 int    `json:"id"`
		DateCreated        string `json:"date_created"`
		DateCreatedGmt     string `json:"date_created_gmt"`
		ProductID          int    `json:"product_id"`
		Status             string `json:"status"`
		Reviewer           string `json:"reviewer"`
		ReviewerEmail      string `json:"reviewer_email"`
		Review             string `json:"review"`
		Rating             int    `json:"rating"`
		Verified           bool   `json:"verified"`
		ReviewerAvatarUrls struct {
			Field1 string `json:"24"`
			Field2 string `json:"48"`
			Field3 string `json:"96"`
		} `json:"reviewer_avatar_urls"`
		Links interface{} `json:"_links"`
	}

	// SearchProductReviewsRequest schema to search for product reviews - https://woocommerce.github.io/woocommerce-rest-api-docs/?javascript#list-all-product-reviews
	SearchProductReviewsRequest struct {
		PaginationRequest

		Context         string   `url:"context,omitempty" json:"context"`
		Search          string   `url:"search,omitempty" json:"search"`
		After           string   `url:"after,omitempty" json:"after"`
		Before          string   `url:"before,omitempty" json:"before"`
		Exclude         []string `url:"exclude,omitempty" json:"exclude"`
		Include         []string `url:"include,omitempty" json:"include"`
		Status          Status   `url:"status,omitempty" json:"status"`
		Reviewer        []string `url:"reviewer,omitempty" json:"reviewer"`
		ReviewerExclude []string `url:"reviewer_exclude,omitempty" json:"reviewer_exclude"`
		ReviewerEmail   []string `url:"reviewer_email,omitempty" json:"reviewer_email"`
		Product         []string `url:"product,omitempty" json:"product"`
	}

	// ProductCreateRequest schema to create a product - https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#create-a-product
	ProductCreateRequest struct {
		Name             string `json:"name"`
		Type             string `json:"type,omitempty"`
		RegularPrice     string `json:"regular_price,omitempty"`
		Description      string `json:"description,omitempty"`
		ShortDescription string `json:"short_description,omitempty"`
		Categories       []struct {
			ID int `json:"id"`
		} `json:"categories,omitempty"`
		Images []struct {
			ID  int    `json:"id,omitempty"`
			Src string `json:"src,omitempty"`
		} `json:"images,omitempty"`
	}

	// ProductUpdateRequest schema to update a product (removed all the readonly fields) https://woocommerce.github.io/woocommerce-rest-api-docs/?shell#product-properties
	ProductUpdateRequest struct {
		Name              string `json:"name,omitempty"`
		Slug              string `json:"slug,omitempty"`
		Type              string `json:"type,omitempty"`
		Status            string `json:"status,omitempty"`
		Featured          bool   `json:"featured,omitempty"`
		CatalogVisibility string `json:"catalog_visibility,omitempty"`
		Description       string `json:"description,omitempty"`
		ShortDescription  string `json:"short_description,omitempty"`
		SKU               string `json:"sku,omitempty"`
		RegularPrice      string `json:"regular_price,omitempty"`
		SalePrice         string `json:"sale_price,omitempty"`
		DateOnSaleFrom    string `json:"date_on_sale_from,omitempty"`
		DateOnSaleFromGmt string `json:"date_on_sale_from_gmt,omitempty"`
		DateOnSaleTo      string `json:"date_on_sale_to,omitempty"`
		DateOnSaleToGmt   string `json:"date_on_sale_to_gmt,omitempty"`
		Virtual           bool   `json:"virtual,omitempty"`
		Downloadable      bool   `json:"downloadable,omitempty"`
		Downloads         []any  `json:"downloads,omitempty"`
		DownloadLimit     int    `json:"download_limit,omitempty"`
		DownloadExpiry    int    `json:"download_expiry,omitempty"`
		ExternalURL       string `json:"external_url,omitempty"`
		ButtonText        string `json:"button_text,omitempty"`
		TaxStatus         string `json:"tax_status,omitempty"`
		TaxClass          string `json:"tax_class,omitempty"`
		ManageStock       bool   `json:"manage_stock,omitempty"`
		StockQuantity     any    `json:"stock_quantity,omitempty"`
		StockStatus       string `json:"stock_status,omitempty"`
		BackOrders        string `json:"backorders,omitempty"`
		SoldIndividually  bool   `json:"sold_individually,omitempty"`
		Weight            string `json:"weight,omitempty"`
		Dimensions        *struct {
			Length string `json:"length,omitempty"`
			Width  string `json:"width,omitempty"`
			Height string `json:"height,omitempty"`
		} `json:"dimensions,omitempty"`
		ShippingClass  string `json:"shipping_class,omitempty"`
		ReviewsAllowed bool   `json:"reviews_allowed,omitempty"`
		UpsellIDs      []int  `json:"upsell_ids,omitempty"`
		CrossSellIDs   []int  `json:"cross_sell_ids,omitempty"`
		ParentID       int    `json:"parent_id,omitempty"`
		PurchaseNote   string `json:"purchase_note,omitempty"`
		Categories     []struct {
			ID int `json:"id"`
		} `json:"categories,omitempty"`
		Tags []struct {
			ID int `json:"id"`
		} `json:"tags,omitempty"`
		Images            []Image            `json:"images,omitempty"`
		Attributes        []ProductAttribute `json:"attributes,omitempty"`
		DefaultAttributes []any              `json:"default_attributes,omitempty"`
		GroupedProducts   []any              `json:"grouped_products,omitempty"`
		MenuOrder         int                `json:"menu_order,omitempty"`
		MetaData          []Metadata         `json:"meta_data,omitempty"`
	}
)
