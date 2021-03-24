package catalogcontroller

import (
	"context"
	"fmt"
	"mini-seller/domain/common/entities/catalogentity"
	"mini-seller/domain/packages/customer/catalogpkg"
	"mini-seller/web/customer/fields/catalogfield"

	"github.com/graphql-go/graphql"
	"github.com/prometheus/common/log"
)

type OrganizationListResolver struct {
	catalogUseCase catalogpkg.IUseCase
}

func NewOrganizationListResolver(catalogUseCase catalogpkg.IUseCase) *OrganizationListResolver {
	return &OrganizationListResolver{catalogUseCase: catalogUseCase}
}

/*
func (olr OrganizationListResolver) GetOrganizationList(ctx context.Context) []catalogfield.OrganizationInfo {
	orgList, err := olr.catalogUseCase.GetOrganizationList(ctx)
	if err != nil {
		return nil, err
	}

	return orgList, nil
}
*/

func (olr OrganizationListResolver) GetOrganizationList() *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(catalogfield.OrganizationInfo),
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			orgList, err := olr.catalogUseCase.GetOrganizationList(p.Context)

			detailResolvers := make([]*OrganizationResolver, 0)
			for _, org := range orgList {
				resolver := OrganizationResolver{
					ctx:            p.Context,
					catalogUseCase: olr.catalogUseCase,
					organization:   org,
					Name:           org.Name,
				}

				detailResolvers = append(detailResolvers, &resolver)
			}

			fmt.Println("detailResolvers", detailResolvers)

			return detailResolvers, err
		},
	}
}

type OrganizationResolver struct {
	ctx            context.Context
	catalogUseCase catalogpkg.IUseCase
	organization   *catalogentity.OrganizationInfo
	Name           string
}

func (or OrganizationResolver) ID() string {
	fmt.Println("ID", or.organization.ID)

	return or.organization.ID
}

// func (or OrganizationResolver) Name() string {
// 	return or.organization.Name
// }

func (or *OrganizationResolver) Email() string {
	return or.organization.Email
}
func (or OrganizationResolver) Phone() string {
	return or.organization.Phone
}
func (or OrganizationResolver) Status() string {
	return or.organization.Status
}
func (or OrganizationResolver) Categories() []*catalogentity.CategoryInfo {
	fmt.Println(1)
	catList, err := or.catalogUseCase.GetCategoryList(or.ctx)
	if err != nil {
		log.Warn(err)
	}
	fmt.Println(2, catList)
	return catList
}
func (or OrganizationResolver) Products() []*catalogentity.ProductInfo {
	prodList, _ := or.catalogUseCase.GetProductList(or.ctx, or.organization.ID, "")
	return prodList
}
