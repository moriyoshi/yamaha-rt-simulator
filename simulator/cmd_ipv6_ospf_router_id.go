package simulator

import "context"

var CmdIpv6OspfRouterId = &CommandSpec{
	[]Token{TIpv6, TOspf, TRouter, TId}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6OspfRouterId = &CommandSpec{
	[]Token{TNo, TIpv6, TOspf, TRouter, TId}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

