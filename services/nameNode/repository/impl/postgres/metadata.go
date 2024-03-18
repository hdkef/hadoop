package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	pkgRepoTr "github.com/hdkef/hadoop/pkg/repository/transactionable"
	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/repository"
	"github.com/lib/pq"
)

type MetadataRepo struct {
	db *sql.DB
}

const (
	queryMetadataCheckPath = "SELECT EXISTS (SELECT 1 FROM metadata WHERE path = $1)"
	queryMetadataInsert    = "INSERT INTO metadata (parent_path,path,m_type,i_node_id,hash,all_block_ids) VALUES ($1,$2,$3,$4,$5,$6)"
	queryMetadataGetByPath = "SELECT parent_path,path,m_type,i_node_id,hash,all_block_ids FROM metadata where path = $1"
	queryMetadataDelete    = "DELETE FROM metadata where path = $1"
)

// CheckPath implements repository.MetadataRepo.
func (m *MetadataRepo) CheckPath(ctx context.Context, path string, tx *pkgRepoTr.Transactionable) bool {
	var exists bool
	var err error

	if tx != nil {
		err = tx.Tx.QueryRow(queryMetadataCheckPath, path).Scan(&exists)
		if err != nil {
			return false
		}
		return exists
	}
	err = m.db.QueryRow(queryMetadataCheckPath, path).Scan(&exists)

	if err != nil {
		return false
	}

	return exists
}

// Delete implements repository.MetadataRepo.
func (m *MetadataRepo) Delete(ctx context.Context, metadata *entity.Metadata, tx *pkgRepoTr.Transactionable) error {
	// if use tx
	if tx != nil {

		stmt, err := tx.Tx.PrepareContext(ctx, queryMetadataDelete)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(metadata.GetPath())
		if err != nil {
			return err
		}

		return nil
	}

	// else
	_, err := m.db.Exec(queryMetadataDelete, metadata.GetPath())
	if err != nil {
		return err
	}

	return nil
}

// Get implements repository.MetadataRepo.
func (m *MetadataRepo) Get(ctx context.Context, path string, tx *pkgRepoTr.Transactionable) (*entity.Metadata, error) {
	var row *sql.Row

	md := &entity.Metadata{}

	if tx != nil {
		row = tx.Tx.QueryRowContext(ctx, queryMetadataGetByPath, path)
	} else {
		row = m.db.QueryRowContext(ctx, queryMetadataGetByPath, path)
	}

	if row != nil {
		var parentPath string
		var path string
		var hash string
		var mType entity.MetadataType
		iNodeIdBytes := []byte{}
		allBlocksIDsBytes := [][]byte{}
		allBlockIDs := []uuid.UUID{}

		err := row.Scan(&parentPath, &path, &mType, &iNodeIdBytes, &hash, &allBlocksIDsBytes)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		md.SetHash(hash)
		md.SetParentPath(parentPath)
		md.SetPath(path)
		md.SetType(mType)

		iNodeId, err := uuid.FromBytes(iNodeIdBytes)
		if err != nil {
			return nil, err
		}
		md.SetINodeID(iNodeId)

		for _, v := range allBlocksIDsBytes {
			bId, err := uuid.FromBytes(v)
			if err != nil {
				return nil, err
			}
			allBlockIDs = append(allBlockIDs, bId)
		}
		md.SetAllBlockIds(allBlockIDs)

		return md, nil
	}

	return nil, errors.New("metadata not found")
}

// Touch implements repository.MetadataRepo.
func (m *MetadataRepo) Touch(ctx context.Context, et *entity.Metadata, tx *pkgRepoTr.Transactionable) error {

	var blockIDs []string
	for _, id := range et.GetAllBlockIds() {
		blockIDs = append(blockIDs, id.String())
	}

	// if use tx
	if tx != nil {

		stmt, err := tx.Tx.PrepareContext(ctx, queryMetadataInsert)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(et.GetParentPath(), et.GetPath(), et.GetType(), et.GetINodeID(), et.GetHash(), pq.Array(blockIDs))
		if err != nil {
			return err
		}

		return nil
	}

	_, err := m.db.Exec(queryMetadataInsert, et.GetParentPath(), et.GetPath(), et.GetType(), et.GetINodeID().String(), et.GetHash(), pq.Array(blockIDs))
	if err != nil {
		return err
	}
	return nil
}

func NewMetadataRepo(db *sql.DB) repository.MetadataRepo {
	if db == nil {
		panic("db is nil")
	}

	return &MetadataRepo{db: db}
}
