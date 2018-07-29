package simulator

import "context"

var CmdQueueInterfaceType = &CommandSpec{
	[]Token{TQueue, TL2Interface, TType, TSomething /* type */, TShapingLevelArg}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdQueuePpTypeType = &CommandSpec{
	[]Token{TQueue, TPp, TType}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQueueInterfaceType = &CommandSpec{
	[]Token{TNo, TQueue, TL2Interface, TType}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQueuePpTypeType = &CommandSpec{
	[]Token{TNo, TQueue, TPp, TType}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
