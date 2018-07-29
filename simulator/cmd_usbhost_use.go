package simulator

import "context"

var CmdUsbhostUse = &CommandSpec{
	[]Token{TUsbhost, TUse}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoUsbhostUse = &CommandSpec{
	[]Token{TNo, TUsbhost, TUse}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

