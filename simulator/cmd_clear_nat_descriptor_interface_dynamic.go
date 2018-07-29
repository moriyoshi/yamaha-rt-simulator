package simulator

import "context"

var CmdClearNatDescriptorInterfaceDynamic = &CommandSpec{
	[]Token{TClear, TNat, TDescriptor, TInterface, TDynamic, &EitherToken{TPp, TTunnel}}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
