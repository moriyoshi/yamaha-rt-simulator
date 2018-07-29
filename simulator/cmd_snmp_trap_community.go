package simulator

import "context"

var CmdSnmpTrapCommunity = &CommandSpec{
	[]Token{TSnmp, TTrap, TCommunity}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpTrapCommunity = &CommandSpec{
	[]Token{TNo, TSnmp, TTrap, TCommunity}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

