package simulator

import "context"

var CmdExternalMemoryAutoSearchTime = &CommandSpec{
	[]Token{TExternalMemory, TAutoSearch, TTime}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoExternalMemoryAutoSearchTime = &CommandSpec{
	[]Token{TNo, TExternalMemory, TAutoSearch, TTime}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

