package simulator

import "context"

var CmdCopyConfig = &CommandSpec{
	[]Token{TCopy, TConfig}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
