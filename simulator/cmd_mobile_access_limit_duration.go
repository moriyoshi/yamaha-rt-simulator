package simulator

import "context"

var CmdMobileAccessLimitDuration = &CommandSpec{
	[]Token{TMobile, TAccess, TLimit, TDuration}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMobileAccessLimitDuration = &CommandSpec{
	[]Token{TNo, TMobile, TAccess, TLimit, TDuration}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

