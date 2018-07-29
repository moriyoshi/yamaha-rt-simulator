package simulator

import "context"

var CmdSipArriveSessionTimerRefresher = &CommandSpec{
	[]Token{TSip, TArrive, TSession, TTimer, TRefresher}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipArriveSessionTimerRefresher = &CommandSpec{
	[]Token{TNo, TSip, TArrive, TSession, TTimer, TRefresher}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

