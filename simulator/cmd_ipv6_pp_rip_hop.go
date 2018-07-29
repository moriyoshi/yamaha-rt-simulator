package simulator

import "context"

var CmdIpv6PpRipHop = &CommandSpec{
	[]Token{TIpv6, TPp, TRip, THop}, 4,
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
