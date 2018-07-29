package simulator

import "context"

var CmdNetvolanteDnsServer = &CommandSpec{
	[]Token{TNetvolanteDns, TServer}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNetvolanteDnsServer = &CommandSpec{
	[]Token{TNo, TNetvolanteDns, TServer}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
