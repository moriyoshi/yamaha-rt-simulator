package simulator

import "context"

var CmdClearMobileAccessLimitationPp = &CommandSpec{
	[]Token{TClear, TMobile, TAccess, TLimitation, TPp}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
