package simulator

import "context"

var CmdBgpRouterId = &CommandSpec{
	[]Token{TBgp, TRouter, TId}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoBgpRouterIdIp_Address = &CommandSpec{
	[]Token{TNo, TBgp, TRouter, TId}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

