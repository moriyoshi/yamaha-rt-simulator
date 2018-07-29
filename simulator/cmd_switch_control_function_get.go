package simulator

import "context"

var CmdSwitchControlFunctionGet = &CommandSpec{
	[]Token{
		TSwitch, TControl, TFunction, TGet,
		&EitherToken{
			TBootRomVersion,
			TCounterFrameRxType,
			TCounterFrameTxType,
			TEnergySaving,
			TFirmwareRevision,
			TLagType,
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
			TModelName,
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
			TSerialNumber,
			TStatusComboPort,
			TStatusCounterFrameRx,
			TStatusCounterFrameTx,
			TStatusCounterOctetRx,
			TStatusCounterOctetTx,
			TStatusFan,
			TStatusFanRpm,
			TStatusLedMode,
			TStatusLoopdetectPort,
			TStatusLoopdetectRecoveryTimer,
			TStatusMacaddressAddr,
			TStatusMacaddressPort,
			TStatusPoeDetectClass,
			TStatusPoeState,
			TStatusPoeSupplyDetail,
			TStatusPoeSupply,
			TStatusPoeSupplyTotal,
			TStatusPoeTemperature,
			TStatusPortSfpRxPower,
			TStatusPortSpeed,
			TSystemMacaddress,
			TSystemName,
			TSystemUptime,
			TVlanAccess,
			TVlanId,
			TVlanMultiple,
			TVlanMultipleUse,
			TVlanPortMode,
			TVlanTrunk,
		},
	}, 4,
	func(_ context.Context, sess *SimulatorSession, tis []TokenInstance) error {
		if !sess.Enabled {
			return AdministratorUseOnly
		}
		return nil
	},
}
