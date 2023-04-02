package model

type (
	// OrderType type
	OrderType string
	// OrderBy type
	OrderBy string
	// Status type
	Status string
	// ProductType type
	ProductType string
	// StockStatus type
	StockStatus string
	// TaxClass type
	TaxClass string

	// ErrorPayload schema
	ErrorPayload struct {
		Code    string `mapstructure:"code" json:"code"`
		Message string `mapstructure:"message" json:"message"`
		Data    struct {
			Status int `mapstructure:"status" json:"status"`
		} `mapstructure:"data" json:"data"`
	}
)

const (
	// LogStrRequest log string key
	LogStrRequest = "request"
	// LogStrKeyEndpoint log endpoint name value
	LogStrKeyEndpoint = "endpoint"

	// BaseURL string
	BaseURL = "https://wp-dev.sendlovebox.xyz/wp-json/wc/v3"

	// OrderTypeAsc ascending
	OrderTypeAsc OrderType = "asc"
	// OrderTypeDesc descending
	OrderTypeDesc OrderType = "desc"

	// OrderByDate to order by date
	OrderByDate OrderBy = "date"
	// OrderByID to order by id
	OrderByID OrderBy = "id"
	// OrderByInclude to order by include
	OrderByInclude OrderBy = "include"
	// OrderByTitle to order by title
	OrderByTitle OrderBy = "title"
	// OrderBySlug to order by slug
	OrderBySlug OrderBy = "slug"
	// OrderByPrice to order by price
	OrderByPrice OrderBy = "price"
	// OrderByPopularity to order by popularity
	OrderByPopularity OrderBy = "popularity"
	// OrderByRating to order by rating
	OrderByRating OrderBy = "rating"

	// StatusAny for all statuses
	StatusAny Status = "any"
	// StatusDraft for draft statuses
	StatusDraft Status = "draft"
	// StatusPending for pending statuses
	StatusPending Status = "pending"
	// StatusPrivate for private statuses
	StatusPrivate Status = "private"
	// StatusPublish for published statuses
	StatusPublish Status = "publish"

	// ProductTypeSimple for simple products
	ProductTypeSimple ProductType = "simple"
	// ProductTypeGrouped for grouped products
	ProductTypeGrouped ProductType = "grouped"
	// ProductTypeExternal for external products
	ProductTypeExternal ProductType = "external"
	// ProductTypeVariable for variable products
	ProductTypeVariable ProductType = "variable"

	// TaxClassStandard for standard tax class
	TaxClassStandard TaxClass = "standard"
	// TaxClassReducedRate for reduced-rate tax class
	TaxClassReducedRate TaxClass = "reduced-rate"
	// TaxClassZeroRate for zero-rate tax class
	TaxClassZeroRate TaxClass = "zero-rate"

	// StockStatusInStock for in stock products
	StockStatusInStock StockStatus = "instock"
	// StockStatusOutOfStock for out of stock products
	StockStatusOutOfStock StockStatus = "outofstock"
	// StockStatusOnBackOrder for on back order products
	StockStatusOnBackOrder StockStatus = "onbackorder"
)
