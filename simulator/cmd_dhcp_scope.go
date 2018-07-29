package simulator

import "context"

var CmdDhcpScope = &CommandSpec{
	[]Token{TDhcp, TScope}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDhcpScope = &CommandSpec{
	[]Token{TNo, TDhcp, TScope}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

