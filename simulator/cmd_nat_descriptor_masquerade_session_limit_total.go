package simulator

import "context"

var CmdNatDescriptorMasqueradeSessionLimitTotal = &CommandSpec{
	[]Token{TNat, TDescriptor, TMasquerade, TSession, TLimit, TTotal}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNatDescriptorMasqueradeSessionLimitTotal = &CommandSpec{
	[]Token{TNo, TNat, TDescriptor, TMasquerade, TSession, TLimit, TTotal}, 7,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

