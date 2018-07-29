package simulator

import "context"

var CmdIpInterfaceMtu = &CommandSpec{
	[]Token{TIp, TL2Interface, TMtu}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpMtu = &CommandSpec{
	[]Token{TIp, TPp, TMtu}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelMtu = &CommandSpec{
	[]Token{TIp, TTunnel, TMtu}, 3,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceMtu = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TMtu}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpMtu = &CommandSpec{
	[]Token{TNo, TIp, TPp, TMtu}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelMtu = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TMtu}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

