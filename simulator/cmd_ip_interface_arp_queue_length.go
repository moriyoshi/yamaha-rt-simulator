package simulator

import "context"

var CmdIpInterfaceArpQueueLength = &CommandSpec{
	[]Token{TIp, TL2Interface, TArp, TQueue, TLength}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceArpQueueLength = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TArp, TQueue, TLength}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

