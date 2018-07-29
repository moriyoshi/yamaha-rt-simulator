package simulator

import "context"

var CmdLuaUse = &CommandSpec{
	[]Token{TLua, TUse}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoLuaUse = &CommandSpec{
	[]Token{TNo, TLua, TUse}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

