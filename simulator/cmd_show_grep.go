package simulator

import "context"

var CmdShowGrep = &CommandSpec{
	[]Token{TShow, TGrep}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
