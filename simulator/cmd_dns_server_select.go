package simulator

import "context"

var CmdDnsServerSelect = &CommandSpec{
	[]Token{TDns, TServer, TSelect}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDnsServerSelect = &CommandSpec{
	[]Token{TNo, TDns, TServer, TSelect}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
