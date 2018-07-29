package simulator

import "context"

var CmdClearNatDescriptorInterfaceDynamicTunnel = &CommandSpec{
	[]Token{TClear, TNat, TDescriptor, TInterface, TDynamic, TTunnel}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
