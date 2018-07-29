package simulator

import "context"

var CmdIpPpRipConnectInterval = &CommandSpec{
	[]Token{TIp, TPp, TRip, TConnect, TInterval}, 5,
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

var CmdNoIpPpRipConnectInterval = &CommandSpec{
	[]Token{TNo, TIp, TPp, TRip, TConnect, TInterval}, 6,
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
