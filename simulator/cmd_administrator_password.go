package simulator

import "context"

var CmdAdministratorPassword = &CommandSpec{
	[]Token{TAdministrator, TPassword, TEncrypted}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
