package simulator

import "context"

var CmdHttpRevisionDownPermit = &CommandSpec{
	[]Token{THttp, TRevisionDown, TPermit}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoHttpRevisionDownPermit = &CommandSpec{
	[]Token{TNo, THttp, TRevisionDown, TPermit}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

