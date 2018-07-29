package simulator

import "context"

var CmdIpInterfaceRipForceToAdvertise = &CommandSpec{
	[]Token{TIp, TL2Interface, TRip, TForceToAdvertise}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpRipForceToAdvertise = &CommandSpec{
	[]Token{TIp, TPp, TRip, TForceToAdvertise}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelRipForceToAdvertise = &CommandSpec{
	[]Token{TIp, TTunnel, TRip, TForceToAdvertise}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceRipForceToAdvertise = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TRip, TForceToAdvertise}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpRipForceToAdvertise = &CommandSpec{
	[]Token{TNo, TIp, TPp, TRip, TForceToAdvertise}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelRipForceToAdvertise = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TRip, TForceToAdvertise}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

