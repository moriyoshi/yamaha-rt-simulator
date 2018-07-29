package simulator

import "context"

var CmdIpInterfaceIntrusionDetectionNoticeInterval = &CommandSpec{
	[]Token{TIp, TL2Interface, TIntrusion, TDetection, TNoticeInterval}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpPpIntrusionDetectionNoticeInterval = &CommandSpec{
	[]Token{TIp, TPp, TIntrusion, TDetection, TNoticeInterval}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdIpTunnelIntrusionDetectionNoticeInterval = &CommandSpec{
	[]Token{TIp, TTunnel, TIntrusion, TDetection, TNoticeInterval}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpInterfaceIntrusionDetectionNotice = &CommandSpec{
	[]Token{TNo, TIp, TL2Interface, TIntrusion, TDetection, TNoticeInterval}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpPpIntrusionDetectionNotice = &CommandSpec{
	[]Token{TNo, TIp, TPp, TIntrusion, TDetection, TNoticeInterval}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoIpTunnelIntrusionDetectionNotice = &CommandSpec{
	[]Token{TNo, TIp, TTunnel, TIntrusion, TDetection, TNoticeInterval}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
