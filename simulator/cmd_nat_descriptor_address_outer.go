package simulator

import "context"

var CmdNatDescriptorAddressOuter = &CommandSpec{
	[]Token{TNat, TDescriptor, TAddress, TOuter}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoNatDescriptorAddressOuter = &CommandSpec{
	[]Token{TNo, TNat, TDescriptor, TAddress, TOuter}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

