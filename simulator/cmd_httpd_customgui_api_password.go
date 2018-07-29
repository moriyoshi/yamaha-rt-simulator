package simulator

import "context"

var CmdHttpdCustomGuiApiPassword = &CommandSpec{
	[]Token{THttpd, TCustomGui, TApi, TPassword}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHttpdCustomGuiApiPassword = &CommandSpec{
	[]Token{TNo, THttpd, TCustomGui, TApi, TPassword}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

