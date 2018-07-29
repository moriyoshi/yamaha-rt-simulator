package simulator

import "context"

var CmdShowNatDescriptorInterfaceBind = &CommandSpec{
	[]Token{TShow, TNat, TDescriptor, TInterface, TBind}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowNatDescriptorInterfaceBindPp = &CommandSpec{
	[]Token{TShow, TNat, TDescriptor, TInterface, TBind, TPp}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowNatDescriptorInterfaceBindTunnel = &CommandSpec{
	[]Token{TShow, TNat, TDescriptor, TInterface, TBind, TTunnel}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

