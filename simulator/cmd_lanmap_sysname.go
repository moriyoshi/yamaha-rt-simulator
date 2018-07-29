package simulator

import "context"

var CmdLanMapSysname = &CommandSpec{
	[]Token{TLanMap, TSysname}, 2,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoLanMapSysname = &CommandSpec{
	[]Token{TNo, TLanMap, TSysname}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

