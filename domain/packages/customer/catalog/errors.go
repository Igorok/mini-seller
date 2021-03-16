package catalog

import "errors"

var (
	// ErrProductNotFound - not found
	ErrProductNotFound = errors.New("product_not_found")
	ErrCatalogLimit    = errors.New("catalog_limit")
)
