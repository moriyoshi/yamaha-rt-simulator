package simulator

import "context"

var CmdIpPpRipDisconnectInterval = &CommandSpec{
	[]Token{TIp, TPp, TRip, TDisconnect, TInterval}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.PP == 0 {
			return InvalidCommandName
		}
		return nil
	},
}

var CmdNoIpPpRipDisconnectInterval = &CommandSpec{
	[]Token{TNo, TIp, TPp, TRip, TDisconnect, TInterval}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.PP == 0 {
			return InvalidCommandName
		}
		return nil
	},
}
