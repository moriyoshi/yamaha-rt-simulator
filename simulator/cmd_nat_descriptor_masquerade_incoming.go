package simulator

import "context"

var CmdNatDescriptorMasqueradeIncoming = &CommandSpec{
	[]Token{TNat, TDescriptor, TMasquerade, TIncoming}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNatDescriptorMasqueradeIncoming = &CommandSpec{
	[]Token{TNo, TNat, TDescriptor, TMasquerade, TIncoming}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

