package simulator

import "context"

var CmdIpInterfaceRipHop = &CommandSpec{
	[]Token{TIp, TL2Interface, TRip, THop}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpRipHopDirectionHop = &CommandSpec{
	[]Token{TIp, TPp, TRip, THop}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelRipHopDirectionHop = &CommandSpec{
	[]Token{TIp, TTunnel, TRip, THop}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceRipHopDirectionHop = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TRip, THop}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpRipHopDirectionHop = &CommandSpec{
	[]Token{TNo, TIp, TPp, TRip, THop}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelRipHopDirectionHop = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TRip, THop}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

