package simulator

import "context"

var CmdNatDescriptorMasqueradeUnconvertiblePort = &CommandSpec{
	[]Token{TNat, TDescriptor, TMasquerade, TUnconvertible, TPort}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNatDescriptorMasqueradeUnconvertiblePort = &CommandSpec{
	[]Token{TNo, TNat, TDescriptor, TMasquerade, TUnconvertible, TPort}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
