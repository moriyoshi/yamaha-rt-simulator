package simulator

import "context"

var CmdIpv6InterfaceDadRetryCount = &CommandSpec{
	[]Token{TIpv6, TL2Interface, TDad, TRetry, TCount}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpv6PpDadRetryCount = &CommandSpec{
	[]Token{TIpv6, TPp, TDad, TRetry, TCount}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6InterfaceDadRetryCount = &CommandSpec{
	[]Token{TNo, TIpv6, TL2Interface, TDad, TRetry, TCount}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6PpDadRetryCount = &CommandSpec{
	[]Token{TNo, TIpv6, TPp, TDad, TRetry, TCount}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

