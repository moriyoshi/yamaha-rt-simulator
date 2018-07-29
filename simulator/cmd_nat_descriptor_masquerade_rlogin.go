package simulator

import "context"

var CmdNatDescriptorMasqueradeRlogin = &CommandSpec{
	[]Token{TNat, TDescriptor, TMasquerade, TRlogin}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNatDescriptorMasqueradeRlogin = &CommandSpec{
	[]Token{TNo, TNat, TDescriptor, TMasquerade, TRlogin}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

