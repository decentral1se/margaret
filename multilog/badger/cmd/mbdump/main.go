// SPDX-License-Identifier: MIT

package main

import (
	"fmt"
	"os"

	"github.com/dgraph-io/badger"
	"github.com/pkg/errors"
	"go.mindeco.de/logging"

	"go.cryptoscope.co/margaret"
	"go.cryptoscope.co/margaret/codec/msgpack"
	librarian "go.cryptoscope.co/margaret/indexes"
	"go.cryptoscope.co/margaret/multilog"
	mlbadger "go.cryptoscope.co/margaret/multilog/badger"
)

var check = logging.CheckFatal

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s <dir> (hasAddr)", os.Args[0])
		os.Exit(1)
	}
	logging.SetupLogging(nil)
	log := logging.Logger(os.Args[0])

	dir := os.Args[1]

	opts := badger.DefaultOptions(dir)

	db, err := badger.Open(opts)
	check(errors.Wrap(err, "error opening database"))

	mlog := mlbadger.New(db, msgpack.New(margaret.BaseSeq(0)))

	addrs, err := mlog.List()
	check(errors.Wrap(err, "error listing multilog"))
	log.Log("mlog", "opened", "list#", len(addrs))

	for i, addr := range addrs {
		log.Log("i", i, "addr", addr)
	}

	// check has
	if len(os.Args) > 2 {
		addr := librarian.Addr(os.Args[2])
		has, err := multilog.Has(mlog, addr)
		log.Log("mlog", "has", "addr", addr, "has?", has, "hasErr", err)
	}
}
