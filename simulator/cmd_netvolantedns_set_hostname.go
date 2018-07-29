package simulator

import "context"

var CmdNetvolanteDnsSetHostname = &CommandSpec{
	[]Token{TNetvolanteDns, TSet, THostname}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

