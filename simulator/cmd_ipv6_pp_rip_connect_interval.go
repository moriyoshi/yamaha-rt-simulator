package simulator

import "context"

var CmdIpv6PpRipConnectInterval = &CommandSpec{
	[]Token{TIpv6, TPp, TRip, TConnect, TInterval}, 5,
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

var CmdNoIpv6PpRipConnectInterval = &CommandSpec{
	[]Token{TNo, TIpv6, TPp, TRip, TConnect, TInterval}, 6,
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
