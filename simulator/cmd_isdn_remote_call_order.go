package simulator

import "context"

var CmdIsdnRemoteCallOrder = &CommandSpec{
	[]Token{TIsdn, TRemote, TCall, TOrder}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnRemoteCallOrder = &CommandSpec{
	[]Token{TNo, TIsdn, TRemote, TCall, TOrder}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

