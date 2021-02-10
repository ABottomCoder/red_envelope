package red_envelope

import (
	_ "github.com/ABottomCoder/account/core/accounts"
	"github.com/ABottomCoder/infra"
	"github.com/ABottomCoder/infra/base"
	"github.com/red_envelope/apis/gorpc"
	_ "github.com/red_envelope/apis/gorpc"
	_ "github.com/red_envelope/apis/web"
	_ "github.com/red_envelope/core/envelopes"
	"github.com/red_envelope/jobs"
)

func init() {
	infra.Register(&base.PropsStarter{})
	infra.Register(&base.DbxDatabaseStarter{})
	infra.Register(&base.ValidatorStarter{})
	infra.Register(&base.GoRPCStarter{})
	infra.Register(&gorpc.GoRpcApiStarter{})
	infra.Register(&jobs.RefundExpiredJobStarter{})
	infra.Register(&base.IrisServerStarter{})
	infra.Register(&infra.WebApiStarter{})
	infra.Register(&base.EurekaStarter{})
	infra.Register(&base.HookStarter{})
}
