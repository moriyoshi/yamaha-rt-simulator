package simulator

import "context"

var CmdNatDescriptorMasqueradeTtlHold = &CommandSpec{
	[]Token{TNat, TDescriptor, TMasquerade, TTtl, THold}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNatDescriptorMasqueradeTtlHold = &CommandSpec{
	[]Token{TNo, TNat, TDescriptor, TMasquerade, TTtl, THold}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

