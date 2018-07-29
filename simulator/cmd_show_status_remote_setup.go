package simulator

import "context"

var CmdShowStatusRemoteSetup = &CommandSpec{
	[]Token{TShow, TStatus, TRemote, TSetup}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

