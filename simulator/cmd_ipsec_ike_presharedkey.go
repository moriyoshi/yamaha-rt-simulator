package simulator

import "context"

var CmdIpsecIkePreSharedKey = &CommandSpec{
	[]Token{TIpsec, TIke, TPreSharedKey}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkePreSharedKey = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TPreSharedKey}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
