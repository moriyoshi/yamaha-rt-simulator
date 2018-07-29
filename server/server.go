package server

import (
	"context"
	"net"
	"sync"
	"time"

	base "github.com/moriyoshi/yamaha-rt-simulator"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

type Server struct {
	ServerConfig *ssh.ServerConfig
	Appliance    base.Appliance
	Log          interface {
		base.DebugLogger
		base.InfoLogger
		base.ErrorLogger
	}
}

func (s *Server) handleNewSSHChan(ctx context.Context, sConn *ssh.ServerConn, newChan ssh.NewChannel) error {
	if newChan.ChannelType() != "session" {
		s.Log.Info("unsupported channel type ", newChan.ChannelType(), " requested")
		newChan.Reject(ssh.UnknownChannelType, "unknown channel type")
		return nil
	}
	sshChan, requests, err := newChan.Accept()
	if err != nil {
		return err
	}
	defer sshChan.Close()

	go func() {
		defer s.Log.Debug("handleNewSSHChan.handleReqChan ended")
	outer:
		for {
			select {
			case <-ctx.Done():
				break outer
			case req := <-requests:
				if req == nil {
					break outer
				}
				switch req.Type {
				case "shell", "pty-req":
					req.Reply(true, nil)
				default:
					if req.WantReply {
						req.Reply(false, nil)
					}
				}
			}
		}
	}()

	s.Appliance.NewSession(func(accb base.AutoCompleteCallback) base.Terminal {
		t := terminal.NewTerminal(sshChan, "")
		t.AutoCompleteCallback = accb
		return t
	}).Run(ctx)
	return nil
}

func (s *Server) handleClient(ctx context.Context, conn *net.TCPConn) error {
	defer func() {
		s.Log.Info("connection to ", conn.RemoteAddr(), " closed")
		conn.Close()
	}()

	s.Log.Info("client ", conn.RemoteAddr(), " connected")

	wg := sync.WaitGroup{}
	innerCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		<-innerCtx.Done()
		conn.SetDeadline(time.Unix(1, 0))
	}()

	sConn, newChanChan, reqChan, err := ssh.NewServerConn(conn, s.ServerConfig)
	if err != nil {
		return err
	}

	wg.Add(1)
	go func() {
		defer s.Log.Debug("handleClient.handleReqChan ended")
		defer wg.Done()
		for _ = range reqChan {
		}
	}()

	wg.Add(1)
	go func() {
		defer s.Log.Debug("handleClient.handleNewChanChan ended")
		defer wg.Done()
		defer cancel()
	outer:
		for newChan := range newChanChan {
			if newChan == nil {
				break outer
			}
			err := s.handleNewSSHChan(ctx, sConn, newChan)
			if err != nil {
				s.Log.Error(err)
				break outer
			}
		}
	}()

	wg.Wait()
	return nil
}

func (s *Server) RunServerEventLoop(ctx context.Context, l *net.TCPListener) (err error) {
	defer s.Log.Debug("RunServerEventLoop ended")
	wg := sync.WaitGroup{}
	newConnChan := make(chan *net.TCPConn)

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(newConnChan)
		var cConn *net.TCPConn
	outer:
		for {
			cConn, err = l.AcceptTCP()
			if err != nil {
				break
			}
			select {
			case <-ctx.Done():
				cConn.Close()
				break outer
			case newConnChan <- cConn:
			}
		}
	}()

outer:
	for {
		select {
		case <-ctx.Done():
			s.Log.Info("interrut signal received")
			l.SetDeadline(time.Unix(1, 0))
			break outer
		case conn := <-newConnChan:
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := s.handleClient(ctx, conn)
				if err != nil {
					s.Log.Error(err)
				}
			}()
		}
	}

	// drain
	for _ = range newConnChan {
	}

	wg.Wait()

	if IsTimeout(err) {
		err = nil
	}

	return
}
