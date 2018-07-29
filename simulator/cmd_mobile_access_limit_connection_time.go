package simulator

import "context"

var CmdMobileAccessLimitConnectionTime = &CommandSpec{
	[]Token{TMobile, TAccess, TLimit, TConnection, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMobileAccessLimitConnectionTime = &CommandSpec{
	[]Token{TNo, TMobile, TAccess, TLimit, TConnection, TTime}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

