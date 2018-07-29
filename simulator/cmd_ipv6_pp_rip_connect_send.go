package simulator

import "context"

var CmdIpv6PpRipConnectSend = &CommandSpec{
	[]Token{TIpv6, TPp, TRip, TConnect, TSend}, 5,
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

var CmdNoIpv6PpRipConnectSend = &CommandSpec{
	[]Token{TNo, TIpv6, TPp, TRip, TConnect, TSend}, 6,
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
