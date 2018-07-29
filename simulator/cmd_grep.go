package simulator

import "context"

var CmdGrep = &CommandSpec{
	[]Token{TGrep}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
