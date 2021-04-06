package resolvers

// go:generate go run github.com/vektah/dataloaden ProductsOrgLoader string []*mini-seller/application/graph/model.Product

import (
	"context"
	"mini-seller/application/graph/model"
	"mini-seller/domain/packages/catalogpkg"
	"net/http"
	"time"

	"github.com/prometheus/common/log"
)

type ctxKeyType struct{ name string }

var ctxKey = ctxKeyType{"catalogCtx"}

type loaders struct {
	productsByOrganization *ProductsOrgLoader
}

// nolint: gosec
func LoaderMiddleware(catalogUseCase catalogpkg.IUseCase, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ldrs := loaders{}

		// set this to zero what happens without dataloading
		wait := 100 * time.Microsecond

		ldrs.productsByOrganization = &ProductsOrgLoader{
			wait:     wait,
			maxBatch: 100,
			fetch: func(keys []string) ([][]*model.Product, []error) {
				errors := make([]error, len(keys))

				prodList, err := catalogUseCase.GetProductList(r.Context(), keys, []string{})
				if err != nil {
					log.Warn(err)
					for i := range keys {
						errors[i] = err
					}
					return nil, errors
				}
				if prodList == nil {
					return nil, errors
				}

				iByKey := make(map[string]int)
				for i, key := range keys {
					iByKey[key] = i
				}

				products := make([][]*model.Product, len(keys))
				for _, prod := range prodList {
					product := &model.Product{
						ID:             prod.ID,
						IDCategory:     prod.IDCategory,
						IDOrganization: prod.IDOrganization,
						Name:           prod.Name,
						Price:          prod.Price,
						Count:          prod.Count,
						Status:         prod.Status,
					}

					i := iByKey[product.IDOrganization]

					products[i] = append(products[i], product)
				}

				return products, errors
			},
		}

		dlCtx := context.WithValue(r.Context(), ctxKey, ldrs)
		next.ServeHTTP(w, r.WithContext(dlCtx))
	})
}

func ctxLoaders(ctx context.Context) loaders {
	return ctx.Value(ctxKey).(loaders)
}
