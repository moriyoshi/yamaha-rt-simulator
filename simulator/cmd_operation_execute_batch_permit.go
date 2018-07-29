package simulator

import "context"

var CmdOperationExecuteBatchPermit = &CommandSpec{
	[]Token{TOperation, TExecute, TBatch, TPermit}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoOperationExecuteBatchPermit = &CommandSpec{
	[]Token{TNo, TOperation, TExecute, TBatch, TPermit}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

