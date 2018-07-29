package simulator

import "context"

var CmdLineMasterclock = &CommandSpec{
	[]Token{TLine, TMasterclock}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoLineMasterclock = &CommandSpec{
	[]Token{TNo, TLine, TMasterclock}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
