package productpkg

import "errors"

var (
	// ErrProductNotFound - not found
	ErrProductNotFound = errors.New("product_not_found")
	ErrListLimit       = errors.New("list_limit")
)
