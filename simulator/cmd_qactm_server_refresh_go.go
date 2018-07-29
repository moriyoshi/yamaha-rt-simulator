package simulator

import "context"

var CmdQacTmServerRefreshGo = &CommandSpec{
	[]Token{TQacTm, TServer, TRefresh, TGo}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

