package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	pkgRepoTr "github.com/hdkef/hadoop/pkg/repository/transactionable"
	messageProto "github.com/hdkef/hadoop/proto/message"
	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/repository"
	"google.golang.org/protobuf/proto"
)

const (
	queryInsertTransactions = "INSERT INTO transactions (is_committed,created_at,lease_time_in_sec,protobuf_bytes) VALUES ($1,$2,$3,$4)"
	queryGetransactionsByID = "SELECT is_committed,created_at,lease_time_in_sec,protobuf_bytes FROM transactions WHERE id = $1"
)

type TransactionsRepo struct {
	db *sql.DB
}

// Add implements repository.TransactionsRepo.
func (t *TransactionsRepo) Add(ctx context.Context, et *entity.Transactions, tx *pkgRepoTr.Transactionable) error {

	createdAt := time.Now()

	protoBuf, err := transactionsToProto(et)
	if err != nil {
		return err
	}

	protoBufBytes, err := proto.Marshal(protoBuf)
	if err != nil {
		return err
	}

	// if use tx
	if tx != nil {

		stmt, err := tx.Tx.PrepareContext(ctx, queryInsertTransactions)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(false, createdAt, et.GetLeaseTimeInSecond(), protoBufBytes)
		if err != nil {
			return err
		}

		return nil
	}

	// else
	_, err = t.db.Exec(queryInsertTransactions, false, createdAt, et.GetLeaseTimeInSecond(), protoBufBytes)
	if err != nil {
		return err
	}

	return nil
}

// Commit implements repository.TransactionsRepo.
func (t *TransactionsRepo) Commit(ctx context.Context, transactionID uuid.UUID, tx *pkgRepoTr.Transactionable) error {
	panic("unimplemented")
}

// Get implements repository.TransactionsRepo.
func (t *TransactionsRepo) Get(ctx context.Context, transactionID uuid.UUID, tx *pkgRepoTr.Transactionable) (*entity.Transactions, error) {
	et := &entity.Transactions{}

	var row *sql.Row

	if tx != nil {
		row = tx.Tx.QueryRowContext(ctx, queryGetransactionsByID, transactionID.String())
	} else {
		row = t.db.QueryRowContext(ctx, queryGetransactionsByID, transactionID.String())
	}

	if row != nil {

		isCommited := false
		createdAt := time.Time{}
		leaseTimeInSecond := uint64(0)
		protoBufBytes := []byte{}
		protoBuf := &messageProto.Transactions{}

		err := row.Scan(&isCommited, &createdAt, &leaseTimeInSecond, &protoBufBytes)
		if err != nil {
			return nil, err
		}

		et.SetIsCommitted(isCommited)
		et.SetCreatedAt(createdAt)
		et.SetLeaseTimeInSecond(leaseTimeInSecond)

		err = proto.Unmarshal(protoBufBytes, protoBuf)
		if err != nil {
			return nil, err
		}

		err = protoToTransactions(et, protoBuf)
		if err != nil {
			return nil, err
		}

		return et, nil
	}

	return nil, errors.New("transactions not found")
}

// GetOneExpired implements repository.TransactionsRepo.
func (t *TransactionsRepo) GetOneExpired(ctx context.Context, tx *pkgRepoTr.Transactionable) (*entity.Transactions, error) {
	panic("unimplemented")
}

// RolledBack implements repository.TransactionsRepo.
func (t *TransactionsRepo) RolledBack(ctx context.Context, transactionID uuid.UUID, tx *pkgRepoTr.Transactionable) error {
	panic("unimplemented")
}

func transactionsToProto(et *entity.Transactions) (*messageProto.Transactions, error) {
	transactionProto := &messageProto.Transactions{
		Metadata: &messageProto.Metadata{
			ParentPath: et.GetMetadata().GetParentPath(),
			Path:       et.GetMetadata().GetPath(),
			Hash:       et.GetMetadata().GetHash(),
		},
	}

	inodeID, err := et.GetMetadata().GetINodeID().MarshalBinary()
	if err != nil {
		return nil, err
	}
	transactionProto.Metadata.INodeID = inodeID

	allBlocksID := [][]byte{}

	for _, v := range et.GetMetadata().GetAllBlockIds() {
		bID, err := v.MarshalBinary()
		if err != nil {
			return nil, err
		}
		allBlocksID = append(allBlocksID, bID)
	}
	transactionProto.Metadata.AllBlockIDs = allBlocksID

	switch et.GetMetadata().GetType() {
	case entity.METADATA_TYPE_DIR:
		transactionProto.Metadata.MType = messageProto.Metadata_DIR
	case entity.METADATA_TYPE_FILE:
		transactionProto.Metadata.MType = messageProto.Metadata_FILE
	}

	switch et.GetAction() {
	case entity.TRANSACTION_ACTION_CREATE:
		transactionProto.Action = messageProto.Transactions_CREATE
	case entity.TRANSACTION_ACTION_UPDATE:
		transactionProto.Action = messageProto.Transactions_UPDATE
	case entity.TRANSACTION_ACTION_DELETE:
		transactionProto.Action = messageProto.Transactions_DELETE
	}

	for _, v := range et.GetBlockTaret() {

		bID, err := v.ID.MarshalBinary()
		if err != nil {
			return nil, err
		}

		transactionProto.BlockTarget = append(transactionProto.BlockTarget, &messageProto.BlockTarget{
			ID:      bID,
			Size:    v.Size,
			NodeIDs: v.NodeIDs,
		})
	}

	return transactionProto, nil
}

func protoToTransactions(et *entity.Transactions, proto *messageProto.Transactions) error {

	md := &entity.Metadata{}
	md.SetHash(proto.GetMetadata().GetHash())
	md.SetParentPath(proto.GetMetadata().GetParentPath())
	md.SetPath(proto.GetMetadata().GetPath())

	iNodeID, err := uuid.FromBytes(proto.GetMetadata().GetINodeID())
	if err != nil {
		return err
	}

	md.SetINodeID(iNodeID)

	allBlockIds := []uuid.UUID{}

	for _, v := range proto.GetMetadata().GetAllBlockIDs() {
		bID, err := uuid.FromBytes(v)
		if err != nil {
			return err
		}
		allBlockIds = append(allBlockIds, bID)
	}
	md.SetAllBlockIds(allBlockIds)

	switch proto.GetMetadata().GetMType() {
	case messageProto.Metadata_DIR:
		md.SetType(entity.METADATA_TYPE_DIR)
	case messageProto.Metadata_FILE:
		md.SetType(entity.METADATA_TYPE_FILE)
	}

	et.SetMetadata(md)

	blockTarget := []*entity.BlockTarget{}

	for _, v := range proto.GetBlockTarget() {
		blockId, err := uuid.FromBytes(v.GetID())
		if err != nil {
			return err
		}

		blockTarget = append(blockTarget, &entity.BlockTarget{
			ID:      blockId,
			Size:    v.Size,
			NodeIDs: v.NodeIDs,
		})

	}

	et.SetBlockTarget(blockTarget)

	return nil
}

func NewTransactionsRepo(db *sql.DB) repository.TransactionsRepo {
	return &TransactionsRepo{
		db: db,
	}
}
