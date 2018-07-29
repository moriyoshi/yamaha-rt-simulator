package simulator

import "context"

var CmdMobileAccessLimitTime = &CommandSpec{
	[]Token{TMobile, TAccess, TLimit, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMobileAccessLimitTimeTime = &CommandSpec{
	[]Token{TNo, TMobile, TAccess, TLimit, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

