package simulator

import "context"

var CmdIsdnCallbackPermitType = &CommandSpec{
	[]Token{TIsdn, TCallback, TPermit, TType}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnCallbackPermitType = &CommandSpec{
	[]Token{TNo, TIsdn, TCallback, TPermit, TType}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

