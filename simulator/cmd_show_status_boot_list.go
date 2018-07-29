package simulator

import "context"

var CmdShowStatusBootList = &CommandSpec{
	[]Token{TShow, TStatus, TBoot, TList}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

