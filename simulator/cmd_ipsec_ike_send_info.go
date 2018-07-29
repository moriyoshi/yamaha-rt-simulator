package simulator

import "context"

var CmdIpsecIkeSendInfo = &CommandSpec{
	[]Token{TIpsec, TIke, TSend, TInfo}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeSendInfo = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TSend, TInfo}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

