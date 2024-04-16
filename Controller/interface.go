package Controller

import (
	"product-service/Controller/Dto/Proto"
	logger "product-service/Logger"
	"sync"
)

type ControllerInterface interface {
	ProductInterface
}

type Controller struct {
	ControllerInterface
	Log *logger.Logger
}

type DataProductServer struct {
	Proto.UnimplementedDataProductServer
	mu          sync.Mutex
	dataProduct *Proto.Product
}
