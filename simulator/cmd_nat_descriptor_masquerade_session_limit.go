package simulator

import "context"

var CmdNatDescriptorMasqueradeSessionLimit = &CommandSpec{
	[]Token{TNat, TDescriptor, TMasquerade, TSession, TLimit}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNatDescriptorMasqueradeSessionLimit = &CommandSpec{
	[]Token{TNo, TNat, TDescriptor, TMasquerade, TSession, TLimit}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

