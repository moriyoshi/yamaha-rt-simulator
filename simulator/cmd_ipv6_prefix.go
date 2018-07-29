package simulator

import "context"

var CmdIpv6Prefix = &CommandSpec{
	[]Token{TIpv6, TPrefix}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6Prefix = &CommandSpec{
	[]Token{TNo, TIpv6, TPrefix}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
