package simulator

import "context"

var CmdOperationExternalMemoryDownloadPermit = &CommandSpec{
	[]Token{TOperation, TExternalMemory, TDownload, TPermit}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoOperationExternalMemoryDownloadPermit = &CommandSpec{
	[]Token{TNo, TOperation, TExternalMemory, TDownload, TPermit}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

