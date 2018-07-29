package simulator

import "context"

var CmdShowStatusIpIgmp = &CommandSpec{
	[]Token{TShow, TStatus, TIp, TIgmp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

