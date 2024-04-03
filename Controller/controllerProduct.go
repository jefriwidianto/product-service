package Controller

import (
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
	"product-service/Controller/Dto/Response"
	"product-service/Repository"
	"product-service/Utils"
)

type ProductInterface interface {
	ListProduct(ctx echo.Context) (err error)
}

func (c Controller) ListProduct(ctx echo.Context) (err error) {
	pagination, err := Utils.Pagination(ctx.QueryParams())
	if err != nil {
		c.Log.Error("invalid parameter limit or page", zap.String("_method", ctx.Request().Method), zap.String("_url_request", ctx.Request().RequestURI))
		return ctx.JSON(http.StatusBadRequest, &Response.Responses{
			Message: http.StatusText(http.StatusBadRequest),
		})
	}

	resp, err := Repository.ApplicationRepository.Product.ListProduct(ctx.Request().Context(), pagination)
	if err != nil {
		c.Log.Error(http.StatusText(http.StatusInternalServerError), zap.String("_method", ctx.Request().Method), zap.String("_url_request", ctx.Request().RequestURI))
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response.Responses{
		Data:    resp,
		Message: http.StatusText(http.StatusOK),
	})
}
