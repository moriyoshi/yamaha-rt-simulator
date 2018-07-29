package simulator

import "context"

var CmdOperationHttpRevisionUpPermit = &CommandSpec{
	[]Token{TOperation, THttp, TRevisionUp, TPermit}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoOperationHttpRevisionUpPermit = &CommandSpec{
	[]Token{TNo, TOperation, THttp, TRevisionUp, TPermit}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

