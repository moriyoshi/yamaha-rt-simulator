package simulator

import "context"

var CmdShowStatusIpPimSparse = &CommandSpec{
	[]Token{TShow, TStatus, TIp, TPim, TSparse}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

