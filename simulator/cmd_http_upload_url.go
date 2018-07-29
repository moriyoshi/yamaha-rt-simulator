package simulator

import "context"

var CmdHttpUploadUrl = &CommandSpec{
	[]Token{THttp, TUpload, TUrl}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHttpUploadUrl = &CommandSpec{
	[]Token{TNo, THttp, TUpload, TUrl}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

