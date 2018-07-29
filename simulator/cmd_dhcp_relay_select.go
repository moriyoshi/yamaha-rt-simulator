package simulator

import "context"

var CmdDhcpRelaySelect = &CommandSpec{
	[]Token{TDhcp, TRelay, TSelect}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDhcpRelaySelect = &CommandSpec{
	[]Token{TNo, TDhcp, TRelay, TSelect}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

