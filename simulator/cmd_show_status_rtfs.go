package simulator

import "context"

var CmdShowStatusRtfs = &CommandSpec{
	[]Token{TShow, TStatus, TRtfs}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

