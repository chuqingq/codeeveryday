package main

import (
	"context"
	"encoding/json"
	"io"
	"os/exec"
	"strings"

	"github.com/chuqingq/go-util"
)

// SubProcess 对os.exec.Cmd的封装，用于启动子进程
type SubProcess struct {
	Cmd     *exec.Cmd
	Stdin   io.WriteCloser
	encoder *json.Encoder
	Stdout  io.ReadCloser
	decoder *json.Decoder
	Stderr  io.ReadCloser
	Alive   bool
	Ctx     context.Context
	Cancel  context.CancelFunc
}

// NewSubProcess 创建一个SubProcess
func NewSubProcess(name string, args ...string) (*SubProcess, error) {
	ctx, cancel := context.WithCancel(context.Background())

	cmd := exec.CommandContext(ctx, name, args...)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		cancel()
		return nil, err
	}
	encoder := json.NewEncoder(stdin)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		cancel()
		return nil, err
	}
	decoder := json.NewDecoder(stdout)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		cancel()
		return nil, err
	}

	return &SubProcess{
		Cmd:     cmd,
		Stdin:   stdin,
		encoder: encoder,
		Stdout:  stdout,
		decoder: decoder,
		Stderr:  stderr,
		Alive:   false,
		Ctx:     ctx,
		Cancel:  cancel,
	}, nil
}

// Start 启动子进程
func (s *SubProcess) Start() error {
	err := s.Cmd.Start()
	if err != nil {
		return err
	}
	s.Alive = true
	return nil
}

// Stop 停止子进程
func (s *SubProcess) Stop() {
	if s.Alive {
		s.Cancel()
		s.Alive = false
		s.Stdin.Close()
		s.Stdout.Close()
		s.Stderr.Close()
		s.Cmd.Wait()
	}
}

// IsAlive 判断子进程是否存活
func (s *SubProcess) IsAlive() bool {
	return s.Alive
}

// Send 向子进程发送消息
func (s *SubProcess) Send(m *util.Message) error {
	err := s.encoder.Encode(m)
	if err == io.ErrClosedPipe || err == io.EOF || strings.Contains(err.Error(), "broken pipe") {
		s.Alive = false
	}
	return err
}

// Recv 从子进程接收消息
func (s *SubProcess) Recv() (*util.Message, error) {
	m := util.NewMessage()
	err := s.decoder.Decode(m)
	if err != nil {
		if err == io.EOF {
			s.Alive = false
		}
		return nil, err
	}
	return m, nil
}
