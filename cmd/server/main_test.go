package main

import "testing"

type fakeServer struct {
	addr string
	err  error
}

func (s *fakeServer) Run(addr string) error {
	s.addr = addr
	return s.err
}

func TestRunUsesDefaultPortWhenUnset(t *testing.T) {
	originalNewServer := newServer
	t.Cleanup(func() {
		newServer = originalNewServer
	})
	t.Setenv("PORT", "")

	server := &fakeServer{}
	newServer = func() serverRunner {
		return server
	}

	if err := run(); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	if server.addr != ":8081" {
		t.Fatalf("expected default port :8081, got %q", server.addr)
	}
}

func TestRunUsesConfiguredPort(t *testing.T) {
	originalNewServer := newServer
	t.Cleanup(func() {
		newServer = originalNewServer
	})
	t.Setenv("PORT", "9090")

	server := &fakeServer{}
	newServer = func() serverRunner {
		return server
	}

	if err := run(); err != nil {
		t.Fatalf("run returned error: %v", err)
	}
	if server.addr != ":9090" {
		t.Fatalf("expected configured port :9090, got %q", server.addr)
	}
}

func TestMainCallsExecute(t *testing.T) {
	originalExecute := execute
	t.Cleanup(func() {
		execute = originalExecute
	})

	called := false
	execute = func() error {
		called = true
		return nil
	}

	main()

	if !called {
		t.Fatal("expected main to call execute")
	}
}