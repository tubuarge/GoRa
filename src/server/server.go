package server

import (
	"github.com/dgraph-io/badger"
	"github.com/hashicorp/raft"
	raftboltdb "github.com/hashicorp/raft-boltdb"
)

func NewRaft(path string) (*raft.Raft, error) {
	fsm, err := NewFSM(path, paht)
	if err != nil {
		return nil, err
	}

	boltDB, err := raftboltdb.NewBoltStore(filepath.Join("./", "raft.db"))
	if err != nil {
		return nil, err
	}

	ss, err := raft.NewFileSnapshotStore("./", os.Stderr)
	if err != nil {
		return nil, err
	}

	addr, err := net.ResolveTCPAddr("tcp", ":15379")
	if err != nil {
		return nil, err
	}

	transport, err := raft.NewTCPTransport(":15379", addr, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return nil, err
	}

	deafultRaftCfg := raft.DefaultConfig()

	r, err := raft.NewRaft(defaultRaftCfg, fsm, boltDB, boltDB, ss, transport)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func NewFSM(dir, valueDir string) (*badger.DB, error) {
	opts := badger.DefaultOptions
	opts.Dir = dir
	opts.ValueDir = valueDir
	opts.SyncWrites = false
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}
	return db, nil
}
