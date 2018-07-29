package simulator

import "context"

var CmdShowStatusUsbhost = &CommandSpec{
	[]Token{TShow, TStatus, TUsbhost}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

