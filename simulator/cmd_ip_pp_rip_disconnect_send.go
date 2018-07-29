package simulator

import "context"

var CmdIpPpRipDisconnectSend = &CommandSpec{
	[]Token{TIp, TPp, TRip, TDisconnect, TSend}, 5,
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

var CmdNoIpPpRipDisconnect = &CommandSpec{
	[]Token{TNo, TIp, TPp, TRip, TDisconnect, TSend}, 6,
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
