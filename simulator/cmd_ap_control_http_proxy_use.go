package simulator

import "context"

var CmdApControlHttpProxyUse = &CommandSpec{
	[]Token{TAp, TControl, THttp, TProxy, TUse}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoApControlHttpProxyUse = &CommandSpec{
	[]Token{TNo, TAp, TControl, THttp, TProxy, TUse}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

