package simulator

import "context"

var CmdSnmpIfindexSwitchStaticIndex = &CommandSpec{
	[]Token{TSnmp, TIfindex, TSwitch, TStatic, TIndex}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpIfindexSwitchStaticIndex = &CommandSpec{
	[]Token{TNo, TSnmp, TIfindex, TSwitch, TStatic, TIndex}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

