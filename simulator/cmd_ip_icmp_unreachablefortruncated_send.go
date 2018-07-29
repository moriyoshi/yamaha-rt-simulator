package simulator

import "context"

var CmdIpIcmpUnreachableForTruncatedSend = &CommandSpec{
	[]Token{TIp, TIcmp, TUnreachableForTruncated, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpIcmpUnreachableForTruncatedSend = &CommandSpec{
	[]Token{TNo, TIp, TIcmp, TUnreachableForTruncated, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

