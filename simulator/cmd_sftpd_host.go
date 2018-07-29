package simulator

import "context"

var CmdSftpdHost = &CommandSpec{
	[]Token{TSftpd, THost}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSftpdHost = &CommandSpec{
	[]Token{TNo, TSftpd, THost}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
