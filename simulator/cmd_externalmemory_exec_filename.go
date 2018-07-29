package simulator

import "context"

var CmdExternalMemoryExecFilename = &CommandSpec{
	[]Token{TExternalMemory, TExec, TFilename}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoExternalMemoryExecFilename = &CommandSpec{
	[]Token{TNo, TExternalMemory, TExec, TFilename}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
