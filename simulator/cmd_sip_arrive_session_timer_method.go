package simulator

import "context"

var CmdSipArriveSessionTimerMethod = &CommandSpec{
	[]Token{TSip, TArrive, TSession, TTimer, TMethod}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipArriveSessionTimerMethod = &CommandSpec{
	[]Token{TNo, TSip, TArrive, TSession, TTimer, TMethod}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

