package simulator

import "context"

var CmdSnmpv2CTrapHost = &CommandSpec{
	[]Token{TSnmpv2C, TTrap, THost}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpv2CTrapHost = &CommandSpec{
	[]Token{TNo, TSnmpv2C, TTrap, THost}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

