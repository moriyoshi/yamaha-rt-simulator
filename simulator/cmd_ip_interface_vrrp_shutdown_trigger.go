package simulator

import "context"

var CmdIpInterfaceVrrpShutdownTrigger = &CommandSpec{
	[]Token{TIp, TL2Interface, TVrrp, TShutdown, TTrigger}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceVrrpShutdownTrigger = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TVrrp, TShutdown, TTrigger}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
