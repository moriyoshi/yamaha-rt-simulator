package simulator

import "context"

var CmdDisconnect = &CommandSpec{
	[]Token{TDisconnect, &EitherToken{TPp, TTunnel}}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
