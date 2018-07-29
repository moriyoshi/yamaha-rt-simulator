package simulator

import "context"

var CmdQueueInterfaceClassProperty = &CommandSpec{
	[]Token{TQueue, TL2Interface, TClass, TProperty}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdQueuePpClassProperty = &CommandSpec{
	[]Token{TQueue, TPp, TClass, TProperty}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQueueInterfaceClassProperty = &CommandSpec{
	[]Token{TNo, TQueue, TL2Interface, TClass, TProperty}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQueuePpClassProperty = &CommandSpec{
	[]Token{TNo, TQueue, TPp, TClass, TProperty}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
