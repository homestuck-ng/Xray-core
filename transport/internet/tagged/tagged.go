package tagged

import (
	"context"

	"github.com/homestuck-ng/xray-core/common/net"
	"github.com/homestuck-ng/xray-core/features/routing"
)

type DialFunc func(ctx context.Context, dispatcher routing.Dispatcher, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc
