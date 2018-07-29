package simulator

import "context"

var CmdNetvolanteDnsUse = &CommandSpec{
	[]Token{TNetvolanteDns, TUse}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNetvolanteDnsUsePp = &CommandSpec{
	[]Token{TNetvolanteDns, TUse, TPp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNetvolanteDnsUse = &CommandSpec{
	[]Token{TNo, TNetvolanteDns, TUse}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNetvolanteDnsUsePp = &CommandSpec{
	[]Token{TNo, TNetvolanteDns, TUse, TPp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
