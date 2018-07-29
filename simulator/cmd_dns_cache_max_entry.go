package simulator

import "context"

var CmdDnsCacheMaxEntry = &CommandSpec{
	[]Token{TDns, TCache, TMax, TEntry}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDnsCacheMaxEntry = &CommandSpec{
	[]Token{TNo, TDns, TCache, TMax, TEntry}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

