package simulator

import "context"

var CmdDnsServerDhcp = &CommandSpec{
	[]Token{TDns, TServer, TDhcp}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDnsServerDhcp = &CommandSpec{
	[]Token{TNo, TDns, TServer, TDhcp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

