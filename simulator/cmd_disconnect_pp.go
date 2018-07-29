package simulator

import "context"

var CmdDisconnectPp = &CommandSpec{
	[]Token{TDisconnect, TPp}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
