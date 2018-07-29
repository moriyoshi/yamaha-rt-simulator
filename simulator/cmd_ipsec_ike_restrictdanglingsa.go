package simulator

import "context"

var CmdIpsecIkeRestrictDanglingSa = &CommandSpec{
	[]Token{TIpsec, TIke, TRestrictDanglingSa}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpsecIkeRestrictDanglingSa = &CommandSpec{
	[]Token{TNo, TIpsec, TIke, TRestrictDanglingSa}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

