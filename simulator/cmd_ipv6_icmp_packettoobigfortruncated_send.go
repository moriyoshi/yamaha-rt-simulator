package simulator

import "context"

var CmdIpv6IcmpPacketTooBigForTruncatedSend = &CommandSpec{
	[]Token{TIpv6, TIcmp, TPacketTooBigForTruncated, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6IcmpPacketTooBigForTruncatedSend = &CommandSpec{
	[]Token{TNo, TIpv6, TIcmp, TPacketTooBigForTruncated, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

