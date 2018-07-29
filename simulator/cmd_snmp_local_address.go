package simulator

import "context"

var CmdSnmpLocalAddress = &CommandSpec{
	[]Token{TSnmp, TLocal, TAddress}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpLocalAddress = &CommandSpec{
	[]Token{TNo, TSnmp, TLocal, TAddress}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

