package simulator

import "context"

var CmdShowStatusOspf = &CommandSpec{
	[]Token{TShow, TStatus, TOspf}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

