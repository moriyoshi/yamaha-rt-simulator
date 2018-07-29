package simulator

import "context"

var CmdDhcpClientClientIdentifier = &CommandSpec{
	[]Token{TDhcp, TClient, TClientIdentifier, &EitherToken{TPool, TPp}}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDhcpClientClientIdentifier = &CommandSpec{
	[]Token{TNo, TDhcp, TClient, TClientIdentifier, &EitherToken{TPool, TPp}}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
