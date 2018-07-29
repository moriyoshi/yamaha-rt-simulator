package simulator

import "context"

var CmdProviderInterfaceDnsServer = &CommandSpec{
	[]Token{TProvider, TL2Interface, TDns, TServer}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoProviderInterfaceDnsServer = &CommandSpec{
	[]Token{TNo, TProvider, TL2Interface, TDns, TServer}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

