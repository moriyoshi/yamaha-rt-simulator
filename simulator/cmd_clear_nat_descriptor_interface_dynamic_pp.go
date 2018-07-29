package simulator

import "context"

var CmdClearNatDescriptorInterfaceDynamicPp = &CommandSpec{
	[]Token{TClear, TNat, TDescriptor, TInterface, TDynamic, TPp}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
