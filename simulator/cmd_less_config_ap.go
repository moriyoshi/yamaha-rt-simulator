package simulator

import "context"

var CmdLessConfigAp = &CommandSpec{
	[]Token{TLess, TConfig, TAp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
