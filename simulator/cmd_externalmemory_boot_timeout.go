package simulator

import "context"

var CmdExternalMemoryBootTimeout = &CommandSpec{
	[]Token{TExternalMemory, TBoot, TTimeout}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoExternalMemoryBootTimeout = &CommandSpec{
	[]Token{TNo, TExternalMemory, TBoot, TTimeout}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

