package midtrans

import "github.com/midtrans/midtrans-go/coreapi"

type (
	Midtrans interface {
	}

	midtransImpl struct {
		client coreapi.Client
	}
)
