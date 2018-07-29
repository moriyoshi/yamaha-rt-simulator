package simulator

import "context"

var CmdShowStatusHeartbeat = &CommandSpec{
	[]Token{TShow, TStatus, THeartbeat}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

