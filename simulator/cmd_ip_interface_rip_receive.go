package simulator

import "context"

var CmdIpInterfaceRipReceive = &CommandSpec{
	[]Token{TIp, TL2Interface, TRip, TReceive}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpRipReceiveReceive = &CommandSpec{
	[]Token{TIp, TPp, TRip, TReceive}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelRipReceiveReceive = &CommandSpec{
	[]Token{TIp, TTunnel, TRip, TReceive}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceRipReceive = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TRip, TReceive}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpRipReceive = &CommandSpec{
	[]Token{TNo, TIp, TPp, TRip, TReceive}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelRipReceive = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TRip, TReceive}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

