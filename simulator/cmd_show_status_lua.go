package simulator

import "context"

var CmdShowStatusLua = &CommandSpec{
	[]Token{TShow, TStatus, TLua}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
