package simulator

import "context"

var CmdExecuteBatch = &CommandSpec{
	[]Token{TExecute, TBatch}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

