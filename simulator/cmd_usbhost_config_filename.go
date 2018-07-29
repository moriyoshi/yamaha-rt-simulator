package simulator

import "context"

var CmdUsbhostConfigFilename = &CommandSpec{
	[]Token{TUsbhost, TConfig, TFilename}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoUsbhostConfigFilename = &CommandSpec{
	[]Token{TNo, TUsbhost, TConfig, TFilename}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

