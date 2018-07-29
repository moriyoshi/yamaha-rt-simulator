package simulator

import "context"

var CmdSnmpDisplayIpcpForce = &CommandSpec{
	[]Token{TSnmp, TDisplay, TIpcp, TForce}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpDisplayIpcpForce = &CommandSpec{
	[]Token{TNo, TSnmp, TDisplay, TIpcp, TForce}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

