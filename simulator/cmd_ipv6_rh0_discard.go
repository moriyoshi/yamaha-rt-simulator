package simulator

import "context"

var CmdIpv6Rh0Discard = &CommandSpec{
	[]Token{TIpv6, TRh0, TDiscard}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6Rh0Discard = &CommandSpec{
	[]Token{TNo, TIpv6, TRh0, TDiscard}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
