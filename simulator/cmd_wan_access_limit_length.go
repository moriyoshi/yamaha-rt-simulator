package simulator

import "context"

var CmdWanAccessLimitLength = &CommandSpec{
	[]Token{TWan, TAccess, TLimit, TLength}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoWanAccessLimitLength = &CommandSpec{
	[]Token{TNo, TWan, TAccess, TLimit, TLength}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

