package main

import (
	"errors"
	"testing"
)

type fakeServer struct {
	addr string
	err  error
}

func (s *fakeServer) Run(args ...string) error {
	if len(args) > 0 {
		s.addr = args[0]
	}
	return s.err
}

func TestDefaultNewServerReturnsRunner(t *testing.T) {
	server := defaultNewServer()
	if server == nil {
		t.Fatal("expected defaultNewServer to return a server runner")
	}
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
	originalFatalf := fatalf
	t.Cleanup(func() {
		fatalf = originalFatalf
	})

	called := false
	execute = func() error {
		called = true
		return nil
	}
	fatalf = func(string, ...any) {
		t.Fatal("fatalf should not be called when execute succeeds")
	}

	main()

	if !called {
		t.Fatal("expected main to call execute")
	}
}

func TestMainCallsFatalfOnExecuteError(t *testing.T) {
	originalExecute := execute
	t.Cleanup(func() {
		execute = originalExecute
	})
	originalFatalf := fatalf
	t.Cleanup(func() {
		fatalf = originalFatalf
	})

	called := false
	execute = func() error {
		return errors.New("boom")
	}
	fatalf = func(format string, args ...any) {
		called = true
		if format != "user-service stopped: %v" {
			t.Fatalf("unexpected fatal format: %q", format)
		}
		if len(args) != 1 || args[0].(error).Error() != "boom" {
			t.Fatalf("unexpected fatal args: %#v", args)
		}
	}

	main()

	if !called {
		t.Fatal("expected main to call fatalf on execute error")
	}
}
