package simulator

import "context"

var CmdConnectPp = &CommandSpec{
	[]Token{TConnect, TPp}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
