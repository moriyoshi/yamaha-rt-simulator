package simulator

import "context"

var CmdShowStatusQos = &CommandSpec{
	[]Token{TShow, TStatus, TQos}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

