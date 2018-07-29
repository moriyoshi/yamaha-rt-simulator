package simulator

import "context"

var CmdSnmpYrswindexSwitchStaticIndex = &CommandSpec{
	[]Token{TSnmp, TYrswindex, TSwitch, TStatic, TIndex}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSnmpYrswindexSwitchStaticIndex = &CommandSpec{
	[]Token{TNo, TSnmp, TYrswindex, TSwitch, TStatic, TIndex}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

