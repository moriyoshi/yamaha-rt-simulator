package simulator

import "context"

var CmdNatDescriptorAddressInner = &CommandSpec{
	[]Token{TNat, TDescriptor, TAddress, TInner}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNatDescriptorAddressInner = &CommandSpec{
	[]Token{TNo, TNat, TDescriptor, TAddress, TInner}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

