package simulator

import "context"

var CmdWanAccessLimitDuration = &CommandSpec{
	[]Token{TWan, TAccess, TLimit, TDuration}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoWanAccessLimitDurationDuration = &CommandSpec{
	[]Token{TNo, TWan, TAccess, TLimit, TDuration}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

