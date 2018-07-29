package simulator

import "context"

var CmdSystemPacketScheduling = &CommandSpec{
	[]Token{TSystem, TPacketScheduling}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSystemPacketScheduling = &CommandSpec{
	[]Token{TNo, TSystem, TPacketScheduling}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
