package Controller

import logger "product-service/Logger"

type ControllerInterface interface {
	ProductInterface
}

type Controller struct {
	ControllerInterface
	Log *logger.Logger
}
