package simulator

import "context"

var CmdHttpdCustomGuiUse = &CommandSpec{
	[]Token{THttpd, TCustomGui, TUse}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHttpdCustomGuiUse = &CommandSpec{
	[]Token{TNo, THttpd, TCustomGui, TUse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

