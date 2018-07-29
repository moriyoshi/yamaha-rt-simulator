package simulator

import "context"

var CmdQueueInterfaceLength = &CommandSpec{
	[]Token{TQueue, TL2Interface, TLength}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdQueuePpLength = &CommandSpec{
	[]Token{TQueue, TPp, TLength}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQueueInterfaceLength = &CommandSpec{
	[]Token{TNo, TQueue, TL2Interface, TLength}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQueuePpLength = &CommandSpec{
	[]Token{TNo, TQueue, TPp, TLength}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

