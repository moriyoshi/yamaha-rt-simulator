package simulator

import "context"

var CmdDhcpClientHostname = &CommandSpec{
	[]Token{TDhcp, TClient, THostname, &EitherToken{TPool, TPp}}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDhcpClientHostname = &CommandSpec{
	[]Token{TNo, TDhcp, TClient, THostname, &EitherToken{TPool, TPp}}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
