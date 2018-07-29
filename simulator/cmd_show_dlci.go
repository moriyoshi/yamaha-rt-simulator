package simulator

import "context"

var CmdShowDlci = &CommandSpec{
	[]Token{TShow, TDlci}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

