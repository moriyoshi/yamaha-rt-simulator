package simulator

import "context"

var CmdDnsService = &CommandSpec{
	[]Token{TDns, TService}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDnsService = &CommandSpec{
	[]Token{TNo, TDns, TService}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

