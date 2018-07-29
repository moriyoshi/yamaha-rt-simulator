package simulator

import "context"

var CmdShowStatusExternalMemory = &CommandSpec{
	[]Token{TShow, TStatus, TExternalMemory}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

