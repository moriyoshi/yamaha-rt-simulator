package simulator

import "context"

var CmdSnmpv2CHost = &CommandSpec{
	[]Token{TSnmpv2C, THost}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpv2CHostHost = &CommandSpec{
	[]Token{TNo, TSnmpv2C, THost}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

