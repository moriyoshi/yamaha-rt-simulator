package simulator

import "context"

var CmdPppIpcpRemoteAddressCheck = &CommandSpec{
	[]Token{TPpp, TIpcp, TRemote, TAddress, TCheck}, 5,
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

var CmdNoPppIpcpRemoteAddressCheck = &CommandSpec{
	[]Token{TNo, TPpp, TIpcp, TRemote, TAddress, TCheck}, 6,
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
