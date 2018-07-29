package simulator

import "context"

var CmdClearMobileAccessLimitation = &CommandSpec{
	[]Token{TClear, TMobile, TAccess, TLimitation}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
