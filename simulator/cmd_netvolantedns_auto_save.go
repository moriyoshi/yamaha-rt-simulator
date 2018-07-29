package simulator

import "context"

var CmdNetvolanteDnsAutoSave = &CommandSpec{
	[]Token{TNetvolanteDns, TAuto, TSave}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNetvolanteDnsAutoSave = &CommandSpec{
	[]Token{TNo, TNetvolanteDns, TAuto, TSave}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

