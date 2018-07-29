package simulator

import "context"

var CmdIpInterfaceDhcpLeaseTime = &CommandSpec{
	[]Token{TIp, TL2Interface, TDhcp, TLease, TTime}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceDhcpLeaseTime = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TDhcp, TLease, TTime}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

