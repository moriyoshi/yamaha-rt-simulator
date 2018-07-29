package simulator

import "context"

var CmdIpPpRipReceive = &CommandSpec{
	[]Token{TIp, TPp, TRip, TReceive}, 4,
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
