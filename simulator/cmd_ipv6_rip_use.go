package simulator

import "context"

var CmdIpv6RipUse = &CommandSpec{
	[]Token{TIpv6, TRip, TUse}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6RipUse = &CommandSpec{
	[]Token{TNo, TIpv6, TRip, TUse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
