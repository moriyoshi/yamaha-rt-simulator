package simulator

import "context"

var CmdIpv6InterfaceVrrpShutdownTrigger = &CommandSpec{
	[]Token{TIpv6, TL2Interface, TVrrp, TShutdown, TTrigger}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6InterfaceVrrpShutdownTrigger = &CommandSpec{
	[]Token{TNo, TIpv6, TL2Interface, TVrrp, TShutdown, TTrigger}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
