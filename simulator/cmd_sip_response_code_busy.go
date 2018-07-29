package simulator

import "context"

var CmdSipResponseCodeBusy = &CommandSpec{
	[]Token{TSip, TResponse, TCode, TBusy}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSipResponseCodeBusy = &CommandSpec{
	[]Token{TNo, TSip, TResponse, TCode, TBusy}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

