package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/signal"
	"time"

	server "github.com/moriyoshi/yamaha-rt-simulator/server"
	yrtsim "github.com/moriyoshi/yamaha-rt-simulator/simulator"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh"
)

var (
	bind        string
	hostKeyFile string
	debug       bool
)

func init() {
	flag.StringVar(&bind, "bind", ":10022", "address:port to listen on")
	flag.StringVar(&hostKeyFile, "hostkey", "./host_key", "host key file")
	flag.BoolVar(&debug, "debug", false, "turn on debug output")
}

func bail(msg interface{}) {
	fmt.Fprintln(os.Stderr, os.Args[0]+":", " ", msg)
	os.Exit(1)
}

var yamahaRTSpec = func() *yrtsim.Spec {
	lan1 := &yrtsim.PhyInterface{
		Name_:   "lan1",
		HwAddr_: yrtsim.MustParseMacAddr("ac:44:f2:39:4f:cb"),
	}
	phyInterfaces := []*yrtsim.PhyInterface{
		lan1,
		&yrtsim.PhyInterface{
			Name_:   "lan2",
			HwAddr_: yrtsim.MustParseMacAddr("ac:44:f2:39:4f:cc"),
		},
		&yrtsim.PhyInterface{
			Name_:   "lan3",
			HwAddr_: yrtsim.MustParseMacAddr("ac:44:f2:39:4f:cd"),
		},
	}
	portVlanInterfaces := []*yrtsim.PortVlanInterface{
		&yrtsim.PortVlanInterface{
			Index: 1,
			Phy:   lan1,
		},
		&yrtsim.PortVlanInterface{
			Index: 2,
			Phy:   lan1,
		},
		&yrtsim.PortVlanInterface{
			Index: 3,
			Phy:   lan1,
		},
		&yrtsim.PortVlanInterface{
			Index: 4,
			Phy:   lan1,
		},
		&yrtsim.PortVlanInterface{
			Index: 5,
			Phy:   lan1,
		},
		&yrtsim.PortVlanInterface{
			Index: 6,
			Phy:   lan1,
		},
		&yrtsim.PortVlanInterface{
			Index: 7,
			Phy:   lan1,
		},
		&yrtsim.PortVlanInterface{
			Index: 8,
			Phy:   lan1,
		},
	}
	portVlanGroupInterfaces := []*yrtsim.PortVlanGroupInterface{
		&yrtsim.PortVlanGroupInterface{
			Name_: "vlan1",
		},
		&yrtsim.PortVlanGroupInterface{
			Name_: "vlan2",
		},
		&yrtsim.PortVlanGroupInterface{
			Name_: "vlan3",
		},
		&yrtsim.PortVlanGroupInterface{
			Name_: "vlan4",
		},
		&yrtsim.PortVlanGroupInterface{
			Name_: "vlan5",
		},
		&yrtsim.PortVlanGroupInterface{
			Name_: "vlan6",
		},
		&yrtsim.PortVlanGroupInterface{
			Name_: "vlan7",
		},
		&yrtsim.PortVlanGroupInterface{
			Name_: "vlan8",
		},
	}

	taggedVlanInterfaces := map[*yrtsim.PhyInterface][]*yrtsim.TaggedVlanInterface{}
	for _, phy := range phyInterfaces {
		ifs := make([]*yrtsim.TaggedVlanInterface, 0, 32)
		for i := 1; i <= 32; i++ {
			ifs = append(ifs, &yrtsim.TaggedVlanInterface{
				Index: i,
				Phy:   lan1,
			})
		}
		taggedVlanInterfaces[phy] = ifs
	}
	return &yrtsim.Spec{
		yrtsim.HardwareSpec{
			ModelName:               "RTX1210",
			PhyInterfaces:           phyInterfaces,
			PortVlanInterfaces:      portVlanInterfaces,
			PortVlanGroupInterfaces: portVlanGroupInterfaces,
			TaggedVlanInterfaces:    taggedVlanInterfaces,
		},
		yrtsim.SystemSpec{
			Firmware: yrtsim.Firmware{
				Version:   "14.01.20",
				Date:      time.Date(2017, time.June, 1, 7, 52, 40, 0, time.UTC),
				Copyright: "Copyright (c) 2018 Moriyoshi Koizumi. All Rights Reserved.",
			},
			Charsets: []yrtsim.NLSInfo{
				yrtsim.NLSInfo{
					Name:     "us.ascii",
					Variants: []string{"ascii", "us.ascii"},
					Encoding: "us-ascii",
					Locale:   "en",
				},
				yrtsim.NLSInfo{
					Name:     "ja.sjis",
					Variants: []string{"sjis", "ja.sjis"},
					Encoding: "Shift_JIS",
					Locale:   "ja",
				},
				yrtsim.NLSInfo{
					Name:     "ja.euc",
					Variants: []string{"euc", "ja.euc"},
					Encoding: "EUC-JP",
					Locale:   "ja",
				},
				yrtsim.NLSInfo{
					Name:     "ja.utf8",
					Variants: []string{"utf8", "ja.utf8"},
					Encoding: "UTF-8",
					Locale:   "ja",
				},
			},
		},
	}
}()

var yamahaRTConfig = &yrtsim.SimulatorConfig{
	Prompt:  "",
	Charset: yamahaRTSpec.Charsets[0],
}

func main() {
	flag.Parse()

	sCfg := &ssh.ServerConfig{
		PasswordCallback: func(conn ssh.ConnMetadata, password []byte) (*ssh.Permissions, error) {
			if len(password) == 0 {
				return nil, nil
			} else {
				return nil, fmt.Errorf("password does not match")
			}
		},
	}

	pem, err := ioutil.ReadFile(hostKeyFile)
	if err != nil {
		bail(err)
	}

	hostKey, err := ssh.ParsePrivateKey(pem)
	if err != nil {
		bail(err)
	}

	sCfg.AddHostKey(hostKey)

	conn, err := net.Listen("tcp", bind)
	if err != nil {
		bail(err)
	}
	defer conn.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)

	errChan := make(chan error)
	log := logrus.New()
	if debug {
		log.SetLevel(logrus.DebugLevel)
		log.Debug("debug logging enabled")
	}

	sim := yrtsim.NewSimulator(yamahaRTSpec, yamahaRTConfig, log)

	go func() {
		srv := &server.Server{
			ServerConfig: sCfg,
			Appliance:    sim,
			Log:          log,
		}
		errChan <- srv.RunServerEventLoop(ctx, conn.(*net.TCPListener))
	}()
outer:
	for {
		select {
		case err = <-errChan:
			if err != nil {
				bail(err)
			}
			break outer
		case <-sigChan:
			cancel()
		}
	}
}
