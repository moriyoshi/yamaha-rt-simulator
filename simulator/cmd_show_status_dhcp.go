package simulator

import "context"

var CmdShowStatusDhcp = &CommandSpec{
	[]Token{TShow, TStatus, TDhcp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

