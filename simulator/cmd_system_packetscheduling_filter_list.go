package simulator

import "context"

var CmdSystemPacketSchedulingFilterList = &CommandSpec{
	[]Token{TSystem, TPacketScheduling, TFilter, TList}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSystemPacketSchedulingFilterList = &CommandSpec{
	[]Token{TNo, TSystem, TPacketScheduling, TFilter, TList}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

