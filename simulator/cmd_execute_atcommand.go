package simulator

import "context"

var CmdExecuteAtCommand = &CommandSpec{
	[]Token{TExecute, TAtCommand}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

