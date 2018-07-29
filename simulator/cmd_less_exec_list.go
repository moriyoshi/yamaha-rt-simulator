package simulator

import "context"

var CmdLessExecList = &CommandSpec{
	[]Token{TLess, TExec, TList}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
