package simulator

import "context"

var CmdSystemPacketBuffer = &CommandSpec{
	[]Token{TSystem, TPacketBuffer}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSystemPacketBufferGroupParameterValue = &CommandSpec{
	[]Token{TNo, TSystem, TPacketBuffer}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

