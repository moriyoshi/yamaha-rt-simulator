package simulator

import "context"

var CmdNtpdate = &CommandSpec{
	[]Token{TNtpdate}, 1,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

