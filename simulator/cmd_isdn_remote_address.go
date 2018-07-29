package simulator

import "context"

var CmdIsdnRemoteAddress = &CommandSpec{
	[]Token{TIsdn, TRemote, TAddress}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIsdnRemoteAddress = &CommandSpec{
	[]Token{TNo, TIsdn, TRemote, TAddress}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
