package simulator

import "context"

var CmdSnmpv3TrapHost = &CommandSpec{
	[]Token{TSnmpv3, TTrap, THost}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpv3TrapHost = &CommandSpec{
	[]Token{TNo, TSnmpv3, TTrap, THost}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

