package simulator

import "context"

var CmdProviderFilterRouting = &CommandSpec{
	[]Token{TProvider, TFilter, TRouting}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoProviderFilterRouting = &CommandSpec{
	[]Token{TNo, TProvider, TFilter, TRouting}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

