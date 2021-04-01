package catalogpkg

import "errors"

var (
	// ErrProductNotFound - not found
	ErrProductNotFound      = errors.New("product_not_found")
	ErrCategoryNotFound     = errors.New("category_not_found")
	ErrOrganizationNotFound = errors.New("organization_not_found")
)
