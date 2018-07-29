package simulator

import "context"

var CmdTerminateLuaFile = &CommandSpec{
	[]Token{TTerminate, TLua, TFile}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
