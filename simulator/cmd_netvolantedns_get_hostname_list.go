package simulator

import "context"

var CmdNetvolanteDnsGetHostnameList = &CommandSpec{
	[]Token{TNetvolanteDns, TGet, THostname, TList}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNetvolanteDnsGetHostnameListPp = &CommandSpec{
	[]Token{TNetvolanteDns, TGet, THostname, TList, TPp}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
