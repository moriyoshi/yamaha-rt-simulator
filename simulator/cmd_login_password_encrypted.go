package simulator

import "context"

var CmdLoginPasswordEncrypted = &CommandSpec{
	[]Token{TLogin, TPassword, TEncrypted}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

