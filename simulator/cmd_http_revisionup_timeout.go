package simulator

import "context"

var CmdHttpRevisionUpTimeout = &CommandSpec{
	[]Token{THttp, TRevisionUp, TTimeout}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHttpRevisionUpTimeout = &CommandSpec{
	[]Token{TNo, THttp, TRevisionUp, TTimeout}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

