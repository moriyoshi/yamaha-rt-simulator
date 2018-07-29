package simulator

import "context"

var CmdProviderAutoConnectForcedDisable = &CommandSpec{
	[]Token{TProvider, TAuto, TConnect, TForced, TDisable}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoProviderAutoConnectForcedDisable = &CommandSpec{
	[]Token{TNo, TProvider, TAuto, TConnect, TForced, TDisable}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

