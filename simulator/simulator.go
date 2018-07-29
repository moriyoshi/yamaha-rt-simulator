package simulator

import (
	"context"
	"fmt"
	"net"
	"strings"
	"sync"
	"text/template"
	"time"

	base "github.com/moriyoshi/yamaha-rt-simulator"
	"golang.org/x/text/encoding"
	encoding_index "golang.org/x/text/encoding/ianaindex"
)

type Firmware struct {
	Version   string
	Date      time.Time
	Copyright string
}

type NetInterface interface {
	Name() string
}

type L2Interface interface {
	NetInterface
	HardwareAddr() net.HardwareAddr
}

type PhyInterface struct {
	Name_   string
	HwAddr_ net.HardwareAddr
}

type Logger interface {
	base.DebugLogger
	base.InfoLogger
	base.ErrorLogger
}

func (p *PhyInterface) Name() string {
	return p.Name_
}

func (p *PhyInterface) HardwareAddr() net.HardwareAddr {
	return p.HwAddr_
}

type PortVlanInterface struct {
	Index int
	Phy   *PhyInterface
}

func (p *PortVlanInterface) Name() string {
	return fmt.Sprintf("%s.%d", p.Phy.Name(), p.Index)
}

type PortVlanGroupInterface struct {
	Name_    string
	_Members []*PortVlanInterface
}

func (p *PortVlanGroupInterface) Name() string {
	return p.Name_
}

type TaggedVlanInterface struct {
	Index int
	Phy   *PhyInterface
}

func (p *TaggedVlanInterface) Name() string {
	return fmt.Sprintf("%s/%d", p.Phy.Name(), p.Index)
}

type PPInterface struct {
	Index int
	NetInterface
}

type L2InterfaceToken struct{}

func (*L2InterfaceToken) Candidates(ctx interface{}, im string, candidates []string, candidatesForDisplay []string) ([]string, []string) {
	sim := ctx.(*SimulatorSession)
	im = strings.ToLower(im)
	for _, if_ := range sim.PhyInterfaces {
		name := if_.Name()
		if strings.HasPrefix(name, im) {
			candidates = append(candidates, name)
			candidatesForDisplay = append(candidatesForDisplay, name)
		}
	}
	for _, if_ := range sim.PortVlanInterfaces {
		name := if_.Name()
		if strings.HasPrefix(name, im) {
			candidates = append(candidates, name)
			candidatesForDisplay = append(candidatesForDisplay, name)
		}
	}
	for _, if_ := range sim.PortVlanGroupInterfaces {
		name := if_.Name()
		if strings.HasPrefix(name, im) {
			candidates = append(candidates, name)
			candidatesForDisplay = append(candidatesForDisplay, name)
		}
	}

	for _, phy := range sim.PhyInterfaces {
		minIndex := 0
		maxIndex := 0
		vlanIfs := sim.TaggedVlanInterfaces[phy]
		if vlanIfs == nil {
			continue
		}
		candidateAvailable := false
		for _, if_ := range vlanIfs {
			name := if_.Name()
			if maxIndex < if_.Index {
				maxIndex = if_.Index
			}
			if minIndex == 0 || minIndex > if_.Index {
				minIndex = if_.Index
			}
			if strings.HasPrefix(name, im) {
				candidates = append(candidates, name)
				candidateAvailable = true
			}
		}
		if candidateAvailable {
			candidatesForDisplay = append(candidatesForDisplay, fmt.Sprintf("%s/<%d-%d>", phy.Name(), minIndex, maxIndex))
		}
	}

	return candidates, candidatesForDisplay
}

func (*L2InterfaceToken) Matches(ctx interface{}, im string) (string, bool) {
	im = strings.ToLower(im)
	sim := ctx.(*SimulatorSession)
	for _, if_ := range sim.PhyInterfaces {
		if if_.Name() == im {
			return im, true
		}
	}
	for _, if_ := range sim.PortVlanInterfaces {
		if if_.Name() == im {
			return im, true
		}
	}
	for _, if_ := range sim.PortVlanGroupInterfaces {
		if if_.Name() == im {
			return im, true
		}
	}
	for _, ifs := range sim.TaggedVlanInterfaces {
		for _, if_ := range ifs {
			if if_.Name() == im {
				return im, true
			}
		}
	}
	return "", false
}

type CharsetToken struct{}

func (*CharsetToken) Candidates(ctx interface{}, im string, candidates []string, candidatesForDisplay []string) ([]string, []string) {
	sim := ctx.(*SimulatorSession)
	im = strings.ToLower(im)
	for _, charset_ := range sim.Charsets {
		name := charset_.Name
		if strings.HasPrefix(name, im) {
			candidates = append(candidates, name)
			candidatesForDisplay = append(candidatesForDisplay, name)
		}
	}
	return candidates, candidatesForDisplay
}

func (*CharsetToken) Matches(ctx interface{}, im string) (string, bool) {
	sim := ctx.(*SimulatorSession)
	im = strings.ToLower(im)
	for _, charset_ := range sim.Charsets {
		name := charset_.Name
		if name == im {
			return name, true
		}
	}
	return "", false
}

var (
	T100Rel                        = &SymbolToken{"100rel"}
	T8021Q                         = &SymbolToken{"802.1q"}
	TAaaa                          = &SymbolToken{"aaaa"}
	TAccelerator                   = &SymbolToken{"accelerator"}
	TAccept                        = &SymbolToken{"accept"}
	TAccessPoint                   = &SymbolToken{"access-point"}
	TAccess                        = &SymbolToken{"access"}
	TAccm                          = &SymbolToken{"accm"}
	TAccount                       = &SymbolToken{"account"}
	TAccumulate                    = &SymbolToken{"accumulate"}
	TAcfc                          = &SymbolToken{"acfc"}
	TActive                        = &SymbolToken{"active"}
	TAddress                       = &SymbolToken{"address"}
	TAdministrator                 = &SymbolToken{"administrator"}
	TAgent                         = &SymbolToken{"agent"}
	TAggregate                     = &SymbolToken{"aggregate"}
	TAlarm                         = &SymbolToken{"alarm"}
	TAlgorithm                     = &SymbolToken{"algorithm"}
	TAlias                         = &SymbolToken{"alias"}
	TAlive                         = &SymbolToken{"alive"}
	TAll                           = &SymbolToken{"all"}
	TAlwaysOn                      = &SymbolToken{"always-on"}
	TAny                           = &SymbolToken{"any"}
	TApi                           = &SymbolToken{"api"}
	TAp                            = &SymbolToken{"ap"}
	TArea                          = &SymbolToken{"area"}
	TArp                           = &SymbolToken{"arp"}
	TArrive                        = &SymbolToken{"arrive"}
	TAspath                        = &SymbolToken{"aspath"}
	TAtCommand                     = &SymbolToken{"at-command"}
	TAt                            = &SymbolToken{"at"}
	TAttribute                     = &SymbolToken{"attribute"}
	TAuthError                     = &SymbolToken{"auth-error"}
	TAuth                          = &SymbolToken{"auth"}
	TAutonomousSystem              = &SymbolToken{"autonomous-system"}
	TAutoSearch                    = &SymbolToken{"auto-search"}
	TAuto                          = &SymbolToken{"auto"}
	TBackup                        = &SymbolToken{"backup"}
	TBackwardCompatibility         = &SymbolToken{"backward-compatibility"}
	TBacp                          = &SymbolToken{"bacp"}
	TBandwidthMeasuring            = &SymbolToken{"bandwidth-measuring"}
	TBandwidth                     = &SymbolToken{"bandwidth"}
	TBap                           = &SymbolToken{"bap"}
	TBatch                         = &SymbolToken{"batch"}
	TBgp                           = &SymbolToken{"bgp"}
	TBind                          = &SymbolToken{"bind"}
	TBlock                         = &SymbolToken{"block"}
	TBootRomVersion                = &SymbolToken{"boot-rom-version"}
	TBoot                          = &SymbolToken{"boot"}
	TBridge_Interface              = &SymbolToken{"bridge_interface"}
	TBridge                        = &SymbolToken{"bridge"}
	TBrightness                    = &SymbolToken{"brightness"}
	TBroadcast                     = &SymbolToken{"broadcast"}
	TBusy                          = &SymbolToken{"busy"}
	TButton                        = &SymbolToken{"button"}
	TCache                         = &SymbolToken{"cache"}
	TCallback                      = &SymbolToken{"callback"}
	TCallee                        = &SymbolToken{"callee"}
	TCaller                        = &SymbolToken{"caller"}
	TCall                          = &SymbolToken{"call"}
	TCallId                        = &SymbolToken{"call-id"}
	TCarrier                       = &SymbolToken{"carrier"}
	TCcp                           = &SymbolToken{"ccp"}
	TCertificate                   = &SymbolToken{"certificate"}
	TCertreq                       = &SymbolToken{"certreq"}
	TChange                        = &SymbolToken{"change"}
	TChannel                       = &SymbolToken{"channel"}
	TChap                          = &SymbolToken{"chap"}
	TCharacter                     = &SymbolToken{"character"}
	TCheck                         = &SymbolToken{"check"}
	TChildExchange                 = &SymbolToken{"child-exchange"}
	TCir                           = &SymbolToken{"cir"}
	TClass                         = &SymbolToken{"class"}
	TClearCounter                  = &SymbolToken{"clear-counter"}
	TClearMacaddressTable          = &SymbolToken{"clear-macaddress-table"}
	TClear                         = &SymbolToken{"clear"}
	TClientIdentifier              = &SymbolToken{"client-identifier"}
	TClient                        = &SymbolToken{"client"}
	TClientPort                    = &SymbolToken{"client_port"}
	TClose                         = &SymbolToken{"close"}
	TCloud                         = &SymbolToken{"cloud"}
	TCode                          = &SymbolToken{"code"}
	TCold                          = &SymbolToken{"cold"}
	TColumns                       = &SymbolToken{"columns"}
	TCommand                       = &SymbolToken{"command"}
	TCommon                        = &SymbolToken{"common"}
	TCommunity                     = &SymbolToken{"community"}
	TCompliant                     = &SymbolToken{"compliant"}
	TCompression                   = &SymbolToken{"compression"}
	TConcentrator                  = &SymbolToken{"concentrator"}
	TConfigAutoSet                 = &SymbolToken{"config-auto-set"}
	TConfig                        = &SymbolToken{"config"}
	TConfigure                     = &SymbolToken{"configure"}
	TConfirm                       = &SymbolToken{"confirm"}
	TCongestion                    = &SymbolToken{"congestion"}
	TConnection                    = &SymbolToken{"connection"}
	TConnect                       = &SymbolToken{"connect"}
	TConsole                       = &SymbolToken{"console"}
	TContext                       = &SymbolToken{"context"}
	TControl                       = &SymbolToken{"control"}
	TConvert                       = &SymbolToken{"convert"}
	TCooperation                   = &SymbolToken{"cooperation"}
	TCopyright                     = &SymbolToken{"copyright"}
	TCopy                          = &SymbolToken{"copy"}
	TCosArg                        = &EqualArgToken{"cos="}
	TCost                          = &SymbolToken{"cost"}
	TCounterFrameRxType            = &SymbolToken{"counter-frame-rx-type"}
	TCounterFrameTxType            = &SymbolToken{"counter-frame-tx-type"}
	TCountHubOverflow              = &SymbolToken{"count-hub-overflow"}
	TCount                         = &SymbolToken{"count"}
	TCrl                           = &SymbolToken{"crl"}
	TCustomGui                     = &SymbolToken{"custom-gui"}
	TDad                           = &SymbolToken{"dad"}
	TDashboard                     = &SymbolToken{"dashboard"}
	TData                          = &SymbolToken{"data"}
	TDate                          = &SymbolToken{"date"}
	TDebug                         = &SymbolToken{"debug"}
	TDefaultRouteAdd               = &SymbolToken{"default-route-add"}
	TDefault                       = &SymbolToken{"default"}
	TDelayTimer                    = &SymbolToken{"delay-timer"}
	TDelete                        = &SymbolToken{"delete"}
	TDequeue                       = &SymbolToken{"dequeue"}
	TDescription                   = &SymbolToken{"description"}
	TDescriptor                    = &SymbolToken{"descriptor"}
	TDe                            = &SymbolToken{"de"}
	TDetection                     = &SymbolToken{"detection"}
	TDfBit                         = &SymbolToken{"df-bit"}
	TDhcpc                         = &SymbolToken{"dhcpc"}
	TDhcp                          = &SymbolToken{"dhcp"}
	TDiagnose                      = &SymbolToken{"diagnose"}
	TDiagnosis                     = &SymbolToken{"diagnosis"}
	TDial                          = &SymbolToken{"dial"}
	TDirectedBroadcast             = &SymbolToken{"directed-broadcast"}
	TDirectory                     = &SymbolToken{"directory"}
	TDisable                       = &SymbolToken{"disable"}
	TDiscard                       = &SymbolToken{"discard"}
	TDisconnect                    = &SymbolToken{"disconnect"}
	TDiscovery                     = &SymbolToken{"discovery"}
	TDisplay                       = &SymbolToken{"display"}
	TDivide                        = &SymbolToken{"divide"}
	TDlci                          = &SymbolToken{"dlci"}
	TDns                           = &SymbolToken{"dns"}
	TDomain                        = &SymbolToken{"domain"}
	TDownload                      = &SymbolToken{"download"}
	TDown                          = &SymbolToken{"down"}
	TDscp                          = &SymbolToken{"dscp"}
	TDsu                           = &SymbolToken{"dsu"}
	TDuplicate                     = &SymbolToken{"duplicate"}
	TDuration                      = &SymbolToken{"duration"}
	TDynamic                       = &SymbolToken{"dynamic"}
	TEap                           = &SymbolToken{"eap"}
	TEchoReply                     = &SymbolToken{"echo-reply"}
	TEcho                          = &SymbolToken{"echo"}
	TEmbedded                      = &SymbolToken{"embedded"}
	TEnable                        = &SymbolToken{"enable"}
	TEncapsulation                 = &SymbolToken{"encapsulation"}
	TEncrypted                     = &SymbolToken{"encrypted"}
	TEncryption                    = &SymbolToken{"encryption"}
	TEncrypt                       = &SymbolToken{"encrypt"}
	TEndId                         = &SymbolToken{"end-id"}
	TEndpoint                      = &SymbolToken{"endpoint"}
	TEnergySaving                  = &SymbolToken{"energy-saving"}
	TEngine                        = &SymbolToken{"engine"}
	TEntire                        = &SymbolToken{"entire"}
	TEntry                         = &SymbolToken{"entry"}
	TEnvironment                   = &SymbolToken{"environment"}
	TEqual                         = &SymbolToken{"equal"}
	TErrorDecryptedIpsec           = &SymbolToken{"error-decrypted-ipsec"}
	TEspEncapsulation              = &SymbolToken{"esp-encapsulation"}
	TEthernet                      = &SymbolToken{"ethernet"}
	TExec                          = &SymbolToken{"exec"}
	TExecute                       = &SymbolToken{"execute"}
	TExit                          = &SymbolToken{"exit"}
	TExport                        = &SymbolToken{"export"}
	TExternalMemory                = &SymbolToken{"external-memory"}
	TExternal                      = &SymbolToken{"external"}
	TFacility                      = &SymbolToken{"facility"}
	TFallback                      = &SymbolToken{"fallback"}
	TFastpathFragmentFunction      = &SymbolToken{"fastpath-fragment-function"}
	TFast                          = &SymbolToken{"fast"}
	TFib                           = &SymbolToken{"fib"}
	TFilename                      = &SymbolToken{"filename"}
	TFile                          = &SymbolToken{"file"}
	TFilter                        = &SymbolToken{"filter"}
	TFirmwareRevision              = &SymbolToken{"firmware-revision"}
	TFirmware                      = &SymbolToken{"firmware"}
	TFlow                          = &SymbolToken{"flow"}
	TFollow                        = &SymbolToken{"follow"}
	TForced                        = &SymbolToken{"forced"}
	TForce                         = &SymbolToken{"force"}
	TForceToAdvertise              = &SymbolToken{"force-to-advertise"}
	TFormat                        = &SymbolToken{"format"}
	TForward                       = &SymbolToken{"forward"}
	TFqdn                          = &SymbolToken{"fqdn"}
	TFragment                      = &SymbolToken{"fragment"}
	TFrom                          = &SymbolToken{"from"}
	TFr                            = &SymbolToken{"fr"}
	TFtp                           = &SymbolToken{"ftp"}
	TFunction                      = &SymbolToken{"function"}
	TGarbageCollect                = &SymbolToken{"garbage-collect"}
	TGateway                       = &SymbolToken{"gateway"}
	TGenerate                      = &SymbolToken{"generate"}
	TGet                           = &SymbolToken{"get"}
	TGo                            = &SymbolToken{"go"}
	TGrep                          = &SymbolToken{"grep"}
	TGroup                         = &SymbolToken{"group"}
	TGuiForwarder                  = &SymbolToken{"gui-forwarder"}
	THash                          = &SymbolToken{"hash"}
	THeartbeat2                    = &SymbolToken{"heartbeat2"}
	THeartbeat                     = &SymbolToken{"heartbeat"}
	THelp                          = &SymbolToken{"help"}
	THide                          = &SymbolToken{"hide"}
	THistoryNum                    = &SymbolToken{"history-num"}
	THistory                       = &SymbolToken{"history"}
	THold                          = &SymbolToken{"hold"}
	THop                           = &SymbolToken{"hop"}
	THostname                      = &SymbolToken{"hostname"}
	THosts                         = &SymbolToken{"hosts"}
	THost                          = &SymbolToken{"host"}
	THttpd                         = &SymbolToken{"httpd"}
	THttpsProxy                    = &SymbolToken{"https-proxy"}
	THttp                          = &SymbolToken{"http"}
	TIcmp                          = &SymbolToken{"icmp"}
	TId                            = &SymbolToken{"id"}
	TIfindex                       = &SymbolToken{"ifindex"}
	TIgmp                          = &SymbolToken{"igmp"}
	TIke                           = &SymbolToken{"ike"}
	TIllegalSpi                    = &SymbolToken{"illegal-spi"}
	TImplicitRoute                 = &SymbolToken{"implicit-route"}
	TImport                        = &SymbolToken{"import"}
	TInarp                         = &SymbolToken{"inarp"}
	TInbound                       = &SymbolToken{"inbound"}
	TIncoming                      = &SymbolToken{"incoming"}
	TIndex                         = &SymbolToken{"index"}
	TInfinity                      = &SymbolToken{"infinity"}
	TInfo                          = &SymbolToken{"info"}
	TInitialize                    = &SymbolToken{"initialize"}
	TInner                         = &SymbolToken{"inner"}
	TInput                         = &SymbolToken{"input"}
	TInterfaceName                 = &SymbolToken{"interface-name"}
	TInterfaceRouteAdd             = &SymbolToken{"interface-route-add"}
	TInterface                     = &SymbolToken{"interface"}
	TInterleave                    = &SymbolToken{"interleave"}
	TInterval                      = &SymbolToken{"interval"}
	TIntrusion                     = &SymbolToken{"intrusion"}
	TInvalidSession                = &SymbolToken{"invalid-session"}
	TIpaddress                     = &SymbolToken{"ipaddress"}
	TIpcomp                        = &SymbolToken{"ipcomp"}
	TIpcp                          = &SymbolToken{"ipcp"}
	TIpip                          = &SymbolToken{"ipip"}
	TIpsec                         = &SymbolToken{"ipsec"}
	TIp                            = &SymbolToken{"ip"}
	TIpv6Cp                        = &SymbolToken{"ipv6cp"}
	TIpv6                          = &SymbolToken{"ipv6"}
	TIsdn                          = &SymbolToken{"isdn"}
	TJate                          = &SymbolToken{"jate"}
	TJoinPrune                     = &SymbolToken{"join-prune"}
	TKeepalive                     = &SymbolToken{"keepalive"}
	TKey                           = &SymbolToken{"key"}
	TKnown                         = &SymbolToken{"known"}
	TL2Ms                          = &SymbolToken{"l2ms"}
	TL2Tp                          = &SymbolToken{"l2tp"}
	TLagType                       = &SymbolToken{"lag-type"}
	TLanMap                        = &SymbolToken{"lan-map"}
	TLan                           = &SymbolToken{"lan"}
	TLcp                           = &SymbolToken{"lcp"}
	TLearning                      = &SymbolToken{"learning"}
	TLeased                        = &SymbolToken{"leased"}
	TLease                         = &SymbolToken{"lease"}
	TLedBrightness                 = &SymbolToken{"led-brightness"}
	TLed                           = &SymbolToken{"led"}
	TLength                        = &SymbolToken{"length"}
	TLess                          = &SymbolToken{"less"}
	TLicenseKey                    = &SymbolToken{"license-key"}
	TLimitation                    = &SymbolToken{"limitation"}
	TLimit                         = &SymbolToken{"limit"}
	TLines                         = &SymbolToken{"lines"}
	TLine                          = &SymbolToken{"line"}
	TLinkAggregation               = &SymbolToken{"link-aggregation"}
	TLinkdown                      = &SymbolToken{"linkdown"}
	TLinkRefresh                   = &SymbolToken{"link-refresh"}
	TLinkUpdown                    = &SymbolToken{"link-updown"}
	TLinkup                        = &SymbolToken{"linkup"}
	TListen                        = &SymbolToken{"listen"}
	TList                          = &SymbolToken{"list"}
	TLmi                           = &SymbolToken{"lmi"}
	TLoad                          = &SymbolToken{"load"}
	TLoadWatch                     = &SymbolToken{"load-watch"}
	TLocal                         = &SymbolToken{"local"}
	TLogin                         = &SymbolToken{"login"}
	TLog                           = &SymbolToken{"log"}
	TLoopback                      = &SymbolToken{"loopback"}
	TLoopdetectCount               = &SymbolToken{"loopdetect-count"}
	TLoopdetectLinkdown            = &SymbolToken{"loopdetect-linkdown"}
	TLoopdetectPortUse             = &SymbolToken{"loopdetect-port-use"}
	TLoopdetectRecoveryTimer       = &SymbolToken{"loopdetect-recovery-timer"}
	TLoopdetectTime                = &SymbolToken{"loopdetect-time"}
	TLoopdetectUseControlPacket    = &SymbolToken{"loopdetect-use-control-packet"}
	TLuac                          = &SymbolToken{"luac"}
	TLua                           = &SymbolToken{"lua"}
	TMacaddressAging               = &SymbolToken{"macaddress-aging"}
	TMacaddressAgingTimer          = &SymbolToken{"macaddress-aging-timer"}
	TMacaddress                    = &SymbolToken{"macaddress"}
	TMacro                         = &SymbolToken{"macro"}
	TMagicnumber                   = &SymbolToken{"magicnumber"}
	TMailNotify                    = &SymbolToken{"mail-notify"}
	TMail                          = &SymbolToken{"mail"}
	TMake                          = &SymbolToken{"make"}
	TManual                        = &SymbolToken{"manual"}
	TMapping                       = &SymbolToken{"mapping"}
	TMappingArg                    = &EqualArgToken{"mapping="}
	TMap                           = &SymbolToken{"map"}
	TMargin                        = &SymbolToken{"margin"}
	TMaskReply                     = &SymbolToken{"mask-reply"}
	TMasquerade                    = &SymbolToken{"masquerade"}
	TMasterclock                   = &SymbolToken{"masterclock"}
	TMaxauthreq                    = &SymbolToken{"maxauthreq"}
	TMaxchallenge                  = &SymbolToken{"maxchallenge"}
	TMaxconfigure                  = &SymbolToken{"maxconfigure"}
	TMaxDetect                     = &SymbolToken{"max-detect"}
	TMaxfailure                    = &SymbolToken{"maxfailure"}
	TMaxlink                       = &SymbolToken{"maxlink"}
	TMaxretry                      = &SymbolToken{"maxretry"}
	TMax                           = &SymbolToken{"max"}
	TMaxterminate                  = &SymbolToken{"maxterminate"}
	TMember                        = &SymbolToken{"member"}
	TMerge                         = &SymbolToken{"merge"}
	TMessageIdControl              = &SymbolToken{"message-id-control"}
	TMethod                        = &SymbolToken{"method"}
	TMinlink                       = &SymbolToken{"minlink"}
	TMirroringDest                 = &SymbolToken{"mirroring-dest"}
	TMirroringSrcRx                = &SymbolToken{"mirroring-src-rx"}
	TMirroringSrcTx                = &SymbolToken{"mirroring-src-tx"}
	TMirroringUse                  = &SymbolToken{"mirroring-use"}
	TMld                           = &SymbolToken{"mld"}
	TMobile                        = &SymbolToken{"mobile"}
	TModeCfg                       = &SymbolToken{"mode-cfg"}
	TModelName                     = &SymbolToken{"model-name"}
	TModem                         = &SymbolToken{"modem"}
	TMode                          = &SymbolToken{"mode"}
	TMonitor                       = &SymbolToken{"monitor"}
	TMp                            = &SymbolToken{"mp"}
	TMroute                        = &SymbolToken{"mroute"}
	TMru                           = &SymbolToken{"mru"}
	TMscbcp                        = &SymbolToken{"mscbcp"}
	TMsext                         = &SymbolToken{"msext"}
	TMss                           = &SymbolToken{"mss"}
	TMtu                           = &SymbolToken{"mtu"}
	TMulticast                     = &SymbolToken{"multicast"}
	TMultipoint                    = &SymbolToken{"multipoint"}
	TMulti                         = &SymbolToken{"multi"}
	TMyname                        = &SymbolToken{"myname"}
	TName                          = &SymbolToken{"name"}
	TNameArg                       = &EqualArgToken{"name="}
	TNat                           = &SymbolToken{"nat"}
	TNatTraversal                  = &SymbolToken{"nat-traversal"}
	TNd                            = &SymbolToken{"nd"}
	TNegotiateStrictly             = &SymbolToken{"negotiate-strictly"}
	TNegotiation                   = &SymbolToken{"negotiation"}
	TNeighbor                      = &SymbolToken{"neighbor"}
	TNetvolanteDns                 = &SymbolToken{"netvolante-dns"}
	TNetwork                       = &SymbolToken{"network"}
	TNgn                           = &SymbolToken{"ngn"}
	TNoEncryption                  = &SymbolToken{"no-encryption"}
	TNo                            = &SymbolToken{"no"}
	TNoticeInterval                = &SymbolToken{"notice-interval"}
	TNotice                        = &SymbolToken{"notice"}
	TNotify                        = &SymbolToken{"notify"}
	TNslookup                      = &SymbolToken{"nslookup"}
	TNsTriggerDad                  = &SymbolToken{"ns-trigger-dad"}
	TNtpdate                       = &SymbolToken{"ntpdate"}
	TNtp                           = &SymbolToken{"ntp"}
	TNumberOnly                    = &SymbolToken{"number-only"}
	TNumber                        = &SymbolToken{"number"}
	TOff                           = &SymbolToken{"off"}
	TOn                            = &SymbolToken{"on"}
	TOpenssh                       = &SymbolToken{"openssh"}
	TOperation                     = &SymbolToken{"operation"}
	TOption                        = &SymbolToken{"option"}
	TOrder                         = &SymbolToken{"order"}
	TOspf                          = &SymbolToken{"ospf"}
	TOuter                         = &SymbolToken{"outer"}
	TOutput                        = &SymbolToken{"output"}
	TOvercurrent                   = &SymbolToken{"overcurrent"}
	TOwn                           = &SymbolToken{"own"}
	TPacketBuffer                  = &SymbolToken{"packet-buffer"}
	TPacketdump                    = &SymbolToken{"packetdump"}
	TPacketScheduling              = &SymbolToken{"packet-scheduling"}
	TPacketTooBigForTruncated      = &SymbolToken{"packet-too-big-for-truncated"}
	TPacketTooBig                  = &SymbolToken{"packet-too-big"}
	TPadi                          = &SymbolToken{"padi"}
	TPadr                          = &SymbolToken{"padr"}
	TPap                           = &SymbolToken{"pap"}
	TParameterProblem              = &SymbolToken{"parameter-problem"}
	TParameter                     = &SymbolToken{"parameter"}
	TPassive                       = &SymbolToken{"passive"}
	TPassThrough                   = &SymbolToken{"pass-through"}
	TPassword                      = &SymbolToken{"password"}
	TPayload                       = &SymbolToken{"payload"}
	TPerformanceTest               = &SymbolToken{"performance-test"}
	TPeriodicPrune                 = &SymbolToken{"periodic-prune"}
	TPermit                        = &SymbolToken{"permit"}
	TPfc                           = &SymbolToken{"pfc"}
	TPfs                           = &SymbolToken{"pfs"}
	TPiafs                         = &SymbolToken{"piafs"}
	TPilot                         = &SymbolToken{"pilot"}
	TPim                           = &SymbolToken{"pim"}
	TPing6                         = &SymbolToken{"ping6"}
	TPing                          = &SymbolToken{"ping"}
	TPin                           = &SymbolToken{"pin"}
	TPki                           = &SymbolToken{"pki"}
	TPNUatype                      = &SymbolToken{"p-n-uatype"}
	TPoeClass                      = &SymbolToken{"poe-class"}
	TPolicy                        = &SymbolToken{"policy"}
	TPool                          = &SymbolToken{"pool"}
	TPop                           = &SymbolToken{"pop"}
	TPortAutoCrossover             = &SymbolToken{"port-auto-crossover"}
	TPortBlockingControlPacket     = &SymbolToken{"port-blocking-control-packet"}
	TPortBlockingDataPacket        = &SymbolToken{"port-blocking-data-packet"}
	TPortFlowControl               = &SymbolToken{"port-flow-control"}
	TPortMirroring                 = &SymbolToken{"port-mirroring"}
	TPortSpeedDownshift            = &SymbolToken{"port-speed-downshift"}
	TPortSpeed                     = &SymbolToken{"port-speed"}
	TPort                          = &SymbolToken{"port"}
	TPortUse                       = &SymbolToken{"port-use"}
	TPppoe                         = &SymbolToken{"pppoe"}
	TPpp                           = &SymbolToken{"ppp"}
	TPp                            = &SymbolToken{"pp"}
	TPptp                          = &SymbolToken{"pptp"}
	TPrecedence                    = &SymbolToken{"precedence"}
	TPreference                    = &SymbolToken{"preference"}
	TPrefix                        = &SymbolToken{"prefix"}
	TPreSharedKey                  = &SymbolToken{"pre-shared-key"}
	TPri                           = &SymbolToken{"pri"}
	TPrivacy                       = &SymbolToken{"privacy"}
	TPrivate                       = &SymbolToken{"private"}
	TProcess                       = &SymbolToken{"process"}
	TProhibit                      = &SymbolToken{"prohibit"}
	TPrompt                        = &SymbolToken{"prompt"}
	TProperty                      = &SymbolToken{"property"}
	TProposalLimitation            = &SymbolToken{"proposal-limitation"}
	TProtocol                      = &SymbolToken{"protocol"}
	TProvider                      = &SymbolToken{"provider"}
	TProxyAccess                   = &SymbolToken{"proxy-access"}
	TProxyarp                      = &SymbolToken{"proxyarp"}
	TProxy                         = &SymbolToken{"proxy"}
	TPublic                        = &SymbolToken{"public"}
	TQacTm                         = &SymbolToken{"qac-tm"}
	TQosDscpRemarkClass            = &SymbolToken{"qos-dscp-remark-class"}
	TQosDscpRemarkType             = &SymbolToken{"qos-dscp-remark-type"}
	TQosPolicingSpeed              = &SymbolToken{"qos-policing-speed"}
	TQosPolicingUse                = &SymbolToken{"qos-policing-use"}
	TQosShapingSpeed               = &SymbolToken{"qos-shaping-speed"}
	TQosShapingUse                 = &SymbolToken{"qos-shaping-use"}
	TQosSpeedUnit                  = &SymbolToken{"qos-speed-unit"}
	TQos                           = &SymbolToken{"qos"}
	TQualified                     = &SymbolToken{"qualified"}
	TQueue                         = &SymbolToken{"queue"}
	TQuit                          = &SymbolToken{"quit"}
	TQvalue                        = &SymbolToken{"qvalue"}
	TRadius                        = &SymbolToken{"radius"}
	TRange                         = &SymbolToken{"range"}
	TRdate                         = &SymbolToken{"rdate"}
	TRead                          = &SymbolToken{"read"}
	TReadOnly                      = &SymbolToken{"read-only"}
	TReadWrite                     = &SymbolToken{"read-write"}
	TRebound                       = &SymbolToken{"rebound"}
	TReceiveBufferSize             = &SymbolToken{"receive-buffer-size"}
	TReceive                       = &SymbolToken{"receive"}
	TRecord                        = &SymbolToken{"record"}
	TRecovery                      = &SymbolToken{"recovery"}
	TRecursive                     = &SymbolToken{"recursive"}
	TRedirect                      = &SymbolToken{"redirect"}
	TRefer                         = &SymbolToken{"refer"}
	TRefresher                     = &SymbolToken{"refresher"}
	TRefresh                       = &SymbolToken{"refresh"}
	TRegisterChecksum              = &SymbolToken{"register-checksum"}
	TRegister                      = &SymbolToken{"register"}
	TReject                        = &SymbolToken{"reject"}
	TRelay                         = &SymbolToken{"relay"}
	TRelease                       = &SymbolToken{"release"}
	TRemote                        = &SymbolToken{"remote"}
	TRemove                        = &SymbolToken{"remove"}
	TRename                        = &SymbolToken{"rename"}
	TRendezvousPoint               = &SymbolToken{"rendezvous-point"}
	TRenumbering                   = &SymbolToken{"renumbering"}
	TRepeatControl                 = &SymbolToken{"repeat-control"}
	TReport                        = &SymbolToken{"report"}
	TRequest                       = &SymbolToken{"request"}
	TRequestUri                    = &SymbolToken{"request-uri"}
	TReric                         = &SymbolToken{"reric"}
	TResetLoopdetect               = &SymbolToken{"reset-loopdetect"}
	TReset                         = &SymbolToken{"reset"}
	TResolv                        = &SymbolToken{"resolv"}
	TResponse                      = &SymbolToken{"response"}
	TRestartPoeSupply              = &SymbolToken{"restart-poe-supply"}
	TRestart                       = &SymbolToken{"restart"}
	TRestrictDanglingSa            = &SymbolToken{"restrict-dangling-sa"}
	TRetry                         = &SymbolToken{"retry"}
	TRevisionDown                  = &SymbolToken{"revision-down"}
	TRevisionUp                    = &SymbolToken{"revision-up"}
	TRfc2131                       = &SymbolToken{"rfc2131"}
	TRh0                           = &SymbolToken{"rh0"}
	TRinging                       = &SymbolToken{"ringing"}
	TRip                           = &SymbolToken{"rip"}
	TRlogin                        = &SymbolToken{"rlogin"}
	TRollback                      = &SymbolToken{"rollback"}
	TRotate                        = &SymbolToken{"rotate"}
	TRouterId                      = &SymbolToken{"router-id"}
	TRouter                        = &SymbolToken{"router"}
	TRoute                         = &SymbolToken{"route"}
	TRouting                       = &SymbolToken{"routing"}
	TRtadv                         = &SymbolToken{"rtadv"}
	TRtfs                          = &SymbolToken{"rtfs"}
	TRule                          = &SymbolToken{"rule"}
	TSa                            = &SymbolToken{"sa"}
	TSave                          = &SymbolToken{"save"}
	TSchedule                      = &SymbolToken{"schedule"}
	TScope                         = &SymbolToken{"scope"}
	TScp                           = &SymbolToken{"scp"}
	TSd                            = &SymbolToken{"sd"}
	TSecondary                     = &SymbolToken{"secondary"}
	TSecret                        = &SymbolToken{"secret"}
	TSecure                        = &SymbolToken{"secure"}
	TSecurity                      = &SymbolToken{"security"}
	TSelection                     = &SymbolToken{"selection"}
	TSelect                        = &SymbolToken{"select"}
	TSendOnlyLinkup                = &SymbolToken{"send-only-linkup"}
	TSend                          = &SymbolToken{"send"}
	TSendWaitTime                  = &SymbolToken{"send-wait-time"}
	TSeparateL2SwitchPort          = &SymbolToken{"separate-l2switch-port"}
	TSerialNumber                  = &SymbolToken{"serial-number"}
	TServer                        = &SymbolToken{"server"}
	TServiceName                   = &SymbolToken{"service-name"}
	TService                       = &SymbolToken{"service"}
	TSession                       = &SymbolToken{"session"}
	TSetDefaultConfig              = &SymbolToken{"set-default-config"}
	TSetDefaultExec                = &SymbolToken{"set-default-exec"}
	TSetSerialBaudrate             = &SymbolToken{"set-serial-baudrate"}
	TSet                           = &SymbolToken{"set"}
	TSetup                         = &SymbolToken{"setup"}
	TSftpd                         = &SymbolToken{"sftpd"}
	TShapingLevelArg               = &EqualArgToken{"shaping-level="}
	TShow                          = &SymbolToken{"show"}
	TShutdown                      = &SymbolToken{"shutdown"}
	TSignalStrength                = &SymbolToken{"signal-strength"}
	TSilent                        = &SymbolToken{"silent"}
	TSimpleService                 = &SymbolToken{"simple-service"}
	TSip                           = &SymbolToken{"sip"}
	TSize                          = &SymbolToken{"size"}
	TSmtp                          = &SymbolToken{"smtp"}
	TSnapshot                      = &SymbolToken{"snapshot"}
	TSnmp                          = &SymbolToken{"snmp"}
	TSnmpv2c                       = &SymbolToken{"snmpv2c"}
	TSnmpv2C                       = &SymbolToken{"snmpv2c"}
	TSnmpv3                        = &SymbolToken{"snmpv3"}
	TSntpd                         = &SymbolToken{"sntpd"}
	TSourceRoute                   = &SymbolToken{"source-route"}
	TSource                        = &SymbolToken{"source"}
	TSparse                        = &SymbolToken{"sparse"}
	TSpeed                         = &SymbolToken{"speed"}
	TSpoof                         = &SymbolToken{"spoof"}
	TSrcport                       = &SymbolToken{"srcport"}
	TSshd                          = &SymbolToken{"sshd"}
	TSsh                           = &SymbolToken{"ssh"}
	TStar                          = &SymbolToken{"star"}
	TStartPoeSupply                = &SymbolToken{"start-poe-supply"}
	TStart                         = &SymbolToken{"start"}
	TStartup                       = &SymbolToken{"startup"}
	TStatic                        = &SymbolToken{"static"}
	TStatistics                    = &SymbolToken{"statistics"}
	TStatusComboPort               = &SymbolToken{"status-combo-port"}
	TStatusCounterFrameRx          = &SymbolToken{"status-counter-frame-rx"}
	TStatusCounterFrameTx          = &SymbolToken{"status-counter-frame-tx"}
	TStatusCounterOctetRx          = &SymbolToken{"status-counter-octet-rx"}
	TStatusCounterOctetTx          = &SymbolToken{"status-counter-octet-tx"}
	TStatusFanRpm                  = &SymbolToken{"status-fan-rpm"}
	TStatusFan                     = &SymbolToken{"status-fan"}
	TStatusLedMode                 = &SymbolToken{"status-led-mode"}
	TStatusLed                     = &SymbolToken{"status-led"}
	TStatusLoopdetectPort          = &SymbolToken{"status-loopdetect-port"}
	TStatusLoopdetectRecoveryTimer = &SymbolToken{"status-loopdetect-recovery-timer"}
	TStatusMacaddressAddr          = &SymbolToken{"status-macaddress-addr"}
	TStatusMacaddressPort          = &SymbolToken{"status-macaddress-port"}
	TStatusPoeDetectClass          = &SymbolToken{"status-poe-detect-class"}
	TStatusPoeState                = &SymbolToken{"status-poe-state"}
	TStatusPoeSupplyDetail         = &SymbolToken{"status-poe-supply-detail"}
	TStatusPoeSupply               = &SymbolToken{"status-poe-supply"}
	TStatusPoeSupplyTotal          = &SymbolToken{"status-poe-supply-total"}
	TStatusPoeTemperature          = &SymbolToken{"status-poe-temperature"}
	TStatusPortSfpRxPower          = &SymbolToken{"status-port-sfp-rx-power"}
	TStatusPortSpeed               = &SymbolToken{"status-port-speed"}
	TStatus                        = &SymbolToken{"status"}
	TStealth                       = &SymbolToken{"stealth"}
	TStopPoeSupply                 = &SymbolToken{"stop-poe-supply"}
	TString                        = &SymbolToken{"string"}
	TStubhost                      = &SymbolToken{"stubhost"}
	TStub                          = &SymbolToken{"stub"}
	TSubject                       = &SymbolToken{"subject"}
	TSummary                       = &SymbolToken{"summary"}
	TSupersede                     = &SymbolToken{"supersede"}
	TSwitchingHub                  = &SymbolToken{"switching-hub"}
	TSwitch                        = &SymbolToken{"switch"}
	TSyscontact                    = &SymbolToken{"syscontact"}
	TSyslocation                   = &SymbolToken{"syslocation"}
	TSyslog                        = &SymbolToken{"syslog"}
	TSysname                       = &SymbolToken{"sysname"}
	TSystemMacaddress              = &SymbolToken{"system-macaddress"}
	TSystemName                    = &SymbolToken{"system-name"}
	TSystem                        = &SymbolToken{"system"}
	TSystemUptime                  = &SymbolToken{"system-uptime"}
	TTable                         = &SymbolToken{"table"}
	TTake                          = &SymbolToken{"take"}
	TTcp                           = &SymbolToken{"tcp"}
	TTechinfo                      = &SymbolToken{"techinfo"}
	TTelnetd                       = &SymbolToken{"telnetd"}
	TTelnet                        = &SymbolToken{"telnet"}
	TTemperature                   = &SymbolToken{"temperature"}
	TTemplate                      = &SymbolToken{"template"}
	TTerminal                      = &SymbolToken{"terminal"}
	TTerminate                     = &SymbolToken{"terminate"}
	TTerminator                    = &SymbolToken{"terminator"}
	TText                          = &SymbolToken{"text"}
	TTftp                          = &SymbolToken{"tftp"}
	TThreshold                     = &SymbolToken{"threshold"}
	TTimeExceeded                  = &SymbolToken{"time-exceeded"}
	TTimeout                       = &SymbolToken{"timeout"}
	TTimer                         = &SymbolToken{"timer"}
	TTimestampReply                = &SymbolToken{"timestamp-reply"}
	TTime                          = &SymbolToken{"time"}
	TTimezone                      = &SymbolToken{"timezone"}
	TTos                           = &SymbolToken{"tos"}
	TTo                            = &SymbolToken{"to"}
	TTotal                         = &SymbolToken{"total"}
	TTraceroute6                   = &SymbolToken{"traceroute6"}
	TTraceroute                    = &SymbolToken{"traceroute"}
	TTraffic                       = &SymbolToken{"traffic"}
	TTransmit                      = &SymbolToken{"transmit"}
	TTransport                     = &SymbolToken{"transport"}
	TTrap                          = &SymbolToken{"trap"}
	TTrigger                       = &SymbolToken{"trigger"}
	TTrust                         = &SymbolToken{"trust"}
	TTtl                           = &SymbolToken{"ttl"}
	TTunnel                        = &SymbolToken{"tunnel"}
	TType                          = &SymbolToken{"type"}
	TUnconvertible                 = &SymbolToken{"unconvertible"}
	TUnqualified                   = &SymbolToken{"unqualified"}
	TUnreachableForTruncated       = &SymbolToken{"unreachable-for-truncated"}
	TUnreachable                   = &SymbolToken{"unreachable"}
	TUpdate                        = &SymbolToken{"update"}
	TUpload                        = &SymbolToken{"upload"}
	TUpnp                          = &SymbolToken{"upnp"}
	TUrl                           = &SymbolToken{"url"}
	TUsbDownload                   = &SymbolToken{"usb-download"}
	TUsbhost                       = &SymbolToken{"usbhost"}
	TUsername                      = &SymbolToken{"username"}
	TUserSpecify                   = &SymbolToken{"user-specify"}
	TUser                          = &SymbolToken{"user"}
	TUse                           = &SymbolToken{"use"}
	TUsm                           = &SymbolToken{"usm"}
	TVacm                          = &SymbolToken{"vacm"}
	TVendorname                    = &SymbolToken{"vendorname"}
	TVersion                       = &SymbolToken{"version"}
	TVidArg                        = &EqualArgToken{"vid="}
	TView                          = &SymbolToken{"view"}
	TVirtualLink                   = &SymbolToken{"virtual-link"}
	TVjc                           = &SymbolToken{"vjc"}
	TVlanAccess                    = &SymbolToken{"vlan-access"}
	TVlanId                        = &SymbolToken{"vlan-id"}
	TVlanMultiple                  = &SymbolToken{"vlan-multiple"}
	TVlanMultipleUse               = &SymbolToken{"vlan-multiple-use"}
	TVlanPortMode                  = &SymbolToken{"vlan-port-mode"}
	TVlan                          = &SymbolToken{"vlan"}
	TVlanTrunk                     = &SymbolToken{"vlan-trunk"}
	TVpn                           = &SymbolToken{"vpn"}
	TVrrp                          = &SymbolToken{"vrrp"}
	TWait                          = &SymbolToken{"wait"}
	TWan                           = &SymbolToken{"wan"}
	TWarning                       = &SymbolToken{"warning"}
	TWatch                         = &SymbolToken{"watch"}
	TWindow                        = &SymbolToken{"window"}
	TWindowScale                   = &SymbolToken{"window-scale"}
	TWins                          = &SymbolToken{"wins"}
	TWol                           = &SymbolToken{"wol"}
	TWrite                         = &SymbolToken{"write"}
	TXauth                         = &SymbolToken{"xauth"}
	TYno                           = &SymbolToken{"yno"}
	TYrifppdisplayatmib2           = &SymbolToken{"yrifppdisplayatmib2"}
	TYrifswitchdisplayatmib2       = &SymbolToken{"yrifswitchdisplayatmib2"}
	TYriftunneldisplayatmib2       = &SymbolToken{"yriftunneldisplayatmib2"}
	TYrswindex                     = &SymbolToken{"yrswindex"}

	TL2Interface = &L2InterfaceToken{}
	TCharsetName = &CharsetToken{}
	TSomething   = &AnyToken{}
	TOnOff       = &EitherToken{TOn, TOff}
)

type CommandHandler func(ctx context.Context, sim *SimulatorSession, tokens []TokenInstance) error

type TokensArityPair struct {
	Tokens []Token
	Arity  int
}

type CommandSpec struct {
	Tokens  interface{}
	Arity   int
	Handler CommandHandler
}

var (
	InvalidCommandName             = fmt.Errorf("invalid command name")
	IllegalKeyword                 = fmt.Errorf("illegal keyword")
	IncorrectPassword              = fmt.Errorf("incorrect password")
	AdministratorUseOnly           = fmt.Errorf("administrator use only")
	WrongNumberOfArgs              = fmt.Errorf("insufficient or too many parameters")
	PPSelectionRequired            = fmt.Errorf("execute pp select command beforehand")
	IntegerArgRequired             = fmt.Errorf("parameter must be an integer")
	ArgOutOfRange                  = fmt.Errorf("parameter out of range")
	UnrecognizedIPAddr             = fmt.Errorf("inrecognized IP address parameter")
	UnsavedChangesInVolatileMemory = &struct{}{}
	PromptSaveSettings             = &struct{}{}
)

func init() {
	var err error
	greetTemplate, err = template.New("").Funcs(template.FuncMap{
		"formatAsRFC822": formatAsRFC822,
	}).Parse(welcomeMessage)
	if err != nil {
		panic(err)
	}
}

type HardwareSpec struct {
	ModelName               string
	PhyInterfaces           []*PhyInterface
	PortVlanInterfaces      []*PortVlanInterface
	PortVlanGroupInterfaces []*PortVlanGroupInterface
	TaggedVlanInterfaces    map[*PhyInterface][]*TaggedVlanInterface
}

type NLSInfo struct {
	Name     string
	Variants []string
	Encoding string
	Locale   string
}

func (nls *NLSInfo) NewEncoder() *encoding.Encoder {
	encoding, err := encoding_index.IANA.Encoding(nls.Encoding)
	if err != nil {
		panic(err)
	}
	return encoding.NewEncoder()
}

type SystemSpec struct {
	Firmware
	HardwareSpec
	Charsets []NLSInfo
}

type Spec struct {
	HardwareSpec
	SystemSpec
}

type InterfaceConfig struct {
	sync.Mutex
	V4DHCP          bool
	V6RA            bool
	V4Address       net.IP
	V4Network       net.IPNet
	V4BroadcastAddr net.IP
	V6Address       net.IP
	V6Network       net.IPNet
}

func (ic *InterfaceConfig) Clone() *InterfaceConfig {
	ic.Lock()
	defer ic.Unlock()
	newConfig := &InterfaceConfig{}
	*newConfig = *ic
	newConfig.Mutex = sync.Mutex{}
	return newConfig
}

type SimulatorConfig struct {
	sync.Mutex
	Prompt     string
	Charset    NLSInfo
	Interfaces map[NetInterface]*InterfaceConfig
}

func (c *SimulatorConfig) SetPrompt(v string) {
	c.Lock()
	c.Prompt = v
	c.Unlock()
}

func (c *SimulatorConfig) Clone() *SimulatorConfig {
	c.Lock()
	defer c.Unlock()
	newConfig := &SimulatorConfig{}
	*newConfig = *c
	newConfig.Mutex = sync.Mutex{}
	for k, v := range c.Interfaces {
		newConfig.Interfaces[k] = v.Clone()
	}
	return newConfig
}

type Simulator struct {
	*Spec
	Log           Logger
	Commands      []*CommandSpec
	CommandTrie   TokenTrie
	RunningConfig *SimulatorConfig
	SavedConfig   *SimulatorConfig
}

func NewSimulator(spec *Spec, config *SimulatorConfig, log Logger) *Simulator {
	sim := &Simulator{
		Spec:          spec,
		SavedConfig:   config.Clone(),
		RunningConfig: config.Clone(),
		Log:           log,
		Commands: []*CommandSpec{
			CmdAccountThreshold,
			CmdAccountThresholdPp,
			CmdAdministrator,
			CmdAdministratorPassword,
			CmdAdministratorRadiusAuth,
			CmdAlarmBatch,
			CmdAlarmEntire,
			CmdAlarmHttpRevisionUp,
			CmdAlarmHttpUpload,
			CmdAlarmLua,
			CmdAlarmMobile,
			CmdAlarmSd,
			CmdAlarmStartup,
			CmdAlarmUsbhost,
			CmdAlias,
			CmdApConfigDirectory,
			CmdApConfigFilename,
			CmdApControlConfigAutoSetUse,
			CmdApControlConfigDelete,
			CmdApControlConfigGet,
			CmdApControlConfigSet,
			CmdApControlFirmwareUpdateGo,
			CmdApControlHttpProxyTimeout,
			CmdApControlHttpProxyUse,
			CmdApSelect,
			CmdAuthUser,
			CmdAuthUserAttribute,
			CmdAuthUserGroup,
			CmdAuthUserGroupAttribute,
			CmdBgpAggregate,
			CmdBgpAggregateFilter,
			CmdBgpAutonomousSystem,
			CmdBgpConfigureRefresh,
			CmdBgpExport,
			CmdBgpExportAspath,
			CmdBgpExportFilter,
			CmdBgpForceToAdvertise,
			CmdBgpImport,
			CmdBgpImportFilter,
			CmdBgpLog,
			CmdBgpNeighbor,
			CmdBgpNeighborPreSharedKey,
			CmdBgpPreference,
			CmdBgpRericInterval,
			CmdBgpRouterId,
			CmdBgpUse,
			CmdBridgeLearningSwitch,
			CmdBridgeLearningStatic,
			CmdBridgeLearningBridgeInterfaceTimer,
			CmdBridgeMember,
			CmdCall,
			CmdClearAccount,
			CmdClearAccountMobile,
			CmdClearAccountNgnData,
			CmdClearAccountPp,
			CmdClearAccountTunnel,
			CmdClearArp,
			CmdClearBootList,
			CmdClearBridgeLearning,
			CmdClearDiagnosisConfigPort,
			CmdClearDnsCache,
			CmdClearExternalMemorySyslog,
			CmdClearHeartbeat2,
			CmdClearHeartbeat2Id,
			CmdClearHeartbeat2Name,
			CmdClearInarp,
			CmdClearIpDynamicRouting,
			CmdClearIpInboundFilter,
			CmdClearIpPolicyFilter,
			CmdClearIpTrafficList,
			CmdClearIpv6DynamicRouting,
			CmdClearIpv6InboundFilter,
			CmdClearIpv6NeighborCache,
			CmdClearIpv6PolicyFilter,
			CmdClearLog,
			CmdClearMobileAccessLimitation,
			CmdClearMobileAccessLimitationPp,
			CmdClearNatDescriptorDynamic,
			CmdClearNatDescriptorInterfaceDynamic,
			CmdClearNatDescriptorInterfaceDynamicPp,
			CmdClearNatDescriptorInterfaceDynamicTunnel,
			CmdClearPppoePassThroughLearning,
			CmdClearPriStatus,
			CmdClearStatus,
			CmdClearSwitchingHubMacaddress,
			CmdClearUrlFilter,
			CmdClearUrlFilterPp,
			CmdClearUrlFilterTunnel,
			CmdCloudVpnBind,
			CmdCloudVpnName,
			CmdCloudVpnOption,
			CmdCloudVpnParameter,
			CmdCloudVpnService,
			CmdCloudVpnSetGo,
			CmdColdStart,
			CmdConfirm,
			CmdConnect,
			CmdConnectPp,
			CmdConnectTunnel,
			CmdConsoleCharacter,
			CmdConsoleColumns,
			CmdConsoleInfo,
			CmdConsoleLines,
			CmdConsolePrompt,
			CmdCooperation,
			CmdCooperationBandwidthMeasuringRemote,
			CmdCooperationLoadWatchControl,
			CmdCooperationLoadWatchRemote,
			CmdCooperationLoadWatchTrigger,
			CmdCooperationPort,
			CmdCooperationTypeGo,
			CmdCopy,
			CmdCopyConfig,
			CmdCopyExec,
			CmdDashboardAccumulate,
			CmdDate,
			CmdDelete,
			CmdDeleteConfig,
			CmdDeleteExec,
			CmdDeletePkiFile,
			CmdDescription,
			CmdDescriptionInterfaceDescription,
			CmdDescriptionYno,
			CmdDhcpClientClientIdentifier,
			CmdDhcpClientHostname,
			CmdDhcpClientOption,
			CmdDhcpConvertLeaseToBind,
			CmdDhcpDuplicateCheck,
			CmdDhcpManualLease,
			CmdDhcpManualRelease,
			CmdDhcpRelaySelect,
			CmdDhcpRelayServer,
			CmdDhcpRelayThreshold,
			CmdDhcpScope,
			CmdDhcpScopeBind,
			CmdDhcpScopeLeaseType,
			CmdDhcpScopeOption,
			CmdDhcpServerRfc2131Compliant,
			CmdDhcpService,
			CmdDiagnoseConfigPortAccess,
			CmdDiagnoseConfigPortMap,
			CmdDiagnosisConfigPortHistoryNum,
			CmdDiagnosisConfigPortMaxDetect,
			CmdDisconnect,
			CmdDisconnectIpConnection,
			CmdDisconnectIpv6Connection,
			CmdDisconnectPp,
			CmdDisconnectTunnel,
			CmdDisconnectUser,
			CmdDisconnectUserUserConnectionNo,
			CmdDnsCacheMaxEntry,
			CmdDnsCacheUse,
			CmdDnsDomain,
			CmdDnsHost,
			CmdDnsNoticeOrder,
			CmdDnsPrivateAddressSpoof,
			CmdDnsServer,
			CmdDnsServerDhcp,
			CmdDnsServerPp,
			CmdDnsServerSelect,
			CmdDnsService,
			CmdDnsServiceAaaaFilter,
			CmdDnsServiceFallback,
			CmdDnsSrcport,
			CmdDnsStatic,
			CmdDnsStaticTypeNameValueTtl,
			CmdDnsSyslogResolv,
			CmdEcho,
			CmdEmbeddedFile,
			CmdEthernetFilter,
			CmdEthernetInterfaceFilter,
			CmdExecuteAtCommand,
			CmdExecuteBatch,
			CmdExternalMemoryAcceleratorCacheSize,
			CmdExternalMemoryAutoSearchTime,
			CmdExternalMemoryBatchFilename,
			CmdExternalMemoryBootPermit,
			CmdExternalMemoryBootTimeout,
			CmdExternalMemoryCacheMode,
			CmdExternalMemoryConfigFilename,
			CmdExternalMemoryExecFilename,
			CmdExternalMemoryPerformanceTestGo,
			CmdExternalMemoryStatisticsFilenamePrefix,
			CmdExternalMemorySyslogFilename,
			CmdFrBackup,
			CmdFrCir,
			CmdFrCompressionUse,
			CmdFrCongestionControl,
			CmdFrDe,
			CmdFrDlci,
			CmdFrInarp,
			CmdFrLmi,
			CmdFrPpDequeueType,
			CmdGrep,
			CmdHeartbeat2Myname,
			CmdHeartbeat2Receive,
			CmdHeartbeat2ReceiveEnable,
			CmdHeartbeat2ReceiveLog,
			CmdHeartbeat2ReceiveMonitor,
			CmdHeartbeat2ReceiveRecordLimit,
			CmdHeartbeat2Transmit,
			CmdHeartbeat2TransmitEnable,
			CmdHeartbeat2TransmitInterval,
			CmdHeartbeat2TransmitLog,
			CmdHeartbeatPreSharedKey,
			CmdHeartbeatReceive,
			CmdHeartbeatSend,
			CmdHelp,
			CmdHttpdCustomGuiApiPassword,
			CmdHttpdCustomGuiApiUse,
			CmdHttpdCustomGuiUse,
			CmdHttpdCustomGuiUser,
			CmdHttpdHost,
			CmdHttpdListen,
			CmdHttpdProxyAccessL2MsPermit,
			CmdHttpdService,
			CmdHttpdTimeout,
			CmdHttpRevisionDownPermit,
			CmdHttpRevisionUpGo,
			CmdHttpRevisionUpPermit,
			CmdHttpRevisionUpProxy,
			CmdHttpRevisionUpSchedule,
			CmdHttpRevisionUpTimeout,
			CmdHttpRevisionUpUrl,
			CmdHttpUpload,
			CmdHttpUploadGo,
			CmdHttpUploadPermit,
			CmdHttpUploadProxy,
			CmdHttpUploadRetryInterval,
			CmdHttpUploadTimeout,
			CmdHttpUploadUrl,
			CmdInterfaceReset,
			CmdInterfaceResetPp,
			CmdIpArpTimer,
			CmdIpFilter,
			CmdIpFilterDirectedBroadcast,
			CmdIpFilterDynamic,
			CmdIpFilterDynamicTimer,
			CmdIpFilterFqdnTimer,
			CmdIpFilterSet,
			CmdIpFilterSourceRoute,
			CmdIpFlowLimit,
			CmdIpFlowTimer,
			CmdIpForwardFilter,
			CmdIpFragmentRemoveDfBit,
			CmdIpHost,
			CmdIpIcmpEchoReplySend,
			CmdIpIcmpEchoReplySendOnlyLinkup,
			CmdIpIcmpErrorDecryptedIpsecSend,
			CmdIpIcmpLog,
			CmdIpIcmpMaskReplySend,
			CmdIpIcmpParameterProblemSend,
			CmdIpIcmpRedirectReceive,
			CmdIpIcmpRedirectSend,
			CmdIpIcmpTimeExceededSend,
			CmdIpIcmpTimestampReplySend,
			CmdIpIcmpUnreachableForTruncatedSend,
			CmdIpIcmpUnreachableSend,
			CmdIPIfAddress,
			CmdIpImplicitRoutePreference,
			CmdIpInboundFilter,
			CmdIpInterfaceArpLog,
			CmdIpInterfaceArpMtuDiscovery,
			CmdIpInterfaceArpQueueLength,
			CmdIpInterfaceArpStatic,
			CmdIpInterfaceDhcpAutoDefaultRouteAdd,
			CmdIpInterfaceDhcpAutoInterfaceRouteAdd,
			CmdIpInterfaceDhcpLeaseTime,
			CmdIpInterfaceDhcpRetry,
			CmdIpInterfaceDhcpService,
			CmdIpInterfaceForwardFilter,
			CmdIpInterfaceIgmp,
			CmdIpInterfaceIgmpStatic,
			CmdIpInterfaceInboundFilterList,
			CmdIpInterfaceIntrusionDetection,
			CmdIpInterfaceIntrusionDetectionNoticeInterval,
			CmdIpInterfaceIntrusionDetectionRepeatControl,
			CmdIpInterfaceIntrusionDetectionReport,
			CmdIpInterfaceIntrusionDetectionThreshold,
			CmdIpInterfaceMtu,
			CmdIpInterfaceNatDescriptor,
			CmdIpInterfaceOspfArea,
			CmdIpInterfaceOspfNeighbor,
			CmdIpInterfacePimSparse,
			CmdIpInterfaceProxyarp,
			CmdIpInterfaceProxyarpVrrp,
			CmdIpInterfaceRebound,
			CmdIpInterfaceRipAuthKey,
			CmdIpInterfaceRipAuthKeyText,
			CmdIpInterfaceRipAuthType,
			CmdIpInterfaceRipFilter,
			CmdIpInterfaceRipForceToAdvertise,
			CmdIpInterfaceRipHop,
			CmdIpInterfaceRipReceive,
			CmdIpInterfaceRipSend,
			CmdIpInterfaceRipTrustGateway,
			CmdIpInterfaceSecondaryAddress,
			CmdIpInterfaceSecureFilter,
			CmdIpInterfaceSecureFilterName,
			CmdIpInterfaceTcpMssLimit,
			CmdIpInterfaceTcpWindowScale,
			CmdIpInterfaceTrafficList,
			CmdIpInterfaceTrafficListThreshold,
			CmdIpInterfaceVrrp,
			CmdIpInterfaceVrrpShutdownTrigger,
			CmdIpInterfaceWolRelay,
			CmdIpipKeepaliveLog,
			CmdIpipKeepaliveUse,
			CmdIpKeepalive,
			CmdIpLocalForwardFilter,
			CmdIpPimSparseJoinPruneSend,
			CmdIpPimSparseLog,
			CmdIpPimSparsePeriodicPruneSend,
			CmdIpPimSparseRegisterChecksum,
			CmdIpPimSparseRendezvousPointStatic,
			CmdIpPolicyAddressGroup,
			CmdIpPolicyFilter,
			CmdIpPolicyFilterSet,
			CmdIpPolicyFilterSetEnable,
			CmdIpPolicyFilterSetSwitch,
			CmdIpPolicyFilterTimer,
			CmdIpPolicyInterfaceGroup,
			CmdIpPolicyService,
			CmdIpPolicyServiceGroup,
			CmdIpPpAddress,
			CmdIpPpForwardFilter,
			CmdIpPpIgmp,
			CmdIpPpIgmpStatic,
			CmdIpPpInboundFilterList,
			CmdIpPpIntrusionDetection,
			CmdIpPpIntrusionDetectionNoticeInterval,
			CmdIpPpIntrusionDetectionRepeatControl,
			CmdIpPpIntrusionDetectionReport,
			CmdIpPpIntrusionDetectionThreshold,
			CmdIpPpMtu,
			CmdIpPpNatDescriptor,
			CmdIpPpOspfArea,
			CmdIpPpOspfNeighbor,
			CmdIpPpPimSparse,
			CmdIpPpRebound,
			CmdIpPpRemoteAddress,
			CmdIpPpRemoteAddressPool,
			CmdIpPpRipAuthKey,
			CmdIpPpRipAuthKeyText,
			CmdIpPpRipAuthType,
			CmdIpPpRipBackupInterface,
			CmdIpPpRipConnectInterval,
			CmdIpPpRipConnectSend,
			CmdIpPpRipDisconnectInterval,
			CmdIpPpRipDisconnectSend,
			CmdIpPpRipFilter,
			CmdIpPpRipForceToAdvertise,
			CmdIpPpRipHoldRouting,
			CmdIpPpRipHop,
			CmdIpPpRipHopDirectionHop,
			CmdIpPpRipReceive,
			CmdIpPpRipReceiveReceive,
			CmdIpPpRipSend,
			CmdIpPpRipTrustGateway,
			CmdIpPpSecureFilter,
			CmdIpPpSecureFilterName,
			CmdIpPpTcpMssLimit,
			CmdIpPpTcpWindowScale,
			CmdIpPpTrafficList,
			CmdIpPpTrafficListThreshold,
			CmdIpRoute,
			CmdIpRouteChangeLog,
			CmdIpRouting,
			CmdIpRoutingProcess,
			CmdIpsecAutoRefresh,
			CmdIpsecIkeAlwaysOn,
			CmdIpsecIkeAuthMethod,
			CmdIpsecIkeBackwardCompatibility,
			CmdIpsecIkeChildExchangeType,
			CmdIpsecIkeDuration,
			CmdIpsecIkeEapMyname,
			CmdIpsecIkeEapRequest,
			CmdIpsecIkeEapSendCertreq,
			CmdIpsecIkeEncryption,
			CmdIpsecIkeEspEncapsulation,
			CmdIpsecIkeGroup,
			CmdIpsecIkeHash,
			CmdIpsecIkeKeepaliveLog,
			CmdIpsecIkeKeepaliveUse,
			CmdIpsecIkeLicenseKey,
			CmdIpsecIkeLicenseKeyUse,
			CmdIpsecIkeLocalAddress,
			CmdIpsecIkeLocalId,
			CmdIpsecIkeLocalName,
			CmdIpsecIkeLog,
			CmdIpsecIkeMessageIdControl,
			CmdIpsecIkeModeCfgAddress,
			CmdIpsecIkeModeCfgAddressPool,
			CmdIpsecIkeModeCfgMethod,
			CmdIpsecIkeNatTraversal,
			CmdIpsecIkeNegotiateStrictly,
			CmdIpsecIkeNegotiationReceive,
			CmdIpsecIkePayloadType,
			CmdIpsecIkePfs,
			CmdIpsecIkePkiFile,
			CmdIpsecIkePreSharedKey,
			CmdIpsecIkeProposalLimitation,
			CmdIpsecIkeQueueLength,
			CmdIpsecIkeRemoteAddress,
			CmdIpsecIkeRemoteId,
			CmdIpsecIkeRemoteName,
			CmdIpsecIkeRestrictDanglingSa,
			CmdIpsecIkeRetry,
			CmdIpsecIkeSendInfo,
			CmdIpsecIkeVersion,
			CmdIpsecIkeXauthMyname,
			CmdIpsecIkeXauthRequest,
			CmdIpsecIpcompType,
			CmdIpsecLogIllegalSpi,
			CmdIpsecRefreshSa,
			CmdIpsecSaDelete,
			CmdIpsecSaPolicy,
			CmdIpsecTransport,
			CmdIpsecTransportTemplate,
			CmdIpsecTunnel,
			CmdIpsecTunnelFastpathFragmentFunctionFollowDfBit,
			CmdIpsecTunnelOuterDfBit,
			CmdIpsecUse,
			CmdIpSimpleService,
			CmdIpStealth,
			CmdIpTosSupersede,
			CmdIpTunnelAddress,
			CmdIpTunnelForwardFilter,
			CmdIpTunnelIgmp,
			CmdIpTunnelIgmpStatic,
			CmdIpTunnelInboundFilterList,
			CmdIpTunnelIntrusionDetection,
			CmdIpTunnelIntrusionDetectionNoticeInterval,
			CmdIpTunnelIntrusionDetectionRepeatControl,
			CmdIpTunnelIntrusionDetectionReport,
			CmdIpTunnelIntrusionDetectionThreshold,
			CmdIpTunnelMtu,
			CmdIpTunnelNatDescriptor,
			CmdIpTunnelOspfArea,
			CmdIpTunnelOspfNeighbor,
			CmdIpTunnelPimSparse,
			CmdIpTunnelRebound,
			CmdIpTunnelRemoteAddress,
			CmdIpTunnelRipAuthKey,
			CmdIpTunnelRipAuthKeyText,
			CmdIpTunnelRipAuthType,
			CmdIpTunnelRipFilter,
			CmdIpTunnelRipForceToAdvertise,
			CmdIpTunnelRipHop,
			CmdIpTunnelRipHopDirectionHop,
			CmdIpTunnelRipReceive,
			CmdIpTunnelRipReceiveReceive,
			CmdIpTunnelRipSend,
			CmdIpTunnelRipTrustGateway,
			CmdIpTunnelSecureFilter,
			CmdIpTunnelSecureFilterName,
			CmdIpTunnelTcpMssLimit,
			CmdIpTunnelTcpWindowScale,
			CmdIpTunnelTrafficList,
			CmdIpTunnelTrafficListThreshold,
			CmdIpv6Filter,
			CmdIpv6FilterDynamic,
			CmdIpv6IcmpEchoReplySend,
			CmdIpv6IcmpEchoReplySendOnlyLinkup,
			CmdIpv6IcmpErrorDecryptedIpsecSend,
			CmdIpv6IcmpLog,
			CmdIpv6IcmpPacketTooBigForTruncatedSend,
			CmdIpv6IcmpPacketTooBigSend,
			CmdIpv6IcmpParameterProblemSend,
			CmdIpv6IcmpRedirectReceive,
			CmdIpv6IcmpRedirectSend,
			CmdIpv6IcmpTimeExceededSend,
			CmdIpv6IcmpUnreachableSend,
			CmdIpv6InboundFilter,
			CmdIpv6InterfaceAddress,
			CmdIpv6InterfaceDadRetryCount,
			CmdIpv6InterfaceDhcpService,
			CmdIpv6InterfaceInboundFilterList,
			CmdIpv6InterfaceMld,
			CmdIpv6InterfaceMldStatic,
			CmdIpv6InterfaceMtu,
			CmdIpv6InterfaceOspfArea,
			CmdIpv6InterfacePrefix,
			CmdIpv6InterfacePrefixChangeLog,
			CmdIpv6InterfaceRipFilter,
			CmdIpv6InterfaceRipHop,
			CmdIpv6InterfaceRipReceive,
			CmdIpv6InterfaceRipSend,
			CmdIpv6InterfaceRipTrustGateway,
			CmdIpv6InterfaceRtadvSend,
			CmdIpv6InterfaceSecureFilter,
			CmdIpv6InterfaceTcpMssLimit,
			CmdIpv6InterfaceTcpWindowScale,
			CmdIpv6InterfaceVrrp,
			CmdIpv6InterfaceVrrpShutdownTrigger,
			CmdIpv6MaxAutoAddress,
			CmdIpv6MulticastRoutingProcess,
			CmdIpv6NdNsTriggerDad,
			CmdIpv6OspfArea,
			CmdIpv6OspfAreaNetwork,
			CmdIpv6OspfConfigureRefresh,
			CmdIpv6OspfExport,
			CmdIpv6OspfExportFromOspf,
			CmdIpv6OspfImport,
			CmdIpv6OspfImportFrom,
			CmdIpv6OspfLog,
			CmdIpv6OspfPreference,
			CmdIpv6OspfRouterId,
			CmdIpv6OspfUse,
			CmdIpv6OspfVirtualLink,
			CmdIpv6PolicyAddressGroup,
			CmdIpv6PolicyFilter,
			CmdIpv6PolicyFilterSet,
			CmdIpv6PolicyFilterSetEnable,
			CmdIpv6PolicyFilterSetSwitch,
			CmdIpv6PolicyInterfaceGroup,
			CmdIpv6PolicyService,
			CmdIpv6PolicyServiceGroup,
			CmdIpv6PpAddress,
			CmdIpv6PpDadRetryCount,
			CmdIpv6PpDhcpService,
			CmdIpv6PpInboundFilterList,
			CmdIpv6PpMld,
			CmdIpv6PpMldStatic,
			CmdIpv6PpMtu,
			CmdIpv6PpOspfArea,
			CmdIpv6PpPrefix,
			CmdIpv6PpPrefixChangeLog,
			CmdIpv6PpRipConnectInterval,
			CmdIpv6PpRipConnectSend,
			CmdIpv6PpRipDisconnectInterval,
			CmdIpv6PpRipDisconnectSend,
			CmdIpv6PpRipFilter,
			CmdIpv6PpRipHoldRouting,
			CmdIpv6PpRipHop,
			CmdIpv6PpRipHopDirectionHop,
			CmdIpv6PpRipReceive,
			CmdIpv6PpRipSend,
			CmdIpv6PpRipTrustGateway,
			CmdIpv6PpRtadvSend,
			CmdIpv6PpSecureFilter,
			CmdIpv6PpTcpMssLimit,
			CmdIpv6PpTcpWindowScale,
			CmdIpv6Prefix,
			CmdIpv6Rh0Discard,
			CmdIpv6RipPreference,
			CmdIpv6RipUse,
			CmdIpv6Route,
			CmdIpv6Routing,
			CmdIpv6RoutingProcess,
			CmdIpv6SourceAddressSelectionRule,
			CmdIpv6Stealth,
			CmdIpv6StealthInterfaceInterface,
			CmdIpv6TunnelAddress,
			CmdIpv6TunnelDhcpService,
			CmdIpv6TunnelInboundFilterList,
			CmdIpv6TunnelMld,
			CmdIpv6TunnelMldStatic,
			CmdIpv6TunnelOspfArea,
			CmdIpv6TunnelPrefix,
			CmdIpv6TunnelPrefixChangeLog,
			CmdIpv6TunnelRipFilter,
			CmdIpv6TunnelRipReceive,
			CmdIpv6TunnelRipSend,
			CmdIpv6TunnelSecureFilter,
			CmdIpv6TunnelTcpMssLimit,
			CmdIpv6TunnelTcpWindowScale,
			CmdIsdnArrivePermit,
			CmdIsdnAutoConnect,
			CmdIsdnCallbackMscbcpUserSpecify,
			CmdIsdnCallbackPermit,
			CmdIsdnCallbackPermitType,
			CmdIsdnCallbackRequest,
			CmdIsdnCallbackRequestType,
			CmdIsdnCallbackResponseTime,
			CmdIsdnCallbackWaitTime,
			CmdIsdnCallBlockTime,
			CmdIsdnCallPermit,
			CmdIsdnCallProhibitTime,
			CmdIsdnDisconnectInputTime,
			CmdIsdnDisconnectIntervalTime,
			CmdIsdnDisconnectOutputTime,
			CmdIsdnDisconnectPolicy,
			CmdIsdnDisconnectTime,
			CmdIsdnDsu,
			CmdIsdnFastDisconnectTime,
			CmdIsdnForcedDisconnectTime,
			CmdIsdnLocalAddress,
			CmdIsdnPiafsArrive,
			CmdIsdnPiafsCall,
			CmdIsdnPiafsControl,
			CmdIsdnRemoteAddress,
			CmdIsdnRemoteCallOrder,
			CmdIsdnTerminator,
			CmdJateNumber,
			CmdL2TpAlwaysOn,
			CmdL2TpHostname,
			CmdL2TpKeepaliveLog,
			CmdL2TpKeepaliveUse,
			CmdL2TpLocalRouterId,
			CmdL2TpRemoteEndId,
			CmdL2TpRemoteRouterId,
			CmdL2TpService,
			CmdL2TpSyslog,
			CmdL2TpTunnelAuth,
			CmdL2TpTunnelDisconnectTime,
			CmdLanBackup,
			CmdLanBackupInterfacePp,
			CmdLanBackupInterfaceTunnel,
			CmdLanBackupRecoveryTime,
			CmdLanCountHubOverflow,
			CmdLanKeepaliveInterval,
			CmdLanKeepaliveLog,
			CmdLanKeepaliveUse,
			CmdLanLinkAggregationStatic,
			CmdLanLinkupSendWaitTime,
			CmdLanMapLog,
			CmdLanMapSnapshotUse,
			CmdLanMapSysname,
			CmdLanMapTerminalWatchInterval,
			CmdLanPortMirroring,
			CmdLanReceiveBufferSize,
			CmdLanShutdown,
			CmdLanType,
			CmdLeasedBackup,
			CmdLeasedKeepaliveDown,
			CmdLess,
			CmdLessConfig,
			CmdLessConfigAp,
			CmdLessConfigFilename,
			CmdLessConfigList,
			CmdLessConfigPp,
			CmdLessConfigSwitch,
			CmdLessConfigTunnel,
			CmdLessExecList,
			CmdLessFileList,
			CmdLessLog,
			CmdLineMasterclock,
			CmdLineType,
			CmdLoad,
			CmdLoginPassword,
			CmdLoginPasswordEncrypted,
			CmdLoginRadiusUse,
			CmdLoginTimer,
			CmdLoginUser,
			CmdLua,
			CmdLuac,
			CmdLuaUse,
			CmdMacro,
			CmdMailNotify,
			CmdMailNotifyStatusExec1,
			CmdMailNotifyStatusExec2,
			CmdMailNotifyStatusFrom,
			CmdMailNotifyStatusServer,
			CmdMailNotifyStatusSubject,
			CmdMailNotifyStatusTimeout,
			CmdMailNotifyStatusTo,
			CmdMailNotifyStatusType,
			CmdMailNotifyStatusUse,
			CmdMailServerName,
			CmdMailServerPop,
			CmdMailServerSmtp,
			CmdMailServerTimeout,
			CmdMailTemplate,
			CmdMakeDirectory,
			CmdMobileAccessLimitConnectionLength,
			CmdMobileAccessLimitConnectionTime,
			CmdMobileAccessLimitDuration,
			CmdMobileAccessLimitLength,
			CmdMobileAccessLimitTime,
			CmdMobileAccessPointName,
			CmdMobileArrivePermit,
			CmdMobileArriveUse,
			CmdMobileAutoConnect,
			CmdMobileCallProhibitAuthErrorCount,
			CmdMobileCarrierMode,
			CmdMobileDialNumber,
			CmdMobileDisconnectInputTime,
			CmdMobileDisconnectOutputTime,
			CmdMobileDisconnectTime,
			CmdMobileDisplayCallerId,
			CmdMobileFirmwareUpdateGo,
			CmdMobilePinCode,
			CmdMobileSignalStrength,
			CmdMobileSignalStrengthGo,
			CmdMobileSyslog,
			CmdMobileUse,
			CmdNatDescriptorAddressInner,
			CmdNatDescriptorAddressOuter,
			CmdNatDescriptorBackwardCompatibility,
			CmdNatDescriptorFtpPort,
			CmdNatDescriptorLog,
			CmdNatDescriptorMasqueradeIncoming,
			CmdNatDescriptorMasqueradePortRange,
			CmdNatDescriptorMasqueradeRemoveDfBit,
			CmdNatDescriptorMasqueradeRlogin,
			CmdNatDescriptorMasqueradeSessionLimit,
			CmdNatDescriptorMasqueradeSessionLimitTotal,
			CmdNatDescriptorMasqueradeStatic,
			CmdNatDescriptorMasqueradeTtlHold,
			CmdNatDescriptorMasqueradeUnconvertiblePort,
			CmdNatDescriptorSip,
			CmdNatDescriptorStatic,
			CmdNatDescriptorTimer,
			CmdNatDescriptorType,
			CmdNetvolanteDnsAutoHostname,
			CmdNetvolanteDnsAutoHostnamePp,
			CmdNetvolanteDnsAutoSave,
			CmdNetvolanteDnsDeleteGo,
			CmdNetvolanteDnsDeleteGoPp,
			CmdNetvolanteDnsGetHostnameList,
			CmdNetvolanteDnsGetHostnameListPp,
			CmdNetvolanteDnsGo,
			CmdNetvolanteDnsGoPp,
			CmdNetvolanteDnsHostnameHost,
			CmdNetvolanteDnsHostnameHostPp,
			CmdNetvolanteDnsPort,
			CmdNetvolanteDnsRegisterTimer,
			CmdNetvolanteDnsRetryInterval,
			CmdNetvolanteDnsRetryIntervalPp,
			CmdNetvolanteDnsServer,
			CmdNetvolanteDnsServerUpdateAddressPort,
			CmdNetvolanteDnsServerUpdateAddressUse,
			CmdNetvolanteDnsSetHostname,
			CmdNetvolanteDnsTimeout,
			CmdNetvolanteDnsTimeoutPp,
			CmdNetvolanteDnsUse,
			CmdNetvolanteDnsUsePp,
			CmdNgnRadiusAccountCallee,
			CmdNgnRadiusAccountCaller,
			CmdNgnRadiusAuthPassword,
			CmdNgnRenumberingLinkRefresh,
			CmdNgnType,
			CmdNoAccountThreshold,
			CmdNoAdministratorRadiusAuth,
			CmdNoAlarmBatch,
			CmdNoAlarmEntire,
			CmdNoAlarmHttpRevisionUpSwitch,
			CmdNoAlarmHttpUpload,
			CmdNoAlarmLua,
			CmdNoAlarmMobile,
			CmdNoAlarmSdSwitch,
			CmdNoAlarmStartup,
			CmdNoAlarmUsbhost,
			CmdNoAlias,
			CmdNoApConfigDirectory,
			CmdNoApConfigFilename,
			CmdNoApControlConfigAutoSetUse,
			CmdNoApControlHttpProxyTimeout,
			CmdNoApControlHttpProxyUse,
			CmdNoApSelect,
			CmdNoAuthUser,
			CmdNoAuthUserAttribute,
			CmdNoAuthUserGroup,
			CmdNoAuthUserGroupAttribute,
			CmdNoBgpAggregate,
			CmdNoBgpAggregateFilter,
			CmdNoBgpAutonomousSystem,
			CmdNoBgpExport,
			CmdNoBgpExportFilter,
			CmdNoBgpForceToAdvertise,
			CmdNoBgpImport,
			CmdNoBgpImportFilter,
			CmdNoBgpLog,
			CmdNoBgpNeighbor,
			CmdNoBgpNeighborPreSharedKey,
			CmdNoBgpPreference,
			CmdNoBgpReric,
			CmdNoBgpRouterIdIp_Address,
			CmdNoBgpUse,
			CmdNoBridgeLearningSwitch,
			CmdNoBridgeLearningStatic,
			CmdNoBridgeLearningBridgeInterfaceTimer,
			CmdNoBridgeMember,
			CmdNoCloudVpnBind,
			CmdNoCloudVpnName,
			CmdNoCloudVpnOption,
			CmdNoCloudVpnParameter,
			CmdNoCloudVpnService,
			CmdNoConsoleColumns,
			CmdNoConsoleInfo,
			CmdNoConsoleLines,
			CmdNoCooperation,
			CmdNoCooperationBandwidthMeasuringRemote,
			CmdNoCooperationLoadWatchControl,
			CmdNoCooperationLoadWatchRemote,
			CmdNoCooperationLoadWatchTrigger,
			CmdNoCooperationPort,
			CmdNoDashboardAccumulate,
			CmdNoDescription,
			CmdNoDescriptionInterfaceDescription,
			CmdNoDescriptionYno,
			CmdNoDhcpClientClientIdentifier,
			CmdNoDhcpClientHostname,
			CmdNoDhcpClientOptionInterfaceOptionValue,
			CmdNoDhcpDuplicateCheck,
			CmdNoDhcpRelaySelect,
			CmdNoDhcpRelayServer,
			CmdNoDhcpRelayThreshold,
			CmdNoDhcpScope,
			CmdNoDhcpScopeBind,
			CmdNoDhcpScopeLeaseType,
			CmdNoDhcpScopeOption,
			CmdNoDhcpServerRfc2131Compliant,
			CmdNoDhcpServiceType,
			CmdNoDnsCacheMaxEntry,
			CmdNoDnsCacheUse,
			CmdNoDnsDomain,
			CmdNoDnsHost,
			CmdNoDnsNoticeOrder,
			CmdNoDnsPrivateAddressSpoof,
			CmdNoDnsServer,
			CmdNoDnsServerDhcp,
			CmdNoDnsServerPp,
			CmdNoDnsServerSelect,
			CmdNoDnsService,
			CmdNoDnsServiceAaaaFilter,
			CmdNoDnsServiceFallback,
			CmdNoDnsSrcport,
			CmdNoDnsStaticTypeNameValue,
			CmdNoDnsSyslogResolv,
			CmdNoEmbeddedFile,
			CmdNoEthernetInterfaceFilter,
			CmdNoExternalMemoryAcceleratorCacheSize,
			CmdNoExternalMemoryAutoSearchTime,
			CmdNoExternalMemoryBatchFilename,
			CmdNoExternalMemoryBootPermit,
			CmdNoExternalMemoryBootTimeout,
			CmdNoExternalMemoryCacheMode,
			CmdNoExternalMemoryConfigFilename,
			CmdNoExternalMemoryExecFilename,
			CmdNoExternalMemoryStatisticsFilenamePrefix,
			CmdNoExternalMemorySyslogFilename,
			CmdNoFrBackup,
			CmdNoFrCir,
			CmdNoFrCompressionUse,
			CmdNoFrCongestionControl,
			CmdNoFrDe,
			CmdNoFrDlci,
			CmdNoFrInarp,
			CmdNoFrLmi,
			CmdNoFrPpDequeueType,
			CmdNoHeartbeat2Myname,
			CmdNoHeartbeat2Receive,
			CmdNoHeartbeat2ReceiveEnable,
			CmdNoHeartbeat2ReceiveLog,
			CmdNoHeartbeat2ReceiveMonitor,
			CmdNoHeartbeat2ReceiveRecordLimit,
			CmdNoHeartbeat2Transmit,
			CmdNoHeartbeat2TransmitEnable,
			CmdNoHeartbeat2TransmitInterval,
			CmdNoHeartbeat2TransmitLog,
			CmdNoHeartbeatPreSharedKey,
			CmdNoHeartbeatReceive,
			CmdNoHttpdCustomGuiApiPassword,
			CmdNoHttpdCustomGuiApiUse,
			CmdNoHttpdCustomGuiUse,
			CmdNoHttpdCustomGuiUser,
			CmdNoHttpdHost,
			CmdNoHttpdListen,
			CmdNoHttpdProxyAccessL2MsPermit,
			CmdNoHttpdService,
			CmdNoHttpdTimeout,
			CmdNoHttpRevisionDownPermit,
			CmdNoHttpRevisionUpPermit,
			CmdNoHttpRevisionUpProxy,
			CmdNoHttpRevisionUpSchedule,
			CmdNoHttpRevisionUpTimeout,
			CmdNoHttpRevisionUpUrl,
			CmdNoHttpUpload,
			CmdNoHttpUploadPermit,
			CmdNoHttpUploadProxy,
			CmdNoHttpUploadRetryInterval,
			CmdNoHttpUploadTimeout,
			CmdNoHttpUploadUrl,
			CmdNoIpArpTimerTimerRetry,
			CmdNoIpFilterDirectedBroadcast,
			CmdNoIpFilterDynamic,
			CmdNoIpFilterDynamicTimer,
			CmdNoIpFilterFilter_NumPass_Reject,
			CmdNoIpFilterFqdnTimerTime,
			CmdNoIpFilterSetNameDirection,
			CmdNoIpFilterSourceRoute,
			CmdNoIpFlowLimit,
			CmdNoIpFlowTimer,
			CmdNoIpForwardFilter,
			CmdNoIpFragmentRemove,
			CmdNoIpHostFqdnValue,
			CmdNoIpIcmpEchoReplySend,
			CmdNoIpIcmpEchoReplySendOnlyLinkup,
			CmdNoIpIcmpErrorDecryptedIpsecSend,
			CmdNoIpIcmpLog,
			CmdNoIpIcmpMaskReplySend,
			CmdNoIpIcmpParameterProblemSend,
			CmdNoIpIcmpRedirectReceive,
			CmdNoIpIcmpRedirectSend,
			CmdNoIpIcmpTimeExceededSendSend,
			CmdNoIpIcmpTimestampReplySend,
			CmdNoIpIcmpUnreachableForTruncatedSend,
			CmdNoIpIcmpUnreachableSend,
			CmdNoIpImplicitRoutePreferencePreference,
			CmdNoIpInboundFilter,
			CmdNoIpInterfaceArpLog,
			CmdNoIpInterfaceArpMtuDiscovery,
			CmdNoIpInterfaceArpQueueLength,
			CmdNoIpInterfaceArpStatic,
			CmdNoIpInterfaceDhcpAutoDefaultRouteAdd,
			CmdNoIpInterfaceDhcpAutoInterfaceRouteAdd,
			CmdNoIpInterfaceDhcpLeaseTime,
			CmdNoIpInterfaceDhcpRetryRetry,
			CmdNoIpInterfaceDhcpService,
			CmdNoIpInterfaceForwardFilter,
			CmdNoIpInterfaceIgmp,
			CmdNoIpInterfaceIgmpStatic,
			CmdNoIpInterfaceInboundFilterList,
			CmdNoIpInterfaceIntrusionDetection,
			CmdNoIpInterfaceIntrusionDetectionNotice,
			CmdNoIpInterfaceIntrusionDetectionRepeatControl,
			CmdNoIpInterfaceIntrusionDetectionReport,
			CmdNoIpInterfaceIntrusionDetectionThreshold,
			CmdNoIpInterfaceMtu,
			CmdNoIpInterfaceNatDescriptor,
			CmdNoIpInterfaceOspfArea,
			CmdNoIpInterfaceOspfNeighbor,
			CmdNoIpInterfacePimSparse,
			CmdNoIpInterfaceProxyarp,
			CmdNoIpInterfaceRebound,
			CmdNoIpInterfaceRipAuthKey,
			CmdNoIpInterfaceRipAuthKeyText,
			CmdNoIpInterfaceRipAuthType,
			CmdNoIpInterfaceRipFilter,
			CmdNoIpInterfaceRipForceToAdvertise,
			CmdNoIpInterfaceRipHopDirectionHop,
			CmdNoIpInterfaceRipReceive,
			CmdNoIpInterfaceRipSend,
			CmdNoIpInterfaceRipTrustGateway,
			CmdNoIpInterfaceSecondaryAddress,
			CmdNoIpInterfaceSecureFilter,
			CmdNoIpInterfaceSecureFilterName,
			CmdNoIpInterfaceTcpMssLimit,
			CmdNoIpInterfaceTcpWindowScale,
			CmdNoIpInterfaceTrafficList,
			CmdNoIpInterfaceTrafficListThreshold,
			CmdNoIpInterfaceVrrp,
			CmdNoIpInterfaceVrrpShutdownTrigger,
			CmdNoIpInterfaceWolRelay,
			CmdNoIpipKeepaliveLog,
			CmdNoIpipKeepaliveUse,
			CmdNoIpKeepalive,
			CmdNoIpLocalForwardFilter,
			CmdNoIpPimSparseJoinPruneSend,
			CmdNoIpPimSparseLog,
			CmdNoIpPimSparsePeriodicPruneSend,
			CmdNoIpPimSparseRegisterChecksum,
			CmdNoIpPimSparseRendezvousPointStatic,
			CmdNoIpPolicyAddressGroup,
			CmdNoIpPolicyFilter,
			CmdNoIpPolicyFilterSet,
			CmdNoIpPolicyFilterSetEnable,
			CmdNoIpPolicyFilterTimer,
			CmdNoIpPolicyInterfaceGroup,
			CmdNoIpPolicyService,
			CmdNoIpPolicyServiceGroup,
			CmdNoIpPpForwardFilter,
			CmdNoIpPpIgmp,
			CmdNoIpPpIgmpStatic,
			CmdNoIpPpInboundFilterListId,
			CmdNoIpPpIntrusionDetection,
			CmdNoIpPpIntrusionDetectionNotice,
			CmdNoIpPpIntrusionDetectionRepeatControl,
			CmdNoIpPpIntrusionDetectionReport,
			CmdNoIpPpIntrusionDetectionThreshold,
			CmdNoIpPpMtu,
			CmdNoIpPpNatDescriptor,
			CmdNoIpPpOspfArea,
			CmdNoIpPpOspfNeighbor,
			CmdNoIpPpPimSparse,
			CmdNoIpPpRebound,
			CmdNoIpPpRemoteAddress,
			CmdNoIpPpRemoteAddressPool,
			CmdNoIpPpRipAuthKey,
			CmdNoIpPpRipAuthKeyText,
			CmdNoIpPpRipAuthType,
			CmdNoIpPpRipBackupInterface,
			CmdNoIpPpRipConnectInterval,
			CmdNoIpPpRipConnectSend,
			CmdNoIpPpRipDisconnect,
			CmdNoIpPpRipDisconnectInterval,
			CmdNoIpPpRipFilter,
			CmdNoIpPpRipForceToAdvertise,
			CmdNoIpPpRipHoldRouting,
			CmdNoIpPpRipHopDirectionHop,
			CmdNoIpPpRipReceive,
			CmdNoIpPpRipSend,
			CmdNoIpPpRipTrustGateway,
			CmdNoIpPpSecureFilter,
			CmdNoIpPpSecureFilterName,
			CmdNoIpPpTcpMssLimit,
			CmdNoIpPpTcpWindowScale,
			CmdNoIpPpTrafficList,
			CmdNoIpPpTrafficListThreshold,
			CmdNoIpRoute,
			CmdNoIpRouteChangeLog,
			CmdNoIpRouting,
			CmdNoIpRoutingProcess,
			CmdNoIpsecAutoRefresh,
			CmdNoIpsecIkeAlwaysOn,
			CmdNoIpsecIkeAuthMethod,
			CmdNoIpsecIkeBackwardCompatibility,
			CmdNoIpsecIkeChildExchangeType,
			CmdNoIpsecIkeDuration,
			CmdNoIpsecIkeEapMyname,
			CmdNoIpsecIkeEapRequest,
			CmdNoIpsecIkeEapSendCertreq,
			CmdNoIpsecIkeEncryption,
			CmdNoIpsecIkeEspEncapsulation,
			CmdNoIpsecIkeGroup,
			CmdNoIpsecIkeHash,
			CmdNoIpsecIkeKeepaliveLog,
			CmdNoIpsecIkeKeepaliveUse,
			CmdNoIpsecIkeLicenseKey,
			CmdNoIpsecIkeLicenseKeyUse,
			CmdNoIpsecIkeLocalAddress,
			CmdNoIpsecIkeLocalIdGateway_IdIp_AddressMask,
			CmdNoIpsecIkeLocalName,
			CmdNoIpsecIkeLog,
			CmdNoIpsecIkeMessageIdControl,
			CmdNoIpsecIkeModeCfgAddress,
			CmdNoIpsecIkeModeCfgAddressPool,
			CmdNoIpsecIkeModeCfgMethod,
			CmdNoIpsecIkeNatTraversal,
			CmdNoIpsecIkeNegotiateStrictly,
			CmdNoIpsecIkeNegotiationReceive,
			CmdNoIpsecIkePayloadType,
			CmdNoIpsecIkePfs,
			CmdNoIpsecIkePkiFile,
			CmdNoIpsecIkePreSharedKey,
			CmdNoIpsecIkeProposalLimitation,
			CmdNoIpsecIkeQueueLength,
			CmdNoIpsecIkeRemoteAddress,
			CmdNoIpsecIkeRemoteId,
			CmdNoIpsecIkeRemoteName,
			CmdNoIpsecIkeRestrictDanglingSa,
			CmdNoIpsecIkeRetry,
			CmdNoIpsecIkeSendInfo,
			CmdNoIpsecIKeVersion,
			CmdNoIpsecIkeXauthMyname,
			CmdNoIpsecIkeXauthRequest,
			CmdNoIpsecIpcompType,
			CmdNoIpsecLogIllegalSpi,
			CmdNoIpsecSaPolicy,
			CmdNoIpsecTransport,
			CmdNoIpsecTransportTemplate,
			CmdNoIpsecTunnel,
			CmdNoIpsecTunnelFastpathFragmentFunctionFollowDfBit,
			CmdNoIpsecTunnelOuterDfBit,
			CmdNoIpsecUse,
			CmdNoIpSimpleService,
			CmdNoIpStealth,
			CmdNoIpTosSupersede,
			CmdNoIpTunnelAddress,
			CmdNoIpTunnelForwardFilter,
			CmdNoIpTunnelIgmp,
			CmdNoIpTunnelIgmpStatic,
			CmdNoIpTunnelInboundFilterListId,
			CmdNoIpTunnelIntrusionDetection,
			CmdNoIpTunnelIntrusionDetectionNotice,
			CmdNoIpTunnelIntrusionDetectionRepeatControl,
			CmdNoIpTunnelIntrusionDetectionReport,
			CmdNoIpTunnelIntrusionDetectionThreshold,
			CmdNoIpTunnelMtu,
			CmdNoIpTunnelNatDescriptor,
			CmdNoIpTunnelOspfArea,
			CmdNoIpTunnelOspfNeighbor,
			CmdNoIpTunnelPimSparse,
			CmdNoIpTunnelRebound,
			CmdNoIpTunnelRemoteAddress,
			CmdNoIpTunnelRipAuthKey,
			CmdNoIpTunnelRipAuthKeyText,
			CmdNoIpTunnelRipAuthType,
			CmdNoIpTunnelRipFilter,
			CmdNoIpTunnelRipForceToAdvertise,
			CmdNoIpTunnelRipHopDirectionHop,
			CmdNoIpTunnelRipReceive,
			CmdNoIpTunnelRipSend,
			CmdNoIpTunnelRipTrustGateway,
			CmdNoIpTunnelSecureFilter,
			CmdNoIpTunnelSecureFilterName,
			CmdNoIpTunnelTcpMssLimit,
			CmdNoIpTunnelTcpWindowScale,
			CmdNoIpTunnelTrafficList,
			CmdNoIpTunnelTrafficListThreshold,
			CmdNoIpv6Filter,
			CmdNoIpv6FilterDynamic,
			CmdNoIpv6IcmpEchoReplySendOnlyLinkup,
			CmdNoIpv6IcmpEchoReplySendSend,
			CmdNoIpv6IcmpErrorDecryptedIpsecSend,
			CmdNoIpv6IcmpLog,
			CmdNoIpv6IcmpPacketTooBigForTruncatedSend,
			CmdNoIpv6IcmpPacketTooBigSend,
			CmdNoIpv6IcmpParameterProblemSend,
			CmdNoIpv6IcmpRedirectReceive,
			CmdNoIpv6IcmpRedirectSend,
			CmdNoIpv6IcmpTimeExceededSend,
			CmdNoIpv6IcmpUnreachableSend,
			CmdNoIpv6InboundFilter,
			CmdNoIpv6InterfaceAddress,
			CmdNoIpv6InterfaceDadRetryCount,
			CmdNoIpv6InterfaceDhcpService,
			CmdNoIpv6InterfaceInboundFilterList,
			CmdNoIpv6InterfaceMld,
			CmdNoIpv6InterfaceMldStatic,
			CmdNoIpv6InterfaceMtu,
			CmdNoIpv6InterfaceOspfArea,
			CmdNoIpv6InterfacePrefix,
			CmdNoIpv6InterfacePrefixChangeLog,
			CmdNoIpv6InterfaceRipFilter,
			CmdNoIpv6InterfaceRipHop,
			CmdNoIpv6InterfaceRipReceive,
			CmdNoIpv6InterfaceRipSend,
			CmdNoIpv6InterfaceRipTrustGateway,
			CmdNoIpv6InterfaceRtadvSend,
			CmdNoIpv6InterfaceSecureFilter,
			CmdNoIpv6InterfaceTcpMssLimit,
			CmdNoIpv6InterfaceTcpWindowScale,
			CmdNoIpv6InterfaceVrrp,
			CmdNoIpv6InterfaceVrrpShutdownTrigger,
			CmdNoIpv6MaxAutoAddress,
			CmdNoIpv6MulticastRoutingProcess,
			CmdNoIpv6NdNsTriggerDad,
			CmdNoIpv6OspfArea,
			CmdNoIpv6OspfAreaNetwork,
			CmdNoIpv6OspfExport,
			CmdNoIpv6OspfExportFromOspf,
			CmdNoIpv6OspfImport,
			CmdNoIpv6OspfImportFrom,
			CmdNoIpv6OspfLog,
			CmdNoIpv6OspfPreference,
			CmdNoIpv6OspfRouterId,
			CmdNoIpv6OspfUse,
			CmdNoIpv6OspfVirtualLink,
			CmdNoIpv6PolicyAddressGroup,
			CmdNoIpv6PolicyFilter,
			CmdNoIpv6PolicyFilterSetEnable,
			CmdNoIpv6PolicyInterfaceGroup,
			CmdNoIpv6PolicyService,
			CmdNoIpv6PolicyServiceGroup,
			CmdNoIpv6PpAddress,
			CmdNoIpv6PpDadRetryCount,
			CmdNoIpv6PpDhcpService,
			CmdNoIpv6PpInboundFilterListId,
			CmdNoIpv6PpMld,
			CmdNoIpv6PpMtu,
			CmdNoIpv6PpOspfArea,
			CmdNoIpv6PpPrefix,
			CmdNoIpv6PpPrefixChangeLog,
			CmdNoIpv6PpRipConnectInterval,
			CmdNoIpv6PpRipConnectSend,
			CmdNoIpv6PpRipDisconnectInterval,
			CmdNoIpv6PpRipDisconnectSend,
			CmdNoIpv6PpRipFilter,
			CmdNoIpv6PpRipHoldRouting,
			CmdNoIpv6PpRipHop,
			CmdNoIpv6PpRipReceive,
			CmdNoIpv6PpRipSend,
			CmdNoIpv6PpRipTrustGateway,
			CmdNoIpv6PpRtadvSend,
			CmdNoIpv6PpSecureFilter,
			CmdNoIpv6PpTcpMssLimit,
			CmdNoIpv6PpTcpWindowScale,
			CmdNoIpv6Prefix,
			CmdNoIpv6Rh0Discard,
			CmdNoIpv6RipPreference,
			CmdNoIpv6RipUse,
			CmdNoIpv6Route,
			CmdNoIpv6Routing,
			CmdNoIpv6RoutingProcess,
			CmdNoIpv6SourceAddressSelectionRule,
			CmdNoIpv6Stealth,
			CmdNoIpv6TunnelAddress,
			CmdNoIpv6TunnelDhcpService,
			CmdNoIpv6TunnelInboundFilterListId,
			CmdNoIpv6TunnelMld,
			CmdNoIpv6TunnelMldStatic,
			CmdNoIpv6TunnelOspfArea,
			CmdNoIpv6TunnelPrefix,
			CmdNoIpv6TunnelPrefixChangeLog,
			CmdNoIpv6TunnelRipFilter,
			CmdNoIpv6TunnelRipReceive,
			CmdNoIpv6TunnelRipSend,
			CmdNoIpv6TunnelSecureFilter,
			CmdNoIpv6TunnelTcpMssLimit,
			CmdNoIpv6TunnelTcpWindowScale,
			CmdNoIsdnArrivePermit,
			CmdNoIsdnAutoConnect,
			CmdNoIsdnCallbackPermit,
			CmdNoIsdnCallbackPermitType,
			CmdNoIsdnCallbackRequest,
			CmdNoIsdnCallbackRequestType,
			CmdNoIsdnCallbackResponseTime,
			CmdNoIsdnCallbackWaitTime,
			CmdNoIsdnCallBlockTime,
			CmdNoIsdnCallPermit,
			CmdNoIsdnCallProhibitTime,
			CmdNoIsdnDisconnectInputTime,
			CmdNoIsdnDisconnectIntervalTime,
			CmdNoIsdnDisconnectOutputTime,
			CmdNoIsdnDisconnectPolicy,
			CmdNoIsdnDisconnectTime,
			CmdNoIsdnDsu,
			CmdNoIsdnForcedDisconnectTime,
			CmdNoisdnLocalAddress,
			CmdNoIsdnPiafsArrive,
			CmdNoIsdnPiafsCall,
			CmdNoIsdnPiafsControl,
			CmdNoIsdnRemoteAddress,
			CmdNoIsdnRemoteCallOrder,
			CmdNoIsdnTerminator,
			CmdNoJateNumber,
			CmdNoL2TpAlwaysOn,
			CmdNoL2TpHostname,
			CmdNoL2TpKeepaliveLog,
			CmdNoL2TpKeepaliveUse,
			CmdNoL2TpLocalRouterId,
			CmdNoL2TpRemoteEndId,
			CmdNoL2TpRemoteRouterId,
			CmdNoL2TpService,
			CmdNoL2TpSyslog,
			CmdNoL2TpTunnelAuth,
			CmdNoL2TpTunnelDisconnectTimeTime,
			CmdNoLanBackup,
			CmdNoLanBackupRecoveryTime,
			CmdNoLanCountHubOverflow,
			CmdNoLanKeepaliveInterval,
			CmdNoLanKeepaliveLog,
			CmdNoLanKeepaliveUse,
			CmdNoLanLinkAggregationStatic,
			CmdNoLanLinkupSendWaitTime,
			CmdNoLanMapLog,
			CmdNoLanMapSnapshotUse,
			CmdNoLanMapSysname,
			CmdNoLanMapTerminalWatchInterval,
			CmdNoLanPortMirroring,
			CmdNoLanReceiveBufferSize,
			CmdNoLanShutdown,
			CmdNoLanType,
			CmdNoLeasedBackup,
			CmdNoLeasedKeepaliveDown,
			CmdNoLineMasterclock,
			CmdNoLineType,
			CmdNoLoginRadiusUse,
			CmdNoLoginTimer,
			CmdNoLoginUser,
			CmdNoLuaUse,
			CmdNoMacroNameEom,
			CmdNoMailNotify,
			CmdNoMailNotifyStatusFrom,
			CmdNoMailNotifyStatusServer,
			CmdNoMailNotifyStatusSubject,
			CmdNoMailNotifyStatusTimeout,
			CmdNoMailNotifyStatusToId,
			CmdNoMailNotifyStatusType,
			CmdNoMailNotifyStatusUse,
			CmdNoMailServerName,
			CmdNoMailServerPop,
			CmdNoMailServerSmtp,
			CmdNoMailServerTimeout,
			CmdNoMailTemplate,
			CmdNoMobileAccessLimitConnectionLength,
			CmdNoMobileAccessLimitConnectionTime,
			CmdNoMobileAccessLimitDuration,
			CmdNoMobileAccessLimitLength,
			CmdNoMobileAccessLimitTimeTime,
			CmdNoMobileAccessPointName,
			CmdNoMobileArrivePermit,
			CmdNoMobileArriveUse,
			CmdNoMobileAutoConnect,
			CmdNoMobileCallProhibitAuthErrorCount,
			CmdNoMobileCarrierMode,
			CmdNoMobileDialNumberDial_String,
			CmdNoMobileDisconnectInputTime,
			CmdNoMobileDisconnectOutputTime,
			CmdNoMobileDisconnectTime,
			CmdNoMobileDisplayCallerId,
			CmdNoMobilePinCode,
			CmdNoMobileSignalStrength,
			CmdNoMobileSyslog,
			CmdNoMobileUse,
			CmdNoNatDescriptorAddressInner,
			CmdNoNatDescriptorAddressOuter,
			CmdNoNatDescriptorBackwardCompatibility,
			CmdNoNatDescriptorFtpPort,
			CmdNoNatDescriptorLog,
			CmdNoNatDescriptorMasqueradeIncoming,
			CmdNoNatDescriptorMasqueradePortRange,
			CmdNoNatDescriptorMasqueradeRemoveDfBit,
			CmdNoNatDescriptorMasqueradeRlogin,
			CmdNoNatDescriptorMasqueradeSessionLimit,
			CmdNoNatDescriptorMasqueradeSessionLimitTotal,
			CmdNoNatDescriptorMasqueradeStatic,
			CmdNoNatDescriptorMasqueradeTtlHold,
			CmdNoNatDescriptorMasqueradeUnconvertiblePort,
			CmdNoNatDescriptorSip,
			CmdNoNatDescriptorStatic,
			CmdNoNatDescriptorTimer,
			CmdNoNatDescriptorType,
			CmdNoNetvolanteDnsAutoHostname,
			CmdNoNetvolanteDnsAutoHostnamePp,
			CmdNoNetvolanteDnsAutoSave,
			CmdNoNetvolanteDnsHostnameHostInterfaceHost,
			CmdNoNetvolanteDnsHostnameHostPpHost,
			CmdNoNetvolanteDnsPort,
			CmdNoNetvolanteDnsRegisterTimer,
			CmdNoNetvolanteDnsRetryInterval,
			CmdNoNetvolanteDnsRetryIntervalPpIntervalCount,
			CmdNoNetvolanteDnsServer,
			CmdNoNetvolanteDnsServerUpdateAddressPort,
			CmdNoNetvolanteDnsServerUpdateAddressUse,
			CmdNoNetvolanteDnsTimeout,
			CmdNoNetvolanteDnsTimeoutPp,
			CmdNoNetvolanteDnsUse,
			CmdNoNetvolanteDnsUsePp,
			CmdNoNgnRadiusAccountCallee,
			CmdNoNgnRadiusAccountCaller,
			CmdNoNgnRadiusAuthPassword,
			CmdNoNgnRenumberingLinkRefresh,
			CmdNoNgnType,
			CmdNoNoIsdnCallbackMscbcpUserSpecify,
			CmdNoNoIsdnFastDisconnectTime,
			CmdNoNtpBackwardCompatibility,
			CmdNoNtpLocalAddress,
			CmdNoOperationButtonFunctionDownload,
			CmdNoOperationExecuteBatchPermit,
			CmdNoOperationExternalMemoryDownloadPermit,
			CmdNoOperationHttpRevisionUpPermit,
			CmdNoOperationUsbDownloadPermit,
			CmdNoOspfArea,
			CmdNoOspfAreaNetwork,
			CmdNoOspfAreaStubhost,
			CmdNoOspfExportFilter,
			CmdNoOspfExportFromOspf,
			CmdNoOspfImportFilter,
			CmdNoOspfImportFrom,
			CmdNoOspfLog,
			CmdNoOspfMergeEqualCostStub,
			CmdNoOspfPreference,
			CmdNoOspfRericInterval,
			CmdNoOspfRouterId,
			CmdNoOspfUse,
			CmdNoOspfVirtualLink,
			CmdNoPkiCertificateFile,
			CmdNoPkiCrlFile,
			CmdNoPpAlwaysOn,
			CmdNoPpAuthAcceptAccept,
			CmdNoPpAuthMultiConnectProhibit,
			CmdNoPpAuthMyname,
			CmdNoPpAuthRequest,
			CmdNoPpAuthUsername,
			CmdNoPpBackup,
			CmdNoPpBackupRecoveryTime,
			CmdNoPpBind,
			CmdNoPpEnable,
			CmdNoPpEncapsulation,
			CmdNoPpKeepaliveInterval,
			CmdNoPpKeepaliveLog,
			CmdNoPpKeepaliveUse,
			CmdNoPpName,
			CmdNoPppBacpMaxconfigure,
			CmdNoPppBacpMaxfailure,
			CmdNoPppBacpMaxterminate,
			CmdNoPppBacpRestart,
			CmdNoPppBapMaxretry,
			CmdNoPppBapRestart,
			CmdNoPppCcpMaxconfigure,
			CmdNoPppCcpMaxfailure,
			CmdNoPppCcpMaxterminate,
			CmdNoPppCcpNoEncryption,
			CmdNoPppCcpRestart,
			CmdNoPppCcpType,
			CmdNoPppChapMaxchallenge,
			CmdNoPppChapRestart,
			CmdNoPppIpcpIpaddress,
			CmdNoPppIpcpMaxconfigure,
			CmdNoPppIpcpMaxfailure,
			CmdNoPppIpcpMaxterminate,
			CmdNoPppIpcpMsext,
			CmdNoPppIpcpRemoteAddressCheck,
			CmdNoPppIpcpRestart,
			CmdNoPppIpcpVjc,
			CmdNoPppIpv6CpUse,
			CmdNoPppLcpAccm,
			CmdNoPppLcpAcfc,
			CmdNoPppLcpMagicnumber,
			CmdNoPppLcpMaxconfigureCount,
			CmdNoPppLcpMaxfailure,
			CmdNoPppLcpMaxterminate,
			CmdNoPppLcpMruMruLength,
			CmdNoPppLcpPfc,
			CmdNoPppLcpRestart,
			CmdNoPppLcpSilent,
			CmdNoPppMpControl,
			CmdNoPppMpDivide,
			CmdNoPppMpInterleave,
			CmdNoPppMpLoadThreshold,
			CmdNoPppMpMaxlink,
			CmdNoPppMpMinlink,
			CmdNoPppMpTimer,
			CmdNoPppMpUse,
			CmdNoPppMscbcpMaxretry,
			CmdNoPppMscbcpRestart,
			CmdNoPppoeAccessConcentrator,
			CmdNoPppoeAutoConnect,
			CmdNoPppoeAutoDisconnect,
			CmdNoPppoeDisconnectTime,
			CmdNoPppoeInvalidSessionForcedClose,
			CmdNoPppoePadiMaxretry,
			CmdNoPppoePadiRestart,
			CmdNoPppoePadrMaxretry,
			CmdNoPppoePadrRestart,
			CmdNoPppoePassThroughMember,
			CmdNoPppoeServiceName,
			CmdNoPppoeTcpMssLimit,
			CmdNoPppoeUse,
			CmdNoPppPapMaxauthreq,
			CmdNoPppPapRestart,
			CmdNoPPSelect,
			CmdNoPptpCallIdMode,
			CmdNoPptpHostname,
			CmdNoPptpKeepaliveInterval,
			CmdNoPptpKeepaliveLog,
			CmdNoPptpKeepaliveUse,
			CmdNoPptpService,
			CmdNoPptpServiceType,
			CmdNoPptpSyslog,
			CmdNoPptpTunnelDisconnectTime,
			CmdNoPptpVendorname,
			CmdNoPptpWindowSize,
			CmdNoPriLeasedChannel,
			CmdNoProviderAutoConnectForcedDisable,
			CmdNoProviderDnsServer,
			CmdNoProviderDnsServerPp,
			CmdNoProviderFilterRouting,
			CmdNoProviderInterfaceBind,
			CmdNoProviderInterfaceDnsServer,
			CmdNoProviderInterfaceName,
			CmdNoProviderIpv6ConnectPp,
			CmdNoProviderNtpdate,
			CmdNoProviderNtpServer,
			CmdNoProviderSelect,
			CmdNoProviderSet,
			CmdNoProviderType,
			CmdNoPv6PpMldStatic,
			CmdNoQacTmClientPermit,
			CmdNoQacTmClientPort,
			CmdNoQacTmClientUpdate,
			CmdNoQacTmPort,
			CmdNoQacTmRedirect,
			CmdNoQacTmServer,
			CmdNoQacTmUnqualifiedClientAccessControl,
			CmdNoQacTmUse,
			CmdNoQacTmVersionMargin,
			CmdNoQacTmWarningUrl,
			CmdNoQueueClassFilter,
			CmdNoQueueInterfaceClassControl,
			CmdNoQueueInterfaceClassFilterList,
			CmdNoQueueInterfaceClassProperty,
			CmdNoQueueInterfaceDefaultClass,
			CmdNoQueueInterfaceDefaultClassSecondary,
			CmdNoQueueInterfaceLength,
			CmdNoQueueInterfaceLengthSecondary,
			CmdNoQueueInterfaceType,
			CmdNoQueuePpClassFilterList,
			CmdNoQueuePpClassProperty,
			CmdNoQueuePpDefaultClass,
			CmdNoQueuePpLength,
			CmdNoQueuePpTypeType,
			CmdNoQueueTunnelClassFilterList,
			CmdNoRadiusAccount,
			CmdNoRadiusAccountPort,
			CmdNoRadiusAccountServer,
			CmdNoRadiusAuth,
			CmdNoRadiusAuthPort,
			CmdNoRadiusAuthServer,
			CmdNoRadiusRetry,
			CmdNoRadiusSecret,
			CmdNoRadiusServer,
			CmdNoRemoteSetupAccept,
			CmdNoRipFilterRuleRule,
			CmdNoRipPreference,
			CmdNoRipTimer,
			CmdNoRipUseUse,
			CmdNoScheduleAt,
			CmdNoSdUseSwitch,
			CmdNoSecurityClass,
			CmdNoSet,
			CmdNoSftpdHost,
			CmdNoSip100Rel,
			CmdNoSipArriveAddressCheck,
			CmdNoSipArriveRingingPNUatype,
			CmdNoSipArriveSessionTimerMethod,
			CmdNoSipArriveSessionTimerRefresher,
			CmdNoSipIpProtocol,
			CmdNoSipLog,
			CmdNoSipOuterAddress,
			CmdNoSipResponseCodeBusy,
			CmdNoSipServer100Rel,
			CmdNoSipServerCallOwnPermit,
			CmdNoSipServerCallRemoteDomain,
			CmdNoSipServerDialNumberOnly,
			CmdNoSipServerDisplayName,
			CmdNoSipServerNumber,
			CmdNoSipServerPilotAddress,
			CmdNoSipServerPrivacy,
			CmdNoSipServerQvalue,
			CmdNoSipServerRegisterRequestUri,
			CmdNoSipServerRegisterTimer,
			CmdNoSipServerSessionTimer,
			CmdNoSipSessionTimer,
			CmdNoSipUse,
			CmdNoSipUserAgent,
			CmdNoSnmpCommunityReadOnly,
			CmdNoSnmpCommunityReadWrite,
			CmdNoSnmpDisplayIpcpForce,
			CmdNoSnmpHostHost,
			CmdNoSnmpIfindexSwitchStaticIndex,
			CmdNoSnmpLocalAddress,
			CmdNoSnmpSyscontact,
			CmdNoSnmpSyslocation,
			CmdNoSnmpSysname,
			CmdNoSnmpTrapCommunity,
			CmdNoSnmpTrapDelayTimer,
			CmdNoSnmpTrapEnableSnmp,
			CmdNoSnmpTrapEnableSwitch,
			CmdNoSnmpTrapEnableSwitchCommon,
			CmdNoSnmpTrapHost,
			CmdNoSnmpTrapLinkUpdownSeparateL2SwitchPort,
			CmdNoSnmpTrapMobileSignalStrength,
			CmdNoSnmpTrapSendLinkdown,
			CmdNoSnmpTrapSendLinkdownPp,
			CmdNoSnmpTrapSendLinkdownTunnel,
			CmdNoSnmpv2CCommunityReadOnly,
			CmdNoSnmpv2CCommunityReadWrite,
			CmdNoSnmpv2CHostHost,
			CmdNoSnmpv2CTrapCommunity,
			CmdNoSnmpv2CTrapHost,
			CmdNoSnmpv3ContextName,
			CmdNoSnmpv3EngineId,
			CmdNoSnmpv3TrapHost,
			CmdNoSnmpv3UsmUser,
			CmdNoSnmpv3VacmAccess,
			CmdNoSnmpv3VacmView,
			CmdNoSnmpYrifppdisplayatmib2,
			CmdNoSnmpYrifswitchdisplayatmib2,
			CmdNoSnmpYriftunneldisplayatmib2,
			CmdNoSnmpYrswindexSwitchStaticIndex,
			CmdNoSntpdHost,
			CmdNoSntpdService,
			CmdNoSpeed,
			CmdNoSpeedPp,
			CmdNoSshdClientAlive,
			CmdNoSshdEncryptAlgorithm,
			CmdNoSshdHideOpensshVersion,
			CmdNoSshdHost,
			CmdNoSshdHostKeyGenerate,
			CmdNoSshdListenPort,
			CmdNoSshdService,
			CmdNoSshdSession,
			CmdNoSshEncryptAlgorithm,
			CmdNoSshKnownHosts,
			CmdNoStatistics,
			CmdNoSwitchConfigDirectoryPath,
			CmdNoSwitchConfigFilename,
			CmdNoSwitchControlFunctionSetFunctionIndex,
			CmdNoSwitchControlMode,
			CmdNoSwitchControlRouteBackupRoute,
			CmdNoSwitchControlUse,
			CmdNoSwitchControlWatchInterval,
			CmdNoSwitchSelect,
			CmdNoSyslogDebug,
			CmdNoSyslogExecuteCommand,
			CmdNoSyslogFacility,
			CmdNoSyslogHost,
			CmdNoSyslogInfo,
			CmdNoSyslogLocalAddress,
			CmdNoSyslogNotice,
			CmdNoSyslogSrcport,
			CmdNoSystemLedBrightness,
			CmdNoSystemPacketBufferGroupParameterValue,
			CmdNoSystemPacketScheduling,
			CmdNoSystemPacketSchedulingFilter,
			CmdNoSystemPacketSchedulingFilterList,
			CmdNoSystemTemperatureThreshold,
			CmdNoTcpLog,
			CmdNoTcpSessionLimit,
			CmdNoTelnetdHost,
			CmdNoTelnetdListen,
			CmdNoTelnetdService,
			CmdNoTelnetdSession,
			CmdNoTftpHost,
			CmdNoTimezone,
			CmdNoTunnelBackup,
			CmdNoTunnelEnable,
			CmdNoTunnelEncapsulation,
			CmdNoTunnelEndpointAddress,
			CmdNoTunnelEndpointLocalAddress,
			CmdNoTunnelEndpointName,
			CmdNoTunnelEndpointRemoteAddress,
			CmdNoTunnelMultipointLimit,
			CmdNoTunnelMultipointLocalName,
			CmdNoTunnelMultipointServer,
			CmdNoTunnelName,
			CmdNoTunnelNgnArrivePermitPermit,
			CmdNoTunnelNgnBandwidthBandwidth,
			CmdNoTunnelNgnCallPermit,
			CmdNoTunnelNgnDisconnectTime,
			CmdNoTunnelNgnFallback,
			CmdNoTunnelNgnInterface,
			CmdNoTunnelNgnRadiusAuth,
			CmdNoTunnelSelect,
			CmdNoTunnelTemplate,
			CmdNoTunnelType,
			CmdNoUpnpExternalAddressRefer,
			CmdNoUpnpMappingTimerType,
			CmdNoUpnpPortMappingTimer,
			CmdNoUpnpSyslog,
			CmdNoUpnpUse,
			CmdNoUrlFilter,
			CmdNoUrlFilterLog,
			CmdNoUrlFilterPort,
			CmdNoUrlFilterReject,
			CmdNoUrlFilterUse,
			CmdNoUrlInterfaceFilter,
			CmdNoUrlPpFilterDir,
			CmdNoUrlTunnelFilterDir,
			CmdNoUsbhostConfigFilename,
			CmdNoUsbhostExecFilenameFromTo,
			CmdNoUsbhostModemFlowControl,
			CmdNoUsbhostModemInitialize,
			CmdNoUsbhostOvercurrentDuration,
			CmdNoUsbhostStatisticsFilenamePrefix,
			CmdNoUsbhostSyslogFilename,
			CmdNoUsbhostUse,
			CmdNoUserAttribute,
			CmdNoVlanInterface,
			CmdNoVlanPortMapping,
			CmdNoWanAccessLimitConnectionLength,
			CmdNoWanAccessLimitConnectionTime,
			CmdNoWanAccessLimitDurationDuration,
			CmdNoWanAccessLimitLength,
			CmdNoWanAccessLimitTime,
			CmdNoWanAccessPointName,
			CmdNoWanAlwaysOn,
			CmdNoWanAuthMyname,
			CmdNoWanAutoConnect,
			CmdNoWanBind,
			CmdNoWanDisconnectInputTime,
			CmdNoWanDisconnectOutputTime,
			CmdNoWanDisconnectTime,
			CmdNoWinsServer,
			CmdNoYnoAccessCode,
			CmdNoYnoGuiForwarderTimeout,
			CmdNoYnoHttpsProxy,
			CmdNoYnoLog,
			CmdNoYnoUseSw,
			CmdNslookup,
			CmdNtpBackwardCompatibility,
			CmdNtpdate,
			CmdNtpLocalAddress,
			CmdOperationButtonFunctionDownload,
			CmdOperationExecuteBatchPermit,
			CmdOperationExternalMemoryDownloadPermit,
			CmdOperationHttpRevisionUpPermit,
			CmdOperationUsbDownloadPermit,
			CmdOspfArea,
			CmdOspfAreaNetwork,
			CmdOspfAreaStubhost,
			CmdOspfConfigureRefresh,
			CmdOspfExportFilter,
			CmdOspfExportFromOspf,
			CmdOspfImportFilter,
			CmdOspfImportFrom,
			CmdOspfLog,
			CmdOspfMergeEqualCostStub,
			CmdOspfPreference,
			CmdOspfRericInterval,
			CmdOspfRouterId,
			CmdOspfUse,
			CmdOspfVirtualLink,
			CmdPkiCertificateFile,
			CmdPkiCrlFile,
			CmdPpAlwaysOn,
			CmdPpAuthAccept,
			CmdPpAuthMultiConnectProhibit,
			CmdPpAuthMyname,
			CmdPpAuthRequest,
			CmdPpAuthUsername,
			CmdPpBackup,
			CmdPpBackupPp,
			CmdPpBackupRecoveryTime,
			CmdPpBackupTunnel,
			CmdPpBind,
			CmdPpDisable,
			CmdPpEnable,
			CmdPpEncapsulation,
			CmdPpKeepaliveInterval,
			CmdPpKeepaliveLog,
			CmdPpKeepaliveUse,
			CmdPpName,
			CmdPppBacpMaxconfigure,
			CmdPppBacpMaxfailure,
			CmdPppBacpMaxterminate,
			CmdPppBacpRestart,
			CmdPppBapMaxretry,
			CmdPppBapRestart,
			CmdPppCcpMaxconfigure,
			CmdPppCcpMaxfailure,
			CmdPppCcpMaxterminate,
			CmdPppCcpNoEncryption,
			CmdPppCcpRestart,
			CmdPppCcpType,
			CmdPppChapMaxchallenge,
			CmdPppChapRestart,
			CmdPppIpcpIpaddress,
			CmdPppIpcpMaxconfigure,
			CmdPppIpcpMaxfailure,
			CmdPppIpcpMaxterminate,
			CmdPppIpcpMsext,
			CmdPppIpcpRemoteAddressCheck,
			CmdPppIpcpRestart,
			CmdPppIpcpVjc,
			CmdPppIpv6CpUse,
			CmdPppLcpAccm,
			CmdPppLcpAcfc,
			CmdPppLcpMagicnumber,
			CmdPppLcpMaxconfigure,
			CmdPppLcpMaxfailure,
			CmdPppLcpMaxterminate,
			CmdPppLcpMru,
			CmdPppLcpPfc,
			CmdPppLcpRestart,
			CmdPppLcpSilent,
			CmdPppMpControl,
			CmdPppMpDivide,
			CmdPppMpInterleave,
			CmdPppMpLoadThreshold,
			CmdPppMpMaxlink,
			CmdPppMpMinlink,
			CmdPppMpTimer,
			CmdPppMpUse,
			CmdPppMscbcpMaxretry,
			CmdPppMscbcpRestart,
			CmdPppoeAccessConcentrator,
			CmdPppoeAutoConnect,
			CmdPppoeAutoDisconnect,
			CmdPppoeDisconnectTime,
			CmdPppoeInvalidSessionForcedClose,
			CmdPppoePadiMaxretry,
			CmdPppoePadiRestart,
			CmdPppoePadrMaxretry,
			CmdPppoePadrRestart,
			CmdPppoePassThroughMember,
			CmdPppoeServiceName,
			CmdPppoeTcpMssLimit,
			CmdPppoeUse,
			CmdPppPapMaxauthreq,
			CmdPppPapRestart,
			CmdPPSelect,
			CmdPptpCallIdMode,
			CmdPptpHostname,
			CmdPptpKeepaliveInterval,
			CmdPptpKeepaliveLog,
			CmdPptpKeepaliveUse,
			CmdPptpService,
			CmdPptpServiceType,
			CmdPptpSyslog,
			CmdPptpTunnelDisconnectTime,
			CmdPptpVendorname,
			CmdPptpWindowSize,
			CmdPriLeasedChannel,
			CmdPriLoopbackActive,
			CmdPriLoopbackPassive,
			CmdPriLoopbackPassivePriOff,
			CmdProviderAutoConnectForcedDisable,
			CmdProviderDnsServer,
			CmdProviderDnsServerPp,
			CmdProviderFilterRouting,
			CmdProviderInterfaceBind,
			CmdProviderInterfaceDnsServer,
			CmdProviderInterfaceName,
			CmdProviderIpv6ConnectPp,
			CmdProviderNtpdate,
			CmdProviderNtpServer,
			CmdProviderSelect,
			CmdProviderSet,
			CmdProviderType,
			CmdQacTmClientPermit,
			CmdQacTmClientPort,
			CmdQacTmClientRefreshGo,
			CmdQacTmClientUpdate,
			CmdQacTmPort,
			CmdQacTmRedirect,
			CmdQacTmServer,
			CmdQacTmServerRefreshGo,
			CmdQacTmUnqualifiedClientAccessControl,
			CmdQacTmUse,
			CmdQacTmVersionMargin,
			CmdQacTmWarningUrl,
			CmdQueueClassFilter,
			CmdQueueClassFilterDscp,
			CmdQueueClassFilterPrecedence,
			CmdQueueInterfaceClassControl,
			CmdQueueInterfaceClassFilterList,
			CmdQueueInterfaceClassProperty,
			CmdQueueInterfaceDefaultClass,
			CmdQueueInterfaceDefaultClassSecondary,
			CmdQueueInterfaceLength,
			CmdQueueInterfaceLengthSecondary,
			CmdQueueInterfaceType,
			CmdQueuePpClassFilterList,
			CmdQueuePpClassProperty,
			CmdQueuePpDefaultClass,
			CmdQueuePpLength,
			CmdQueuePpType,
			CmdQueuePpTypeType,
			CmdQueueTunnelClassFilterList,
			CmdQuit,
			CmdRadiusAccount,
			CmdRadiusAccountPort,
			CmdRadiusAccountServer,
			CmdRadiusAuth,
			CmdRadiusAuthPort,
			CmdRadiusAuthServer,
			CmdRadiusRetry,
			CmdRadiusSecret,
			CmdRadiusServer,
			CmdRdate,
			CmdRemoteSetup,
			CmdRemoteSetupAccept,
			CmdRename,
			CmdRestart,
			CmdRestartConfig,
			CmdRipFilterRule,
			CmdRipPreference,
			CmdRipTimer,
			CmdRipUse,
			CmdRollbackTimer,
			CmdRotateExternalMemorySyslog,
			CmdRtfsFormat,
			CmdRtfsGarbageCollect,
			CmdSave,
			CmdScheduleAt,
			CmdScp,
			CmdSdUse,
			CmdSecurityClass,
			CmdSet,
			CmdSetDefaultConfig,
			CmdSetDefaultExec,
			CmdSetSerialBaudrate,
			CmdSftpdHost,
			CmdShowAccount,
			CmdShowAccountMobile,
			CmdShowAccountNgnData,
			CmdShowAccountPp,
			CmdShowAccountTunnel,
			CmdShowAlias,
			CmdShowArp,
			CmdShowBridgeLearning,
			CmdShowCommand,
			CmdShowCommandHistory,
			CmdShowConfig,
			CmdShowConfigFilename,
			CmdShowConfigPp,
			CmdShowConfigTunnel,
			CmdShowCopyright,
			CmdShowDiagnosisConfigPortAccess,
			CmdShowDiagnosisConfigPortMap,
			CmdShowDlci,
			CmdShowDnsCache,
			CmdShowEnvironment,
			CmdShowExecList,
			CmdShowFileList,
			CmdShowGrep,
			CmdShowHistory,
			CmdShowIpConnection,
			CmdShowIpConnectionPp,
			CmdShowIpConnectionTunnel,
			CmdShowIpIntrusionDetection,
			CmdShowIpIntrusionDetectionPp,
			CmdShowIpIntrusionDetectionTunnel,
			CmdShowIpMroute,
			CmdShowIpRipTable,
			CmdShowIpRoute,
			CmdShowIpsecSa,
			CmdShowIpsecSaGateway,
			CmdShowIpSecureFilter,
			CmdShowIpSecureFilterPp,
			CmdShowIpSecureFilterTunnel,
			CmdShowIpTrafficList,
			CmdShowIpTrafficListPp,
			CmdShowIpTrafficListTunnel,
			CmdShowIpv6Address,
			CmdShowIpv6AddressPp,
			CmdShowIpv6AddressTunnel,
			CmdShowIpv6Connection,
			CmdShowIpv6ConnectionPp,
			CmdShowIpv6ConnectionTunnel,
			CmdShowIpv6MrouteFib,
			CmdShowIpv6NeighborCache,
			CmdShowIpv6Ospf,
			CmdShowIpv6RipTable,
			CmdShowIpv6Route,
			CmdShowLanMap,
			CmdShowLess,
			CmdShowLineMasterclock,
			CmdShowLog,
			CmdShowMacro,
			CmdShowNatDescriptorAddress,
			CmdShowNatDescriptorInterfaceAddress,
			CmdShowNatDescriptorInterfaceAddressPp,
			CmdShowNatDescriptorInterfaceAddressTunnel,
			CmdShowNatDescriptorInterfaceBind,
			CmdShowNatDescriptorInterfaceBindPp,
			CmdShowNatDescriptorInterfaceBindTunnel,
			CmdShowNatDescriptorMasqueradePortSummary,
			CmdShowNatDescriptorMasqueradeSessionSummary,
			CmdShowPkiCertificateSummary,
			CmdShowPkiCrl,
			CmdShowPpConnectTime,
			CmdShowPppoePassThroughLearning,
			CmdShowSet,
			CmdShowSshdPublicKey,
			CmdShowStatus,
			CmdShowStatusBackup,
			CmdShowStatusBgpNeighbor,
			CmdShowStatusBoot,
			CmdShowStatusBootAll,
			CmdShowStatusBootList,
			CmdShowStatusCloudVpn,
			CmdShowStatusCooperation,
			CmdShowStatusDhcp,
			CmdShowStatusDhcpc,
			CmdShowStatusEthernetFilter,
			CmdShowStatusExternalMemory,
			CmdShowStatusHeartbeat,
			CmdShowStatusHeartbeat2,
			CmdShowStatusHeartbeat2Id,
			CmdShowStatusHeartbeat2Name,
			CmdShowStatusIpIgmp,
			CmdShowStatusIpInboundFilter,
			CmdShowStatusIpip,
			CmdShowStatusIpKeepalive,
			CmdShowStatusIpPimSparse,
			CmdShowStatusIpPolicyFilter,
			CmdShowStatusIpPolicyService,
			CmdShowStatusIpv6Dhcp,
			CmdShowStatusIpv6InboundFilter,
			CmdShowStatusIpv6Mld,
			CmdShowStatusIpv6PolicyFilter,
			CmdShowStatusIpv6PolicyFilterIdType,
			CmdShowStatusIpv6PolicyService,
			CmdShowStatusL2Tp,
			CmdShowStatusLua,
			CmdShowStatusMailService,
			CmdShowStatusMobileSignalStrength,
			CmdShowStatusNetvolanteDns,
			CmdShowStatusNetvolanteDnsPp,
			CmdShowStatusNgn,
			CmdShowStatusOspf,
			CmdShowStatusPacketBuffer,
			CmdShowStatusPacketScheduling,
			CmdShowStatusPp,
			CmdShowStatusPptp,
			CmdShowStatusQacTm,
			CmdShowStatusQacTmClient,
			CmdShowStatusQacTmQualified,
			CmdShowStatusQacTmServer,
			CmdShowStatusQacTmUnqualified,
			CmdShowStatusQos,
			CmdShowStatusRemoteSetup,
			CmdShowStatusRtfs,
			CmdShowStatusSd,
			CmdShowStatusStatusLed,
			CmdShowStatusSwitchControl,
			CmdShowStatusSwitchControlRouteBackup,
			CmdShowStatusSwitchingHubMacaddress,
			CmdShowStatusTunnel,
			CmdShowStatusUpnp,
			CmdShowStatusUsbhost,
			CmdShowStatusUser,
			CmdShowStatusUserHistory,
			CmdShowStatusVlan,
			CmdShowStatusVrrp,
			CmdShowStatusYno,
			CmdShowTechinfo,
			CmdShowUrlFilter,
			CmdShowUrlFilterPp,
			CmdShowUrlFilterTunnel,
			CmdSip100Rel,
			CmdSipArriveAddressCheck,
			CmdSipArriveRingingPNUatype,
			CmdSipArriveSessionTimerMethod,
			CmdSipArriveSessionTimerRefresher,
			CmdSipIpProtocol,
			CmdSipLog,
			CmdSipOuterAddress,
			CmdSipResponseCodeBusy,
			CmdSipServer,
			CmdSipServer100Rel,
			CmdSipServerCallOwnPermit,
			CmdSipServerCallRemoteDomain,
			CmdSipServerConnect,
			CmdSipServerDialNumberOnly,
			CmdSipServerDisconnect,
			CmdSipServerDisplayName,
			CmdSipServerPilotAddress,
			CmdSipServerPrivacy,
			CmdSipServerQvalue,
			CmdSipServerRegisterRequestUri,
			CmdSipServerRegisterTimer,
			CmdSipServerSessionTimer,
			CmdSipSessionTimer,
			CmdSipUse,
			CmdSipUserAgent,
			CmdSnmpCommunityReadOnly,
			CmdSnmpCommunityReadWrite,
			CmdSnmpDisplayIpcpForce,
			CmdSnmpHost,
			CmdSnmpIfindexSwitchStaticIndex,
			CmdSnmpLocalAddress,
			CmdSnmpSyscontact,
			CmdSnmpSyslocation,
			CmdSnmpSysname,
			CmdSnmpTrapCommunity,
			CmdSnmpTrapDelayTimer,
			CmdSnmpTrapDelayTimerOff,
			CmdSnmpTrapEnableSnmp,
			CmdSnmpTrapEnableSnmpTrap,
			CmdSnmpTrapEnableSwitch,
			CmdSnmpTrapEnableSwitchCommon,
			CmdSnmpTrapHost,
			CmdSnmpTrapLinkUpdownSeparateL2SwitchPort,
			CmdSnmpTrapMobileSignalStrength,
			CmdSnmpTrapSendLinkdown,
			CmdSnmpTrapSendLinkdownPp,
			CmdSnmpTrapSendLinkdownTunnel,
			CmdSnmpv2CCommunityReadOnly,
			CmdSnmpv2CCommunityReadWrite,
			CmdSnmpv2CHost,
			CmdSnmpv2CTrapCommunity,
			CmdSnmpv2CTrapHost,
			CmdSnmpv3ContextName,
			CmdSnmpv3EngineId,
			CmdSnmpv3Host,
			CmdSnmpv3TrapHost,
			CmdSnmpv3UsmUser,
			CmdSnmpv3VacmAccess,
			CmdSnmpv3VacmView,
			CmdSnmpYrifppdisplayatmib2,
			CmdSnmpYrifswitchdisplayatmib2,
			CmdSnmpYriftunneldisplayatmib2,
			CmdSnmpYrswindexSwitchStaticIndex,
			CmdSntpdHost,
			CmdSntpdService,
			CmdSpeed,
			CmdSpeedPp,
			CmdSsh,
			CmdSshdClientAlive,
			CmdSshdEncryptAlgorithm,
			CmdSshdHideOpensshVersion,
			CmdSshdHost,
			CmdSshdHostAny,
			CmdSshdHostKeyGenerate,
			CmdSshdHostLan,
			CmdSshdHostNone,
			CmdSshdListen,
			CmdSshdService,
			CmdSshdSession,
			CmdSshEncryptAlgorithm,
			CmdSshKnownHosts,
			CmdStatistics,
			CmdSwitchConfigDirectory,
			CmdSwitchConfigFilename,
			CmdSwitchControlConfigGet,
			CmdSwitchControlConfigGetInterface,
			CmdSwitchControlConfigSet,
			CmdSwitchControlFirmwareUploadGo,
			CmdSwitchControlFunctionDefault,
			CmdSwitchControlFunctionExecute,
			CmdSwitchControlFunctionGet,
			CmdSwitchControlFunctionSet,
			CmdSwitchControlMode,
			CmdSwitchControlRouteBackup,
			CmdSwitchControlUse,
			CmdSwitchControlWatchInterval,
			CmdSwitchSelect,
			CmdSyslogDebug,
			CmdSyslogExecuteCommand,
			CmdSyslogFacility,
			CmdSyslogHost,
			CmdSyslogInfo,
			CmdSyslogLocalAddress,
			CmdSyslogNotice,
			CmdSyslogSrcport,
			CmdSystemLedBrightness,
			CmdSystemPacketBuffer,
			CmdSystemPacketScheduling,
			CmdSystemPacketSchedulingFilter,
			CmdSystemPacketSchedulingFilterList,
			CmdSystemTemperatureThreshold,
			CmdTakeLanMapSnapshot,
			CmdTcpLog,
			CmdTcpSessionLimit,
			CmdTelnet,
			CmdTelnetdHost,
			CmdTelnetdListen,
			CmdTelnetdService,
			CmdTelnetdSession,
			CmdTerminateLua,
			CmdTerminateLuaFile,
			CmdTftpHost,
			CmdTime,
			CmdTimezone,
			CmdTraceroute,
			CmdTraceroute6,
			CmdTunnelBackup,
			CmdTunnelDisable,
			CmdTunnelEnable,
			CmdTunnelEncapsulation,
			CmdTunnelEndpointAddress,
			CmdTunnelEndpointLocalAddress,
			CmdTunnelEndpointName,
			CmdTunnelEndpointRemoteAddress,
			CmdTunnelMultipointLimit,
			CmdTunnelMultipointLocalName,
			CmdTunnelMultipointServer,
			CmdTunnelName,
			CmdTunnelNgnArrivePermit,
			CmdTunnelNgnBandwidth,
			CmdTunnelNgnCallPermit,
			CmdTunnelNgnDisconnectTime,
			CmdTunnelNgnFallback,
			CmdTunnelNgnInterface,
			CmdTunnelNgnRadiusAuth,
			CmdTunnelSelect,
			CmdTunnelTemplate,
			CmdTunnelType,
			CmdUpnpExternalAddressRefer,
			CmdUpnpPortMappingTimer,
			CmdUpnpPortMappingTimerType,
			CmdUpnpSyslog,
			CmdUpnpUse,
			CmdUrlFilter,
			CmdUrlFilterLog,
			CmdUrlFilterPort,
			CmdUrlFilterReject,
			CmdUrlFilterUse,
			CmdUrlInterfaceFilter,
			CmdUrlPpFilter,
			CmdUrlTunnelFilter,
			CmdUsbhostConfigFilename,
			CmdUsbhostExecFilename,
			CmdUsbhostModemFlowControl,
			CmdUsbhostModemInitialize,
			CmdUsbhostOvercurrentDuration,
			CmdUsbhostStatisticsFilenamePrefix,
			CmdUsbhostSyslogFilename,
			CmdUsbhostUse,
			CmdUserAttribute,
			CmdVlanInterface8021Q,
			CmdVlanPortMapping,
			CmdWanAccessLimitConnectionLength,
			CmdWanAccessLimitConnectionTime,
			CmdWanAccessLimitDuration,
			CmdWanAccessLimitLength,
			CmdWanAccessLimitTime,
			CmdWanAccessPointName,
			CmdWanAlwaysOn,
			CmdWanAuthMyname,
			CmdWanAutoConnect,
			CmdWanBind,
			CmdWanDisconnectInputTime,
			CmdWanDisconnectOutputTime,
			CmdWanDisconnectTime,
			CmdWinsServer,
			CmdWolSend,
			CmdYnoAccessCode,
			CmdYnoGuiForwarderTimeout,
			CmdYnoHttpsProxy,
			CmdYnoLog,
			CmdYnoUse,
		},
	}
	for _, spec := range sim.Commands {
		switch t := spec.Tokens.(type) {
		case []Token:
			sim.CommandTrie.Add(t, spec.Arity, spec)
		case []TokensArityPair:
			for _, t := range t {
				sim.CommandTrie.Add(t.Tokens, t.Arity, spec)
			}
		default:
			panic("never get here")
		}
	}
	return sim
}

func (sim *Simulator) NewSession(tf base.TerminalFactory) base.ApplianceSession {
	sess := &SimulatorSession{
		Simulator: sim,
	}
	sess.Terminal = tf(sess.autoCompleteCallback)
	sess.Charset = sim.SavedConfig.Charset
	sess.Reset()
	return sess
}

func (sim *Simulator) ResolveInterface(name string) NetInterface {
	name = strings.ToLower(name)
	for _, if_ := range sim.PhyInterfaces {
		if if_.Name() == name {
			return if_
		}
	}
	for _, if_ := range sim.PortVlanInterfaces {
		if if_.Name() == name {
			return if_
		}
	}
	for _, if_ := range sim.PortVlanGroupInterfaces {
		if if_.Name() == name {
			return if_
		}
	}
	for _, ifs := range sim.TaggedVlanInterfaces {
		for _, if_ := range ifs {
			if if_.Name() == name {
				return if_
			}
		}
	}
	return nil
}
