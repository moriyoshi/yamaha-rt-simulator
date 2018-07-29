package simulator

import "context"

var CmdLua = &CommandSpec{
	[]Token{TLua}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

