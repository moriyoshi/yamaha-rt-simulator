package simulator

import "context"

var CmdIpv6RipPreference = &CommandSpec{
	[]Token{TIpv6, TRip, TPreference}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6RipPreference = &CommandSpec{
	[]Token{TNo, TIpv6, TRip, TPreference}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
