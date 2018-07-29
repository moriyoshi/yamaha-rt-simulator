package simulator

import "context"

var CmdIsdnCallbackRequest = &CommandSpec{
	[]Token{TIsdn, TCallback, TRequest}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnCallbackRequest = &CommandSpec{
	[]Token{TNo, TIsdn, TCallback, TRequest}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

