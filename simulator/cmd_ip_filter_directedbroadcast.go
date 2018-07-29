package simulator

import "context"

var CmdIpFilterDirectedBroadcast = &CommandSpec{
	[]Token{TIp, TFilter, TDirectedBroadcast}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpFilterDirectedBroadcast = &CommandSpec{
	[]Token{TNo, TIp, TFilter, TDirectedBroadcast}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
