package simulator

import "context"

var CmdIpRoutingProcess = &CommandSpec{
	[]Token{TIp, TRouting, TProcess}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpRoutingProcess = &CommandSpec{
	[]Token{TNo, TIp, TRouting, TProcess}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

