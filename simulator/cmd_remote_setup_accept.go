package simulator

import "context"

var CmdRemoteSetupAccept = &CommandSpec{
	[]Token{TRemote, TSetup, TAccept}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoRemoteSetupAccept = &CommandSpec{
	[]Token{TNo, TRemote, TSetup, TAccept}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
