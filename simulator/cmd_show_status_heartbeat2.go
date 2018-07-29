package simulator

import "context"

var CmdShowStatusHeartbeat2 = &CommandSpec{
	[]Token{TShow, TStatus, THeartbeat2}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowStatusHeartbeat2Id = &CommandSpec{
	[]Token{TShow, TStatus, THeartbeat2, TId}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowStatusHeartbeat2Name = &CommandSpec{
	[]Token{TShow, TStatus, THeartbeat2, TName}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

