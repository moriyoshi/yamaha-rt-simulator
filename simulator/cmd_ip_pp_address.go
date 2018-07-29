package simulator

import "context"

var CmdIpPpAddress = &CommandSpec{
	[]Token{TIp, TPp, TAddress}, 3,
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
