package simulator

import "context"

var CmdNatDescriptorMasqueradePortRange = &CommandSpec{
	[]Token{TNat, TDescriptor, TMasquerade, TPort, TRange}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNatDescriptorMasqueradePortRange = &CommandSpec{
	[]Token{TNo, TNat, TDescriptor, TMasquerade, TPort, TRange}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

