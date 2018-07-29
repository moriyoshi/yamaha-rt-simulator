package simulator

import "context"

var CmdExternalMemoryBatchFilename = &CommandSpec{
	[]Token{TExternalMemory, TBatch, TFilename}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoExternalMemoryBatchFilename = &CommandSpec{
	[]Token{TNo, TExternalMemory, TBatch, TFilename}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

