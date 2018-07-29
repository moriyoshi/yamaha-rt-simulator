package simulator

import "context"

var CmdIpv6IcmpParameterProblemSend = &CommandSpec{
	[]Token{TIpv6, TIcmp, TParameterProblem, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6IcmpParameterProblemSend = &CommandSpec{
	[]Token{TNo, TIpv6, TIcmp, TParameterProblem, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

