// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graph

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/talgat-ruby/lessons-go/projects/expense-tracker-micro/api/internal/graphql/graph/model"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) unmarshalNCategory2githubᚗcomᚋtalgatᚑrubyᚋlessonsᚑgoᚋprojectsᚋexpenseᚑtrackerᚋinternalᚋgraphqlᚋgraphᚋmodelᚐCategory(ctx context.Context, v any) (model.Category, error) {
	var res model.Category
	err := res.UnmarshalGQL(v)
	return res, graphql.ErrorOnPath(ctx, err)
}

func (ec *executionContext) marshalNCategory2githubᚗcomᚋtalgatᚑrubyᚋlessonsᚑgoᚋprojectsᚋexpenseᚑtrackerᚋinternalᚋgraphqlᚋgraphᚋmodelᚐCategory(ctx context.Context, sel ast.SelectionSet, v model.Category) graphql.Marshaler {
	return v
}

// endregion ***************************** type.gotpl *****************************
