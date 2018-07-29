package simulator

import "context"

var CmdNatDescriptorBackwardCompatibility = &CommandSpec{
	[]Token{TNat, TDescriptor, TBackwardCompatibility}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNatDescriptorBackwardCompatibility = &CommandSpec{
	[]Token{TNo, TNat, TDescriptor, TBackwardCompatibility}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

