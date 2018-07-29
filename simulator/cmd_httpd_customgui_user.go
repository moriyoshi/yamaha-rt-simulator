package simulator

import "context"

var CmdHttpdCustomGuiUser = &CommandSpec{
	[]Token{THttpd, TCustomGui, TUser}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHttpdCustomGuiUser = &CommandSpec{
	[]Token{TNo, THttpd, TCustomGui, TUser}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

