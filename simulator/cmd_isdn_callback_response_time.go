package simulator

import "context"

var CmdIsdnCallbackResponseTime = &CommandSpec{
	[]Token{TIsdn, TCallback, TResponse, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnCallbackResponseTime = &CommandSpec{
	[]Token{TNo, TIsdn, TCallback, TResponse, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

