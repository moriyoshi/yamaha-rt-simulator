package simulator

import "context"

var CmdHttpRevisionUpPermit = &CommandSpec{
	[]Token{THttp, TRevisionUp, TPermit}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHttpRevisionUpPermit = &CommandSpec{
	[]Token{TNo, THttp, TRevisionUp, TPermit}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

