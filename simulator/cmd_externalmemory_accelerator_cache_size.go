package simulator

import "context"

var CmdExternalMemoryAcceleratorCacheSize = &CommandSpec{
	[]Token{TExternalMemory, TAccelerator, TCache, TSize}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoExternalMemoryAcceleratorCacheSize = &CommandSpec{
	[]Token{TNo, TExternalMemory, TAccelerator, TCache, TSize}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

