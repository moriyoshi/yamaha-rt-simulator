package simulator

import "context"

var CmdNetvolanteDnsGo = &CommandSpec{
	[]Token{TNetvolanteDns, TGo}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNetvolanteDnsGoPp = &CommandSpec{
	[]Token{TNetvolanteDns, TGo, TPp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
