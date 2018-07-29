package simulator

import "context"

var CmdSnmpv2CCommunityReadWrite = &CommandSpec{
	[]Token{TSnmpv2C, TCommunity, TReadWrite}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpv2CCommunityReadWrite = &CommandSpec{
	[]Token{TNo, TSnmpv2C, TCommunity, TReadWrite}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

