package simulator

import "context"

var CmdLuac = &CommandSpec{
	[]Token{TLuac}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

