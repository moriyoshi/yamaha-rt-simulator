package simulator

import "context"

var CmdClearNatDescriptorDynamic = &CommandSpec{
	[]Token{TClear, TNat, TDescriptor, TDynamic}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

