package simulator

import "context"

var CmdQueueInterfaceLengthSecondary = &CommandSpec{
	[]Token{TQueue, TL2Interface, TLength, TSecondary}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQueueInterfaceLengthSecondary = &CommandSpec{
	[]Token{TNo, TQueue, TL2Interface, TLength, TSecondary}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

