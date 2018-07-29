package simulator

import "context"

var CmdCopyExec = &CommandSpec{
	[]Token{TCopy, TExec}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

