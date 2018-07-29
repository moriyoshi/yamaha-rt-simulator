package simulator

import "context"

var CmdIpIcmpTimestampReplySend = &CommandSpec{
	[]Token{TIp, TIcmp, TTimestampReply, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpIcmpTimestampReplySend = &CommandSpec{
	[]Token{TNo, TIp, TIcmp, TTimestampReply, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

