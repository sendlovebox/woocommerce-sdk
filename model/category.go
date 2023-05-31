package model

type (
	// SearchCategoriesRequest params to search categories
	SearchCategoriesRequest struct {
		PaginationRequest

		Context   string   `url:"context,omitempty" json:"context"`
		Search    string   `url:"search,omitempty" json:"search"`
		Exclude   []string `url:"exclude,omitempty" json:"exclude"`
		Include   []string `url:"include,omitempty" json:"include"`
		Parent    int      `url:"parent,omitempty" json:"parent"`
		Product   int      `url:"product,omitempty" json:"product"`
		HideEmpty bool     `url:"hide_empty,omitempty" json:"hide_empty"`
		Slug      string   `url:"slug,omitempty" json:"slug"`
	}

	// Category schema
	Category struct {
		ID          int    `mapstructure:"id" json:"id"`
		Name        string `mapstructure:"name" json:"name"`
		Slug        string `mapstructure:"slug" json:"slug"`
		Parent      int    `mapstructure:"parent" json:"parent"`
		Description string `mapstructure:"description" json:"description"`
		Display     string `mapstructure:"display" json:"display"`
		Image       Image  `mapstructure:"image,omitempty" json:"image,omitempty"`
		MenuOrder   int    `mapstructure:"menu_order" json:"menu_order"`
		Count       int    `mapstructure:"count" json:"count"`
		Links       struct {
			Self []struct {
				Href string `mapstructure:"href" json:"href"`
			} `mapstructure:"self" json:"self"`
			Collection []struct {
				Href string `mapstructure:"href" json:"href"`
			} `mapstructure:"collection" json:"collection"`
			Up []struct {
				Href string `mapstructure:"href" json:"href"`
			} `mapstructure:"up" json:"up"`
		} `mapstructure:"_links" json:"_links"`
	}
)
