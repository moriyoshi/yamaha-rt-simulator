package simulator

import "context"

var CmdQueueInterfaceDefaultClass = &CommandSpec{
	[]Token{TQueue, TL2Interface, TDefault, TClass}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdQueuePpDefaultClass = &CommandSpec{
	[]Token{TQueue, TPp, TDefault, TClass}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQueueInterfaceDefaultClass = &CommandSpec{
	[]Token{TNo, TQueue, TL2Interface, TDefault, TClass}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQueuePpDefaultClass = &CommandSpec{
	[]Token{TNo, TQueue, TPp, TDefault, TClass}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

