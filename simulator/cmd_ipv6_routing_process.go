package simulator

import "context"

var CmdIpv6RoutingProcess = &CommandSpec{
	[]Token{TIpv6, TRouting, TProcess}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6RoutingProcess = &CommandSpec{
	[]Token{TNo, TIpv6, TRouting, TProcess}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
