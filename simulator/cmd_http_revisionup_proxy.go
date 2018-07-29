package simulator

import "context"

var CmdHttpRevisionUpProxy = &CommandSpec{
	[]Token{THttp, TRevisionUp, TProxy}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHttpRevisionUpProxy = &CommandSpec{
	[]Token{TNo, THttp, TRevisionUp, TProxy}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

