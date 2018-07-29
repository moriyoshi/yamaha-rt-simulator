package simulator

import "context"

var CmdSnmpCommunityReadWrite = &CommandSpec{
	[]Token{TSnmp, TCommunity, TReadWrite}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpCommunityReadWrite = &CommandSpec{
	[]Token{TNo, TSnmp, TCommunity, TReadWrite}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

