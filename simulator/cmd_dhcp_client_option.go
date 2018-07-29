package simulator

import "context"

var CmdDhcpClientOption = &CommandSpec{
	[]Token{TDhcp, TClient, TOption, &EitherToken{TPool, TPp}}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDhcpClientOptionInterfaceOptionValue = &CommandSpec{
	[]Token{TNo, TDhcp, TClient, TOption, &EitherToken{TPool, TPp}}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
