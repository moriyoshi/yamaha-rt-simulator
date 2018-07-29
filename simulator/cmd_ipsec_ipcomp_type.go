package simulator

import "context"

var CmdIpsecIpcompType = &CommandSpec{
	[]Token{TIpsec, TIpcomp, TType}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIpcompType = &CommandSpec{
	[]Token{TNo, TIpsec, TIpcomp, TType}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

