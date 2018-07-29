package simulator

import "context"

var CmdIsdnCallbackMscbcpUserSpecify = &CommandSpec{
	[]Token{TIsdn, TCallback, TMscbcp, TUserSpecify}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNoIsdnCallbackMscbcpUserSpecify = &CommandSpec{
	[]Token{TNo, TNo, TIsdn, TCallback, TMscbcp, TUserSpecify}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

