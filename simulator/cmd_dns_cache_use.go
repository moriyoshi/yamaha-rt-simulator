package simulator

import "context"

var CmdDnsCacheUse = &CommandSpec{
	[]Token{TDns, TCache, TUse}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDnsCacheUse = &CommandSpec{
	[]Token{TNo, TDns, TCache, TUse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

