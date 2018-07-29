package simulator

import "context"

var CmdIpIcmpParameterProblemSend = &CommandSpec{
	[]Token{TIp, TIcmp, TParameterProblem, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpIcmpParameterProblemSend = &CommandSpec{
	[]Token{TNo, TIp, TIcmp, TParameterProblem, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

