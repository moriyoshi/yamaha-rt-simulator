package simulator

import "context"

var CmdNetvolanteDnsPort = &CommandSpec{
	[]Token{TNetvolanteDns, TPort}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNetvolanteDnsPort = &CommandSpec{
	[]Token{TNo, TNetvolanteDns, TPort}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

