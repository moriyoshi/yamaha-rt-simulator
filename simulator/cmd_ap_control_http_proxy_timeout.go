package simulator

import "context"

var CmdApControlHttpProxyTimeout = &CommandSpec{
	[]Token{TAp, TControl, THttp, TProxy, TTimeout}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoApControlHttpProxyTimeout = &CommandSpec{
	[]Token{TNo, TAp, TControl, THttp, TProxy, TTimeout}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

