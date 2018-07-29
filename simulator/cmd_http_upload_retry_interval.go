package simulator

import "context"

var CmdHttpUploadRetryInterval = &CommandSpec{
	[]Token{THttp, TUpload, TRetry, TInterval}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHttpUploadRetryInterval = &CommandSpec{
	[]Token{TNo, THttp, TUpload, TRetry, TInterval}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

