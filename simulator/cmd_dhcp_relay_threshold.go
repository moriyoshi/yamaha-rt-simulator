package simulator

import "context"

var CmdDhcpRelayThreshold = &CommandSpec{
	[]Token{TDhcp, TRelay, TThreshold}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDhcpRelayThreshold = &CommandSpec{
	[]Token{TNo, TDhcp, TRelay, TThreshold}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

