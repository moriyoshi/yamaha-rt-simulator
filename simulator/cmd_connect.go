package simulator

import "context"

var CmdConnect = &CommandSpec{
	[]Token{TConnect, &EitherToken{TPp, TTunnel}}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
