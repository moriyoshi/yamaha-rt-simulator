package simulator

import "context"

var CmdOperationUsbDownloadPermit = &CommandSpec{
	[]Token{TOperation, TUsbDownload, TPermit}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoOperationUsbDownloadPermit = &CommandSpec{
	[]Token{TNo, TOperation, TUsbDownload, TPermit}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

