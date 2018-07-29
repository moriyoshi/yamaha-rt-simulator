package simulator

import "context"

var CmdQueueInterfaceDefaultClassSecondary = &CommandSpec{
	[]Token{TQueue, TL2Interface, TDefault, TClass, TSecondary}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQueueInterfaceDefaultClassSecondary = &CommandSpec{
	[]Token{TNo, TQueue, TL2Interface, TDefault, TClass, TSecondary}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

