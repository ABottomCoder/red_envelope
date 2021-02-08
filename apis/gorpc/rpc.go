package gorpc

import (
	"github.com/ABottomCoder/infra"
	"github.com/ABottomCoder/infra/base"
)

type GoRpcApiStarter struct {
	infra.BaseStarter
}

func (g *GoRpcApiStarter) Init(ctx infra.StarterContext) {
	base.RpcRegister(new(EnvelopeRpc))
}
