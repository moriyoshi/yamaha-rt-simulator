package simulator

import "context"

var CmdShowPpConnectTime = &CommandSpec{
	[]Token{TShow, TPp, TConnect, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

