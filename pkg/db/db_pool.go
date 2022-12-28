package db

import (
	"math"
	"sync"

	"github.com/DeBankDeFi/db-replicator/pkg/utils"
	"github.com/DeBankDeFi/db-replicator/pkg/utils/pb"
	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/protobuf/proto"
)

const (
	LastBlockInfo = "rpl_last_bk"
	DBInfo        = "rpl_db_info"
)

type dbWrap struct {
	id     int32
	db     DB
	path   string
	dbType string
	isMeta bool
}

// DBPool is a pool of Chain's All DBs.
type DBPool struct {
	sync.RWMutex
	dbs      map[int32]dbWrap
	metaDBID int32
}

func NewDBPool() *DBPool {
	return &DBPool{
		dbs:      make(map[int32]dbWrap),
		metaDBID: math.MinInt32,
	}
}

// RegisterMetaDB registers a DB.
// if isMetaDB is true, this DB will be used to store msgHead.
func (p *DBPool) Register(id int32, path string, dbType string, db DB, isMetaDB bool) {
	p.Lock()
	defer p.Unlock()
	p.dbs[id] = dbWrap{
		id:     id,
		db:     db,
		path:   path,
		dbType: dbType,
		isMeta: isMetaDB,
	}
	if isMetaDB && p.metaDBID != math.MinInt32 {
		panic("meta db already registered")
	}
	if isMetaDB {
		p.metaDBID = id
	}
}

func (p *DBPool) Open(dbInfo *pb.DBInfo) (err error) {
	p.Lock()
	defer p.Unlock()
	if _, ok := p.dbs[dbInfo.Id]; ok {
		return nil
	}
	switch dbInfo.DbType {
	case "leveldb":
		for _, db := range p.dbs {
			if db.path == dbInfo.DbPath {
				return nil
			}
		}
		db, err := NewLDB(dbInfo.DbPath)
		if err != nil {
			return err
		}
		p.dbs[dbInfo.Id] = dbWrap{
			id:     dbInfo.Id,
			db:     db,
			path:   dbInfo.DbPath,
			dbType: dbInfo.DbType,
			isMeta: dbInfo.IsMeta,
		}
		if dbInfo.IsMeta && p.metaDBID != math.MinInt32 {
			return utils.ErrMetaDBAlreadyRegistered
		}
		if dbInfo.IsMeta {
			p.metaDBID = dbInfo.Id
		}
	}
	return nil
}

// GetBlockInfo returns the last BlockInfo from metaDB.
func (p *DBPool) GetBlockInfo() (header *pb.BlockInfo, err error) {
	p.RLock()
	defer p.RUnlock()
	if p.metaDBID == math.MinInt32 {
		return nil, utils.ErrNoMetaDBRegistered
	}
	db := p.dbs[p.metaDBID]
	lastMsgHeadBuf, err := db.db.Get([]byte(LastBlockInfo))
	if err != nil && err != leveldb.ErrNotFound {
		return nil, err
	}
	header = &pb.BlockInfo{
		BlockNum: -1,
		MsgOffset: -1,
	}
	if err == leveldb.ErrNotFound || len(lastMsgHeadBuf) == 0 {
		return header, nil
	}
	if err := proto.Unmarshal(lastMsgHeadBuf, header); err != nil {
		return nil, nil
	}
	return header, nil
}

func (p *DBPool) GetDBInfo() (info *pb.DBInfoList, err error) {
	p.RLock()
	defer p.RUnlock()
	if p.metaDBID == math.MinInt32 {
		return nil, utils.ErrNoMetaDBRegistered
	}
	info = &pb.DBInfoList{}
	for id, db := range p.dbs {
		info.DbInfos = append(info.DbInfos, &pb.DBInfo{
			Id:     id,
			DbType: db.dbType,
			DbPath: db.path,
			IsMeta: db.isMeta,
		})
	}
	return info, nil
}

func (p *DBPool) GetDB(id int32) (db DB, err error) {
	p.RLock()
	defer p.RUnlock()
	dbWrap, ok := p.dbs[id]
	if !ok {
		return nil, leveldb.ErrNotFound
	}
	return dbWrap.db, nil
}

func (p *DBPool) GetDBID(path string) (id int32, err error) {
	p.RLock()
	defer p.RUnlock()
	for _, db := range p.dbs {
		if db.path == path {
			return db.id, nil
		}
	}
	return -1, leveldb.ErrNotFound
}

func (p *DBPool) Marshal(batchs []BatchWithID) (batchItems []*pb.BatchItem, err error) {
	p.RLock()
	defer p.RUnlock()
	for _, item := range batchs {
		tmp := make([]byte, len(item.B.Dump()))
		copy(tmp, item.B.Dump())
		batchItems = append(batchItems, &pb.BatchItem{
			Id:   int32(item.ID),
			Data: tmp,
		})
	}
	return batchItems, nil
}

func (p *DBPool) WriteBlockInfo(header *pb.BlockInfo) (err error) {
	p.RLock()
	defer p.RUnlock()
	if p.metaDBID == math.MinInt32 {
		return utils.ErrNoMetaDBRegistered
	}
	lastMsgHeadBuf, err := proto.Marshal(header)
	if err != nil {
		return err
	}
	metaBatch := p.dbs[p.metaDBID].db.NewBatch()
	err = metaBatch.Put([]byte(LastBlockInfo), lastMsgHeadBuf)
	if err != nil {
		return err
	}
	err = metaBatch.Write()
	if err != nil {
		return err
	}
	return nil
}

func (p *DBPool) WriteBatchs(batchs []BatchWithID) (err error) {
	p.RLock()
	defer p.RUnlock()
	for _, item := range batchs {
		err = item.B.Write()
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *DBPool) WriteBatchItems(items []*pb.BatchItem) (err error) {
	p.RLock()
	defer p.RUnlock()
	for _, item := range items {
		db := p.dbs[item.Id]
		batch := db.db.NewBatch()
		if err := batch.Load(item.Data); err != nil {
			return err
		}
		if err := batch.Write(); err != nil {
			return err
		}
	}
	return nil
}
