package simulator

import "context"

var CmdIsdnCallbackPermit = &CommandSpec{
	[]Token{TIsdn, TCallback, TPermit}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnCallbackPermit = &CommandSpec{
	[]Token{TNo, TIsdn, TCallback, TPermit}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

