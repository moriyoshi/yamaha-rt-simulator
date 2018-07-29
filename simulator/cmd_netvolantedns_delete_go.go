package simulator

import "context"

var CmdNetvolanteDnsDeleteGo = &CommandSpec{
	[]Token{TNetvolanteDns, TDelete, TGo}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNetvolanteDnsDeleteGoPp = &CommandSpec{
	[]Token{TNetvolanteDns, TDelete, TGo, TPp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

