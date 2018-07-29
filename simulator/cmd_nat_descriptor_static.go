package simulator

import "context"

var CmdNatDescriptorStatic = &CommandSpec{
	[]Token{TNat, TDescriptor, TStatic}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNatDescriptorStatic = &CommandSpec{
	[]Token{TNo, TNat, TDescriptor, TStatic}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
