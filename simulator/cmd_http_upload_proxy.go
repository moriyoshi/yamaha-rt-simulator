package simulator

import "context"

var CmdHttpUploadProxy = &CommandSpec{
	[]Token{THttp, TUpload, TProxy}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHttpUploadProxy = &CommandSpec{
	[]Token{TNo, THttp, TUpload, TProxy}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

