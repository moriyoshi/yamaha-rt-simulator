package simulator

import "context"

var CmdShowStatusNetvolanteDns = &CommandSpec{
	[]Token{TShow, TStatus, TNetvolanteDns}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowStatusNetvolanteDnsPp = &CommandSpec{
	[]Token{TShow, TStatus, TNetvolanteDns, TPp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
