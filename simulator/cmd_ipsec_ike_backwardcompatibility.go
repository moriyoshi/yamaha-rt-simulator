package simulator

import "context"

var CmdIpsecIkeBackwardCompatibility = &CommandSpec{
	[]Token{TIpsec, TIke, TBackwardCompatibility}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeBackwardCompatibility = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TBackwardCompatibility}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
