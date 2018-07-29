package simulator

import "context"

var CmdHttpUploadGo = &CommandSpec{
	[]Token{THttp, TUpload, TGo}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

