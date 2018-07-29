package simulator

import "context"

var CmdExternalMemoryBootPermit = &CommandSpec{
	[]Token{TExternalMemory, TBoot, TPermit}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoExternalMemoryBootPermit = &CommandSpec{
	[]Token{TNo, TExternalMemory, TBoot, TPermit}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

