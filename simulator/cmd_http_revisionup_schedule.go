package simulator

import "context"

var CmdHttpRevisionUpSchedule = &CommandSpec{
	[]Token{THttp, TRevisionUp, TSchedule}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHttpRevisionUpSchedule = &CommandSpec{
	[]Token{TNo, THttp, TRevisionUp, TSchedule}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

