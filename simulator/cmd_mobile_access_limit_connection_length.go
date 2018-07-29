package simulator

import "context"

var CmdMobileAccessLimitConnectionLength = &CommandSpec{
	[]Token{TMobile, TAccess, TLimit, TConnection, TLength}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMobileAccessLimitConnectionLength = &CommandSpec{
	[]Token{TNo, TMobile, TAccess, TLimit, TConnection, TLength}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

