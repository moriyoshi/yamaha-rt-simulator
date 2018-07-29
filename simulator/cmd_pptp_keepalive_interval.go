package simulator

import "context"

var CmdPptpKeepaliveInterval = &CommandSpec{
	[]Token{TPptp, TKeepalive, TInterval}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoPptpKeepaliveInterval = &CommandSpec{
	[]Token{TNo, TPptp, TKeepalive, TInterval}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
