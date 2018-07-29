package simulator

import "context"

var CmdNatDescriptorType = &CommandSpec{
	[]Token{TNat, TDescriptor, TType}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNatDescriptorType = &CommandSpec{
	[]Token{TNo, TNat, TDescriptor, TType}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

