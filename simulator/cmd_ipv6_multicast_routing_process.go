package simulator

import "context"

var CmdIpv6MulticastRoutingProcess = &CommandSpec{
	[]Token{TIpv6, TMulticast, TRouting, TProcess}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6MulticastRoutingProcess = &CommandSpec{
	[]Token{TNo, TIpv6, TMulticast, TRouting, TProcess}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

