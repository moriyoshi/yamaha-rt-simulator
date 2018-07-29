package simulator

import "context"

var CmdSshEncryptAlgorithm = &CommandSpec{
	[]Token{TSsh, TEncrypt, TAlgorithm}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSshEncryptAlgorithm = &CommandSpec{
	[]Token{TNo, TSsh, TEncrypt, TAlgorithm}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

