package simulator

import "context"

var CmdShowStatusDhcpc = &CommandSpec{
	[]Token{TShow, TStatus, TDhcpc}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

