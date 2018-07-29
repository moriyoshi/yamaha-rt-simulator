package simulator

import "context"

var CmdPppCcpNoEncryption = &CommandSpec{
	[]Token{TPpp, TCcp, TNoEncryption}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.PP == 0 {
			return InvalidCommandName
		}
		return nil
	},
}

var CmdNoPppCcpNoEncryption = &CommandSpec{
	[]Token{TNo, TPpp, TCcp, TNoEncryption}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.PP == 0 {
			return InvalidCommandName
		}
		return nil
	},
}
