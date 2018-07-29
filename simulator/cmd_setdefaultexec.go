package simulator

import "context"

var CmdSetDefaultExec = &CommandSpec{
	[]Token{TSetDefaultExec}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

