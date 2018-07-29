package simulator

import "context"

var CmdSnmpCommunityReadOnly = &CommandSpec{
	[]Token{TSnmp, TCommunity, TReadOnly}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpCommunityReadOnly = &CommandSpec{
	[]Token{TNo, TSnmp, TCommunity, TReadOnly}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

