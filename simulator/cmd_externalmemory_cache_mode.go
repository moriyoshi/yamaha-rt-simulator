package simulator

import "context"

var CmdExternalMemoryCacheMode = &CommandSpec{
	[]Token{TExternalMemory, TCache, TMode}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoExternalMemoryCacheMode = &CommandSpec{
	[]Token{TNo, TExternalMemory, TCache, TMode}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

