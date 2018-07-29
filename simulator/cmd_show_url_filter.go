package simulator

import "context"

var CmdShowUrlFilter = &CommandSpec{
	[]Token{TShow, TUrl, TFilter}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowUrlFilterPp = &CommandSpec{
	[]Token{TShow, TUrl, TFilter, TPp}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdShowUrlFilterTunnel = &CommandSpec{
	[]Token{TShow, TUrl, TFilter, TTunnel}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
