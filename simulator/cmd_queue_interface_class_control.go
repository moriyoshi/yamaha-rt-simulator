package simulator

import "context"

var CmdQueueInterfaceClassControl = &CommandSpec{
	[]Token{TQueue, TL2Interface, TClass, TControl}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQueueInterfaceClassControl = &CommandSpec{
	[]Token{TNo, TQueue, TL2Interface, TClass, TControl}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

