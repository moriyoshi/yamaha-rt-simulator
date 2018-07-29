package simulator

import "context"

var CmdFrCompressionUse = &CommandSpec{
	[]Token{TFr, TCompression, TUse}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoFrCompressionUse = &CommandSpec{
	[]Token{TNo, TFr, TCompression, TUse}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

