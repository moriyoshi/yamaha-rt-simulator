package simulator

import "context"

var CmdQacTmClientRefreshGo = &CommandSpec{
	[]Token{TQacTm, TClient, TRefresh, TGo}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

