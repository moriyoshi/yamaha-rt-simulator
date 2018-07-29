package simulator

import "context"

var CmdHttpRevisionUpGo = &CommandSpec{
	[]Token{THttp, TRevisionUp, TGo}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

