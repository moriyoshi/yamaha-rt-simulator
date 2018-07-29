package simulator

import "context"

var CmdShowStatusNgn = &CommandSpec{
	[]Token{TShow, TStatus, TNgn}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

