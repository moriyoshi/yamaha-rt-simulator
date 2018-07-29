package simulator

import "context"

var CmdIsdnLocalAddress = &CommandSpec{
	[]Token{TIsdn, TLocal, TAddress}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoisdnLocalAddress = &CommandSpec{
	[]Token{TNo, TIsdn, TLocal, TAddress}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
