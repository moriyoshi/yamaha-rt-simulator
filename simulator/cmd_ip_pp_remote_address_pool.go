package simulator

import "context"

var CmdIpPpRemoteAddressPool = &CommandSpec{
	[]Token{TIp, TPp, TRemote, TAddress, TPool}, 5,
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

var CmdNoIpPpRemoteAddressPool = &CommandSpec{
	[]Token{TNo, TIp, TPp, TRemote, TAddress, TPool}, 6,
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
