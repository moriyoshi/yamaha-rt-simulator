package simulator

import "context"

var CmdIpv6IcmpRedirectSend = &CommandSpec{
	[]Token{TIpv6, TIcmp, TRedirect, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6IcmpRedirectSend = &CommandSpec{
	[]Token{TNo, TIpv6, TIcmp, TRedirect, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

