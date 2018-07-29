package simulator

import "context"

var CmdSnmpv2CCommunityReadOnly = &CommandSpec{
	[]Token{TSnmpv2C, TCommunity, TReadOnly}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpv2CCommunityReadOnly = &CommandSpec{
	[]Token{TNo, TSnmpv2C, TCommunity, TReadOnly}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

