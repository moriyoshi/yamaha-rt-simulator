package simulator

import "context"

var CmdIsdnCallbackWaitTime = &CommandSpec{
	[]Token{TIsdn, TCallback, TWait, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnCallbackWaitTime = &CommandSpec{
	[]Token{TNo, TIsdn, TCallback, TWait, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

