package simulator

import "context"

var CmdShowStatusStatusLed = &CommandSpec{
	[]Token{TShow, TStatus, TStatusLed}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

