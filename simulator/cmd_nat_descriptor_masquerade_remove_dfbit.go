package simulator

import "context"

var CmdNatDescriptorMasqueradeRemoveDfBit = &CommandSpec{
	[]Token{TNat, TDescriptor, TMasquerade, TRemove, TDfBit}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNatDescriptorMasqueradeRemoveDfBit = &CommandSpec{
	[]Token{TNo, TNat, TDescriptor, TMasquerade, TRemove, TDfBit}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

