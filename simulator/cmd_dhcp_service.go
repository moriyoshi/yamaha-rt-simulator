package simulator

import "context"

var CmdDhcpService = &CommandSpec{
	[]Token{TDhcp, TService}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDhcpServiceType = &CommandSpec{
	[]Token{TNo, TDhcp, TService}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

