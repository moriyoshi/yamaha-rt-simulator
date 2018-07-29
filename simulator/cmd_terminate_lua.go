package simulator

import "context"

var CmdTerminateLua = &CommandSpec{
	[]Token{TTerminate, TLua, TFile}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
