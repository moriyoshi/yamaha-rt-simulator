package simulator

import "context"

var CmdIpPpRipConnectSend = &CommandSpec{
	[]Token{TIp, TPp, TRip, TConnect, TSend}, 5,
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

var CmdNoIpPpRipConnectSend = &CommandSpec{
	[]Token{TNo, TIp, TPp, TRip, TConnect, TSend}, 6,
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
