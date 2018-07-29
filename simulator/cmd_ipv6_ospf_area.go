package simulator

import "context"

var CmdIpv6OspfArea = &CommandSpec{
	[]Token{TIpv6, TOspf, TArea}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6OspfArea = &CommandSpec{
	[]Token{TNo, TIpv6, TOspf, TArea}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

