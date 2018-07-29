package simulator

import "context"

var CmdIpInterfaceArpMtuDiscovery = &CommandSpec{
	[]Token{TIp, TL2Interface, TArp, TMtu, TDiscovery}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceArpMtuDiscovery = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TArp, TMtu, TDiscovery}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

