package exported

import (
	"github.com/AlexY0905/specs-actors-yst-1/actors/builtin/account"
	"github.com/AlexY0905/specs-actors-yst-1/actors/builtin/cron"
	init_ "github.com/AlexY0905/specs-actors-yst-1/actors/builtin/init"
	"github.com/AlexY0905/specs-actors-yst-1/actors/builtin/market"
	"github.com/AlexY0905/specs-actors-yst-1/actors/builtin/miner"
	"github.com/AlexY0905/specs-actors-yst-1/actors/builtin/multisig"
	"github.com/AlexY0905/specs-actors-yst-1/actors/builtin/paych"
	"github.com/AlexY0905/specs-actors-yst-1/actors/builtin/power"
	"github.com/AlexY0905/specs-actors-yst-1/actors/builtin/reward"
	"github.com/AlexY0905/specs-actors-yst-1/actors/builtin/system"
	"github.com/AlexY0905/specs-actors-yst-1/actors/builtin/verifreg"
	"github.com/AlexY0905/specs-actors-yst-1/actors/runtime"
)

func BuiltinActors() []runtime.VMActor {
	return []runtime.VMActor{
		account.Actor{},
		cron.Actor{},
		init_.Actor{},
		market.Actor{},
		miner.Actor{},
		multisig.Actor{},
		paych.Actor{},
		power.Actor{},
		reward.Actor{},
		system.Actor{},
		verifreg.Actor{},
	}
}
