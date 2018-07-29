package simulator

import "context"

var CmdSnmpTrapHost = &CommandSpec{
	[]Token{TSnmp, TTrap, THost}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpTrapHost = &CommandSpec{
	[]Token{TNo, TSnmp, TTrap, THost}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

