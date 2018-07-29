package simulator

import "context"

var CmdIpv6Routing = &CommandSpec{
	[]Token{TIpv6, TRouting}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6Routing = &CommandSpec{
	[]Token{TNo, TIpv6, TRouting}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
