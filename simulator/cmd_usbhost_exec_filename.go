package simulator

import "context"

var CmdUsbhostExecFilename = &CommandSpec{
	[]Token{TUsbhost, TExec, TFilename}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoUsbhostExecFilenameFromTo = &CommandSpec{
	[]Token{TNo, TUsbhost, TExec, TFilename}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

