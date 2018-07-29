package simulator

import "context"

var CmdIpv6IcmpPacketTooBigSend = &CommandSpec{
	[]Token{TIpv6, TIcmp, TPacketTooBig, TSend}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpv6IcmpPacketTooBigSend = &CommandSpec{
	[]Token{TNo, TIpv6, TIcmp, TPacketTooBig, TSend}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

