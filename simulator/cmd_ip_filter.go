package simulator

import "context"

var CmdIpFilter = &CommandSpec{
	[]Token{TIp, TFilter}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpFilterFilter_NumPass_Reject = &CommandSpec{
	[]Token{TNo, TIp, TFilter}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

