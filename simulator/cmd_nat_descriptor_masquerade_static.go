package simulator

import "context"

var CmdNatDescriptorMasqueradeStatic = &CommandSpec{
	[]Token{TNat, TDescriptor, TMasquerade, TStatic}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNatDescriptorMasqueradeStatic = &CommandSpec{
	[]Token{TNo, TNat, TDescriptor, TMasquerade, TStatic}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

