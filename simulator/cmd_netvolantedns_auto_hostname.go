package simulator

import "context"

var CmdNetvolanteDnsAutoHostname = &CommandSpec{
	[]Token{TNetvolanteDns, TAuto, THostname}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNetvolanteDnsAutoHostnamePp = &CommandSpec{
	[]Token{TNetvolanteDns, TAuto, THostname, TPp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNetvolanteDnsAutoHostname = &CommandSpec{
	[]Token{TNo, TNetvolanteDns, TAuto, THostname}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNetvolanteDnsAutoHostnamePp = &CommandSpec{
	[]Token{TNo, TNetvolanteDns, TAuto, THostname, TPp}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

