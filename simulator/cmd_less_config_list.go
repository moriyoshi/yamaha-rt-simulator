package simulator

import "context"

var CmdLessConfigList = &CommandSpec{
	[]Token{TLess, TConfig, TList}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
