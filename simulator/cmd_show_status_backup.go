package simulator

import "context"

var CmdShowStatusBackup = &CommandSpec{
	[]Token{TShow, TStatus, TBackup}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

