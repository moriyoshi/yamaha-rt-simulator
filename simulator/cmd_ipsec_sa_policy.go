package simulator

import "context"

var CmdIpsecSaPolicy = &CommandSpec{
	[]Token{TIpsec, TSa, TPolicy}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecSaPolicy = &CommandSpec{
	[]Token{TNo, TIpsec, TSa, TPolicy}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
