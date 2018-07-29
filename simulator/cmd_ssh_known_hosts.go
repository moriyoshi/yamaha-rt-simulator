package simulator

import "context"

var CmdSshKnownHosts = &CommandSpec{
	[]Token{TSsh, TKnown, THosts}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSshKnownHosts = &CommandSpec{
	[]Token{TNo, TSsh, TKnown, THosts}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

