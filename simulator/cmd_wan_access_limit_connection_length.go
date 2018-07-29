package simulator

import "context"

var CmdWanAccessLimitConnectionLength = &CommandSpec{
	[]Token{TWan, TAccess, TLimit, TConnection, TLength}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoWanAccessLimitConnectionLength = &CommandSpec{
	[]Token{TNo, TWan, TAccess, TLimit, TConnection, TLength}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

