package simulator

import "context"

var CmdIpPpRipHoldRouting = &CommandSpec{
	[]Token{TIp, TPp, TRip, THold, TRouting}, 5,
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

var CmdNoIpPpRipHoldRouting = &CommandSpec{
	[]Token{TNo, TIp, TPp, TRip, THold, TRouting}, 6,
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
