package simulator

import "context"

var CmdShowLineMasterclock = &CommandSpec{
	[]Token{TShow, TLine, TMasterclock}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

