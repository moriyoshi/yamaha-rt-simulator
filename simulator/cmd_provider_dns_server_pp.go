package simulator

import "context"

var CmdProviderDnsServerPp = &CommandSpec{
	[]Token{TProvider, TDns, TServer, TPp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoProviderDnsServerPp = &CommandSpec{
	[]Token{TNo, TProvider, TDns, TServer, TPp}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

