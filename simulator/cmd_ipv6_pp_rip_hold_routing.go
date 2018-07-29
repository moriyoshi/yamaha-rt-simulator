package simulator

import "context"

var CmdIpv6PpRipHoldRouting = &CommandSpec{
	[]Token{TIpv6, TPp, TRip, THold, TRouting}, 5,
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

var CmdNoIpv6PpRipHoldRouting = &CommandSpec{
	[]Token{TNo, TIpv6, TPp, TRip, THold, TRouting}, 6,
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
