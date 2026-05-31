package udp

import (
	"github.com/homestuck-ng/xray-core/common"
	"github.com/homestuck-ng/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
