package simulator

import "context"

var CmdLanShutdown = &CommandSpec{
	[]Token{TLan, TShutdown}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoLanShutdown = &CommandSpec{
	[]Token{TNo, TLan, TShutdown}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

