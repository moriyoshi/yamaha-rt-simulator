package simulator

import "context"

var CmdIpv6OspfUse = &CommandSpec{
	[]Token{TIpv6, TOspf, TUse}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6OspfUse = &CommandSpec{
	[]Token{TNo, TIpv6, TOspf, TUse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

