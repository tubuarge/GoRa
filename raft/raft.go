package raft

import (
	"fmt"

	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"

	"../store"
)

func NewRaft(path string) (*raft.Raft, error) {
	fsm, err := NewFSM(path, path)
	if err != nil {
		return nil, fmt.Errorf("Couldn't create new FSM: %v", err)
	}

	boltDB, err := raftboltdb.NewBoltStore(filepath.Join("./", "raft.db"))
	if err != nil {
		return nil, fmt.Errorf("Couldn't create new Bolt Store: %v", err)
	}

	ss, err := raft.NewFileSnapshotStore("./", 1, os.Stderr)
	if err != nil {
		return nil, fmt.Errorf("Couldn't create new File Snapshot Store: %v", err)
	}

	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:5379")
	if err != nil {
		return nil, fmt.Errorf("Couldn't create new TCP endpoint: %v", err)
	}

	transport, err := raft.NewTCPTransport(":15379", addr, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return nil, fmt.Errorf("Couldn't create new TCP transport: %v", err)
	}

	defaultRaftCfg := raft.DefaultConfig()
	defaultRaftCfg.LocalID = raft.ServerID("local1")

	r, err := raft.NewRaft(defaultRaftCfg, fsm, boltDB, boltDB, ss, transport)
	if err != nil {
		return nil, fmt.Errorf("Couldn't Create new Raft: %v", err)
	}

	return r, nil
}

func NewFSM(dir, valueDir string) (raft.FSM, error) {
	badgerOpt := badger.DefaultOptions("./")
	badgerDB, err := badger.Open(badgerOpt)
	if err != nil {
		return nil, err
	}

	fsmStore := store.NewBadger(badgerDB)
	return fsmStore, nil
}
