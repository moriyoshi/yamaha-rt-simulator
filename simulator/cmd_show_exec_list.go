package simulator

import "context"

var CmdShowExecList = &CommandSpec{
	[]Token{TShow, TExec, TList}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
