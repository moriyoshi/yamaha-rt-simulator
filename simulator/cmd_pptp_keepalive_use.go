package simulator

import "context"

var CmdPptpKeepaliveUse = &CommandSpec{
	[]Token{TPptp, TKeepalive, TUse}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoPptpKeepaliveUse = &CommandSpec{
	[]Token{TNo, TPptp, TKeepalive, TUse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
