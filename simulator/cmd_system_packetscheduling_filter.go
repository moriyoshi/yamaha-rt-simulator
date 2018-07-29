package simulator

import "context"

var CmdSystemPacketSchedulingFilter = &CommandSpec{
	[]Token{TSystem, TPacketScheduling, TFilter}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSystemPacketSchedulingFilter = &CommandSpec{
	[]Token{TNo, TSystem, TPacketScheduling, TFilter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
