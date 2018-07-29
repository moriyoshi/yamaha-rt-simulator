package simulator

import "context"

var CmdDnsServiceAaaaFilter = &CommandSpec{
	[]Token{TDns, TService, TAaaa, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoDnsServiceAaaaFilter = &CommandSpec{
	[]Token{TNo, TDns, TService, TAaaa, TFilter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

