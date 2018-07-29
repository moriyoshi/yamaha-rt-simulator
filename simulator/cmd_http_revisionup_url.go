package simulator

import "context"

var CmdHttpRevisionUpUrl = &CommandSpec{
	[]Token{THttp, TRevisionUp, TUrl}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHttpRevisionUpUrl = &CommandSpec{
	[]Token{TNo, THttp, TRevisionUp, TUrl}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

