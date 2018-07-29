package simulator

import "context"

var keysForSwitchControlFunctionSet = &EitherToken{
	TCounterFrameRxType,
	TCounterFrameTxType,
	TEnergySaving,
	TLedBrightness,
	TLoopdetectCount,
	TLoopdetectLinkdown,
	TLoopdetectPortUse,
	TLoopdetectRecoveryTimer,
	TLoopdetectTime,
	TLoopdetectUseControlPacket,
	TMacaddressAging,
	TMacaddressAgingTimer,
	TMirroringDest,
	TMirroringSrcRx,
	TMirroringSrcTx,
	TMirroringUse,
	TPoeClass,
	TPortAutoCrossover,
	TPortBlockingControlPacket,
	TPortBlockingDataPacket,
	TPortFlowControl,
	TPortSpeedDownshift,
	TPortSpeed,
	TPortUse,
	TQosDscpRemarkClass,
	TQosDscpRemarkType,
	TQosPolicingSpeed,
	TQosPolicingUse,
	TQosShapingSpeed,
	TQosShapingUse,
	TQosSpeedUnit,
	TSystemName,
	TVlanAccess,
	TVlanId,
	TVlanMultiple,
	TVlanMultipleUse,
	TVlanPortMode,
	TVlanTrunk,
}

var CmdSwitchControlFunctionSet = &CommandSpec{
	[]Token{TSwitch, TControl, TFunction, TSet, keysForSwitchControlFunctionSet}, 5,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}

var CmdNoSwitchControlFunctionSetFunctionIndex = &CommandSpec{
	[]Token{TNo, TSwitch, TControl, TFunction, TSet, keysForSwitchControlFunctionSet}, 6,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
