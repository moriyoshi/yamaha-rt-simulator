package simulator

import "context"

var CmdIpv6Route = &CommandSpec{
	[]Token{TIpv6, TRoute}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6Route = &CommandSpec{
	[]Token{TNo, TIpv6, TRoute}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
