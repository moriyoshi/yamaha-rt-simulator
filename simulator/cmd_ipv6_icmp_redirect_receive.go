package simulator

import "context"

var CmdIpv6IcmpRedirectReceive = &CommandSpec{
	[]Token{TIpv6, TIcmp, TRedirect, TReceive}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6IcmpRedirectReceive = &CommandSpec{
	[]Token{TNo, TIpv6, TIcmp, TRedirect, TReceive}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

