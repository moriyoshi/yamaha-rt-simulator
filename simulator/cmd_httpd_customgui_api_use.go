package simulator

import "context"

var CmdHttpdCustomGuiApiUse = &CommandSpec{
	[]Token{THttpd, TCustomGui, TApi, TUse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHttpdCustomGuiApiUse = &CommandSpec{
	[]Token{TNo, THttpd, TCustomGui, TApi, TUse}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

