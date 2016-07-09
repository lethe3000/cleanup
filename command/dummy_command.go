package command
import (
    "github.com/urfave/cli"
    "fmt"
)

func NewDummyCommand() cli.Command {
	return cli.Command{
		Name:      "dummy",
		Usage:     "dummy command",
		ArgsUsage: " ",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "data-dir", Value: "", Usage: "Path to the etcd data dir"},
			cli.StringFlag{Name: "wal-dir", Value: "", Usage: "Path to the etcd wal dir"},
			cli.StringFlag{Name: "backup-dir", Value: "", Usage: "Path to the backup dir"},
			cli.StringFlag{Name: "backup-wal-dir", Value: "", Usage: "Path to the backup wal dir"},
		},
		Action: handleDummy,
	}
}

// handleBackup handles a request that intends to do a backup.
func handleDummy(c *cli.Context) error {
    fmt.Println("Dummy command:")
    fmt.Println(c.GlobalString("docker-entry"))
//    log.Println("Dummy command")
//	var srcWAL string
//	var destWAL string
//
//	srcSnap := path.Join(c.String("data-dir"), "member", "snap")
//	destSnap := path.Join(c.String("backup-dir"), "member", "snap")
//
//	if c.String("wal-dir") != "" {
//		srcWAL = c.String("wal-dir")
//	} else {
//		srcWAL = path.Join(c.String("data-dir"), "member", "wal")
//	}
//
//	if c.String("backup-wal-dir") != "" {
//		destWAL = c.String("backup-wal-dir")
//	} else {
//		destWAL = path.Join(c.String("backup-dir"), "member", "wal")
//	}
//
//	if err := fileutil.CreateDirAll(destSnap); err != nil {
//		log.Fatalf("failed creating backup snapshot dir %v: %v", destSnap, err)
//	}
//	ss := snap.New(srcSnap)
//	snapshot, err := ss.Load()
//	if err != nil && err != snap.ErrNoSnapshot {
//		log.Fatal(err)
//	}
//	var walsnap walpb.Snapshot
//	if snapshot != nil {
//		walsnap.Index, walsnap.Term = snapshot.Metadata.Index, snapshot.Metadata.Term
//		newss := snap.New(destSnap)
//		if err = newss.SaveSnap(*snapshot); err != nil {
//			log.Fatal(err)
//		}
//	}
//
//	w, err := wal.OpenForRead(srcWAL, walsnap)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer w.Close()
//	wmetadata, state, ents, err := w.ReadAll()
//	switch err {
//	case nil:
//	case wal.ErrSnapshotNotFound:
//		fmt.Printf("Failed to find the match snapshot record %+v in wal %v.", walsnap, srcWAL)
//		fmt.Printf("etcdctl will add it back. Start auto fixing...")
//	default:
//		log.Fatal(err)
//	}
//	var metadata etcdserverpb.Metadata
//	pbutil.MustUnmarshal(&metadata, wmetadata)
//	idgen := idutil.NewGenerator(0, time.Now())
//	metadata.NodeID = idgen.Next()
//	metadata.ClusterID = idgen.Next()
//
//	neww, err := wal.Create(destWAL, pbutil.MustMarshal(&metadata))
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer neww.Close()
//	if err := neww.Save(state, ents); err != nil {
//		log.Fatal(err)
//	}
//	if err := neww.SaveSnapshot(walsnap); err != nil {
//		log.Fatal(err)
//	}
//
	return nil
}

