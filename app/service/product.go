package service

import (
	"github.com/micrease/gorme"
	mctx "github.com/micrease/micrease-core/context"
	"github.com/micrease/micrease-core/errs"
	"meshop-api/app/model"
	"meshop-api/app/repository"
)

type Product struct {
}

func (this *Product) PageList(ctx *mctx.Context) *gorme.PageResult[model.Product] {
	repo := repository.NewProduct(ctx.Orm)
	result, err := repo.Paginate(1, 10)
	errs.PanicIf(err, 5000, "查询失败")

	return result
}
