package simulator

import "context"

var CmdWanAccessLimitConnectionTime = &CommandSpec{
	[]Token{TWan, TAccess, TLimit, TConnection, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoWanAccessLimitConnectionTime = &CommandSpec{
	[]Token{TNo, TWan, TAccess, TLimit, TConnection, TTime}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

