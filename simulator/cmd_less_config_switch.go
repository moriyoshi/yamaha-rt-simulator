package simulator

import "context"

var CmdLessConfigSwitch = &CommandSpec{
	[]Token{TLess, TConfig, TSwitch}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
