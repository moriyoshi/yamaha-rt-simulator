package simulator

import "context"

var CmdRemoteSetup = &CommandSpec{
	[]Token{TRemote, TSetup}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
