package simulator

import "context"

var CmdIpIcmpTimeExceededSend = &CommandSpec{
	[]Token{TIp, TIcmp, TTimeExceeded, TSend, TSomething /* send */, TRebound}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpIcmpTimeExceededSendSend = &CommandSpec{
	[]Token{TNo, TIp, TIcmp, TTimeExceeded, TSend, TSomething /* send */, TRebound}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

