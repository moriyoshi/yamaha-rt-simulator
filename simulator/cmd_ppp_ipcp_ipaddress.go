package simulator

import "context"

var CmdPppIpcpIpaddress = &CommandSpec{
	[]Token{TPpp, TIpcp, TIpaddress}, 3,
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

var CmdNoPppIpcpIpaddress = &CommandSpec{
	[]Token{TNo, TPpp, TIpcp, TIpaddress}, 4,
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
