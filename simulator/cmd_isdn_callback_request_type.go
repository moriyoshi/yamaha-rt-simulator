package simulator

import "context"

var CmdIsdnCallbackRequestType = &CommandSpec{
	[]Token{TIsdn, TCallback, TRequest, TType}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnCallbackRequestType = &CommandSpec{
	[]Token{TNo, TIsdn, TCallback, TRequest, TType}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

