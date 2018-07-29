package simulator

import "context"

var CmdShowNatDescriptorInterfaceAddress = &CommandSpec{
	[]Token{TShow, TNat, TDescriptor, TInterface, TAddress}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowNatDescriptorInterfaceAddressPp = &CommandSpec{
	[]Token{TShow, TNat, TDescriptor, TInterface, TAddress, TPp}, 7,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowNatDescriptorInterfaceAddressTunnel = &CommandSpec{
	[]Token{TShow, TNat, TDescriptor, TInterface, TAddress, TTunnel}, 7,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

