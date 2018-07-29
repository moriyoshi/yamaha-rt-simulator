package simulator

import "context"

var CmdHttpdListen = &CommandSpec{
	[]Token{THttpd, TListen}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHttpdListen = &CommandSpec{
	[]Token{TNo, THttpd, TListen}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

