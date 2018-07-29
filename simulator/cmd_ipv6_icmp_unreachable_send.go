package simulator

import "context"

var CmdIpv6IcmpUnreachableSend = &CommandSpec{
	[]Token{TIpv6, TIcmp, TUnreachable, TSend, TSomething /* send */, TRebound}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6IcmpUnreachableSend = &CommandSpec{
	[]Token{TNo, TIpv6, TIcmp, TUnreachable, TSend, TSomething /* send */, TRebound}, 7,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

