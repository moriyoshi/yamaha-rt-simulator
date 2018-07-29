package simulator

import "context"

var CmdSshdEncryptAlgorithm = &CommandSpec{
	[]Token{TSshd, TEncrypt, TAlgorithm}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSshdEncryptAlgorithm = &CommandSpec{
	[]Token{TNo, TSshd, TEncrypt, TAlgorithm}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

