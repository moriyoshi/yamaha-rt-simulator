package simulator

import "context"

var CmdDnsServiceFallback = &CommandSpec{
	[]Token{TDns, TService, TFallback}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDnsServiceFallback = &CommandSpec{
	[]Token{TNo, TDns, TService, TFallback}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

