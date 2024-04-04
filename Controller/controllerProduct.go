package Controller

import (
	"errors"
	"github.com/labstack/echo"
	"go.uber.org/zap"
	"net/http"
	"product-service/Controller/Dto/Response"
	"product-service/Repository"
	"product-service/Utils"
)

type ProductInterface interface {
	ListProduct(ctx echo.Context) (err error)
	DetailProduct(ctx echo.Context) (err error)
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

func (c Controller) DetailProduct(ctx echo.Context) (err error) {
	productId := ctx.Param("id")
	exists, err := Repository.ApplicationRepository.Product.CheckExistsProductId(ctx.Request().Context(), productId)
	if err != nil {
		c.Log.Error(err.Error(), zap.String("_method", ctx.Request().Method), zap.String("_url_request", ctx.Request().RequestURI))
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Message: err.Error(),
		})
	}

	if !exists {
		c.Log.Error("Product Id Not Found", zap.String("_method", ctx.Request().Method), zap.String("_url_request", ctx.Request().RequestURI))
		return ctx.JSON(http.StatusBadRequest, &Response.Responses{
			Message: errors.New("Product Id Not Found").Error(),
		})
	}

	data, err := Repository.ApplicationRepository.Product.DetailProduct(ctx.Request().Context(), productId)
	if err != nil {
		c.Log.Error(err.Error(), zap.String("_method", ctx.Request().Method), zap.String("_url_request", ctx.Request().RequestURI))
		return ctx.JSON(http.StatusInternalServerError, &Response.Responses{
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, &Response.Responses{
		Data:    data,
		Message: http.StatusText(http.StatusOK),
	})
}
