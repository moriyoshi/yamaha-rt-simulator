package simulator

import "context"

var CmdNatDescriptorSip = &CommandSpec{
	[]Token{TNat, TDescriptor, TSip}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNatDescriptorSip = &CommandSpec{
	[]Token{TNo, TNat, TDescriptor, TSip}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

