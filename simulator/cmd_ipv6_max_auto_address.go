package simulator

import "context"

var CmdIpv6MaxAutoAddress = &CommandSpec{
	[]Token{TIpv6, TMax, TAuto, TAddress}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6MaxAutoAddress = &CommandSpec{
	[]Token{TNo, TIpv6, TMax, TAuto, TAddress}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

