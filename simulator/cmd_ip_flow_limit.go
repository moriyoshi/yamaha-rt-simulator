package simulator

import "context"

var CmdIpFlowLimit = &CommandSpec{
	[]Token{TIp, TFlow, TLimit}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpFlowLimit = &CommandSpec{
	[]Token{TNo, TIp, TFlow, TLimit}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

