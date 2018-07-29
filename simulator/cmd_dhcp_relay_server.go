package simulator

import "context"

var CmdDhcpRelayServer = &CommandSpec{
	[]Token{TDhcp, TRelay, TServer}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDhcpRelayServer = &CommandSpec{
	[]Token{TNo, TDhcp, TRelay, TServer}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

