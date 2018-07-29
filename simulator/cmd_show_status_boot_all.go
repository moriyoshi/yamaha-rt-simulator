package simulator

import "context"

var CmdShowStatusBootAll = &CommandSpec{
	[]Token{TShow, TStatus, TBoot, TAll}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

