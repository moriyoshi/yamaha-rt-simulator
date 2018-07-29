package simulator

import "context"

var CmdProviderIpv6ConnectPp = &CommandSpec{
	[]Token{TProvider, TIpv6, TConnect, TPp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoProviderIpv6ConnectPp = &CommandSpec{
	[]Token{TNo, TProvider, TIpv6, TConnect, TPp}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

