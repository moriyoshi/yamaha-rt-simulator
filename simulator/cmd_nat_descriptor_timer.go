package simulator

import "context"

var CmdNatDescriptorTimer = &CommandSpec{
	[]Token{TNat, TDescriptor, TTimer}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNatDescriptorTimer = &CommandSpec{
	[]Token{TNo, TNat, TDescriptor, TTimer}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
