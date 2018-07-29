package simulator

import "context"

var CmdIpsecUse = &CommandSpec{
	[]Token{TIpsec, TUse}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecUse = &CommandSpec{
	[]Token{TNo, TIpsec, TUse}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

