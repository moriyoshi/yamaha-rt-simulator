package simulator

import "context"

var CmdSnmpv2CTrapCommunity = &CommandSpec{
	[]Token{TSnmpv2C, TTrap, TCommunity}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpv2CTrapCommunity = &CommandSpec{
	[]Token{TNo, TSnmpv2C, TTrap, TCommunity}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

