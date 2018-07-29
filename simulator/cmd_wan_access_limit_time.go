package simulator

import "context"

var CmdWanAccessLimitTime = &CommandSpec{
	[]Token{TWan, TAccess, TLimit, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoWanAccessLimitTime = &CommandSpec{
	[]Token{TNo, TWan, TAccess, TLimit, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

