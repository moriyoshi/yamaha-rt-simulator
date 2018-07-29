package simulator

import "context"

var CmdMobileAccessLimitLength = &CommandSpec{
	[]Token{TMobile, TAccess, TLimit, TLength}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoMobileAccessLimitLength = &CommandSpec{
	[]Token{TNo, TMobile, TAccess, TLimit, TLength}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

