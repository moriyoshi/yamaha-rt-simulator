package simulator

import "context"

var CmdIpInterfaceRipAuthType = &CommandSpec{
	[]Token{TIp, TL2Interface, TRip, TAuth, TType}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpRipAuthType = &CommandSpec{
	[]Token{TIp, TPp, TRip, TAuth, TType}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.PP == 0 {
			return InvalidCommandName
		}
		return nil
	},
}

var CmdIpTunnelRipAuthType = &CommandSpec{
	[]Token{TIp, TTunnel, TRip, TAuth, TType}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.Tunnel == 0 {
			return InvalidCommandName
		}
		return nil
	},
}

var CmdNoIpInterfaceRipAuthType = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TRip, TAuth, TType}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpRipAuthType = &CommandSpec{
	[]Token{TNo, TIp, TPp, TRip, TAuth, TType}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.PP == 0 {
			return InvalidCommandName
		}
		return nil
	},
}

var CmdNoIpTunnelRipAuthType = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TRip, TAuth, TType}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		if sess.Tunnel == 0 {
			return InvalidCommandName
		}
		return nil
	},
}
