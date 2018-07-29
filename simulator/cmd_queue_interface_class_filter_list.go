package simulator

import "context"

var CmdQueueInterfaceClassFilterList = &CommandSpec{
	[]Token{TQueue, TL2Interface, TClass, TFilter, TList}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdQueuePpClassFilterList = &CommandSpec{
	[]Token{TQueue, TPp, TClass, TFilter, TList}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdQueueTunnelClassFilterList = &CommandSpec{
	[]Token{TQueue, TTunnel, TClass, TFilter, TList}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQueueInterfaceClassFilterList = &CommandSpec{
	[]Token{TNo, TQueue, TL2Interface, TClass, TFilter, TList}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQueuePpClassFilterList = &CommandSpec{
	[]Token{TNo, TQueue, TPp, TClass, TFilter, TList}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoQueueTunnelClassFilterList = &CommandSpec{
	[]Token{TNo, TQueue, TTunnel, TClass, TFilter, TList}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

