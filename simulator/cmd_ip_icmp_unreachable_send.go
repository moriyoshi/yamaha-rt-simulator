package simulator

import "context"

var CmdIpIcmpUnreachableSend = &CommandSpec{
	[]Token{TIp, TIcmp, TUnreachable, TSend, TSomething /* send */, TRebound}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpIcmpUnreachableSend = &CommandSpec{
	[]Token{TNo, TIp, TIcmp, TUnreachable, TSend, TSomething /* send */, TRebound}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

