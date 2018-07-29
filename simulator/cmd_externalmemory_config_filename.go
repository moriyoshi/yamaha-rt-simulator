package simulator

import "context"

var CmdExternalMemoryConfigFilename = &CommandSpec{
	[]Token{TExternalMemory, TConfig, TFilename}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoExternalMemoryConfigFilename = &CommandSpec{
	[]Token{TNo, TExternalMemory, TConfig, TFilename}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
