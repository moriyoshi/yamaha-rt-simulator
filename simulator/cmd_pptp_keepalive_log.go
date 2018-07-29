package simulator

import "context"

var CmdPptpKeepaliveLog = &CommandSpec{
	[]Token{TPptp, TKeepalive, TLog}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoPptpKeepaliveLog = &CommandSpec{
	[]Token{TNo, TPptp, TKeepalive, TLog}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
