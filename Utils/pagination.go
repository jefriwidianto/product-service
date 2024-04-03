package Utils

import (
	"net/url"
	"product-service/Controller/Dto/Request"
	"strconv"
)

func Pagination(param url.Values) (resp Request.ListParam, err error) {
	resp.Limit = 10
	resp.Page = 0

	if param.Get("page") != "" {
		page, err := strconv.Atoi(param.Get("page"))
		if err != nil {
			return resp, err
		}

		resp.Page = page - 1
	}

	if param.Get("limit") != "" {
		limit, err := strconv.Atoi(param.Get("limit"))
		if err != nil {
			return resp, err
		}

		resp.Limit = limit
	}

	resp.Offset = resp.Limit * resp.Page
	return resp, err
}
