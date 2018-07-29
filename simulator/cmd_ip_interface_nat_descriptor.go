package simulator

import "context"

var CmdIpInterfaceNatDescriptor = &CommandSpec{
	[]Token{TIp, TL2Interface, TNat, TDescriptor}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpNatDescriptor = &CommandSpec{
	[]Token{TIp, TPp, TNat, TDescriptor}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelNatDescriptor = &CommandSpec{
	[]Token{TIp, TTunnel, TNat, TDescriptor}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceNatDescriptor = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TNat, TDescriptor}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpNatDescriptor = &CommandSpec{
	[]Token{TNo, TIp, TPp, TNat, TDescriptor}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelNatDescriptor = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TNat, TDescriptor}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

