package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/google/uuid"
	pkgRepoTr "github.com/hdkef/hadoop/pkg/repository/transactionable"
	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/hdkef/hadoop/services/nameNode/repository"
)

const (
	queryGetByINodeID = "SELECT FROM i_nodes_blocks (blocks_id,node_id,blocks_index,size) WHERE i_node_id = $1"
	queryDeleteINode  = "DELETE FROM i_nodes_blocks WHERE i_node_id = $1"
)

type INodeRepo struct {
	db *sql.DB
}

// Create implements repository.INodeRepo.
func (i *INodeRepo) Create(ctx context.Context, inode *entity.INode, tx *pkgRepoTr.Transactionable) error {

	q, val, err := i.queryInsert(inode)
	if err != nil {
		return err
	}

	// if use tx
	if tx != nil {

		stmt, err := tx.Tx.PrepareContext(ctx, q)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(val...)
		if err != nil {
			return err
		}

		return nil
	}

	// else
	_, err = i.db.Exec(q, val...)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements repository.INodeRepo.
func (i *INodeRepo) Delete(ctx context.Context, inodeID uuid.UUID, tx *pkgRepoTr.Transactionable) error {
	// if use tx
	if tx != nil {

		stmt, err := tx.Tx.PrepareContext(ctx, queryDeleteINode)
		if err != nil {
			return err
		}
		defer stmt.Close()

		_, err = stmt.Exec(inodeID)
		if err != nil {
			return err
		}

		return nil
	}

	// else
	_, err := i.db.Exec(queryDeleteINode, inodeID)
	if err != nil {
		return err
	}

	return nil
}

// Get implements repository.INodeRepo.
func (i *INodeRepo) Get(ctx context.Context, inodeID uuid.UUID, tx *pkgRepoTr.Transactionable) (*entity.INode, error) {
	et := &entity.INode{}
	et.SetID(inodeID)

	var rows *sql.Rows
	var err error

	if tx != nil {
		rows, err = tx.Tx.QueryContext(ctx, queryGetByINodeID, inodeID.String())
	} else {
		rows, err = i.db.QueryContext(ctx, queryGetByINodeID, inodeID.String())
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	blocksIDNodeIds := make(map[uuid.UUID][]string)
	blockTargets := make(map[uint32]*entity.BlockTarget)

	// Iterate over the rows returned by the query
	for rows.Next() {
		var blockID uuid.UUID
		var NodeID string
		var Index uint32
		var size uint64

		err := rows.Scan(&blockID, &NodeID, &Index, &size)
		if err != nil {
			return nil, err
		}

		// Populate the maps
		blocksIDNodeIds[blockID] = append(blocksIDNodeIds[blockID], NodeID)
		blockTargets[Index] = &entity.BlockTarget{
			ID:   blockID,
			Size: size,
		}
	}

	// Construct allBlockIds and blocks slice
	allBlockIds := make([]uuid.UUID, len(blockTargets))
	blocks := make([]*entity.BlockTarget, len(blockTargets))

	// Populate allBlockIds and blocks slice
	for i, b := range blockTargets {
		nodeIds := blocksIDNodeIds[b.ID]
		allBlockIds[i] = b.ID
		blocks[i] = &entity.BlockTarget{
			ID:      b.ID,
			NodeIDs: nodeIds,
			Size:    b.Size,
		}
	}

	// Set the attributes of et
	et.SetAllBlockIds(allBlockIds)
	et.SetBlocks(blocks)

	return et, nil
}

func NewINodeRepo(db *sql.DB) repository.INodeRepo {

	if db == nil {
		panic("db is nil")
	}

	return &INodeRepo{
		db: db,
	}
}

func (i *INodeRepo) queryInsert(inode *entity.INode) (string, []interface{}, error) {

	if len(inode.GetAllBlockIds()) != len(inode.GetBlocks()) {
		return "", nil, errors.New("unmatch size of blocks & blocks id")
	}

	query := "INSERT INTO i_nodes_blocks (i_node_id,blocks_id,blocks_index,node_id,size) VALUES "

	var placeholders []string
	var values []interface{}

	idx := 0

	for i := range inode.GetAllBlockIds() {
		blocks := inode.GetBlocks()[i]
		for _, nodeID := range blocks.NodeIDs {
			values = append(values, inode.GetID(), blocks.ID, i, nodeID, blocks.Size)
			placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d, $%d)", (idx*5)+1, (idx*5)+2, (idx*5)+3, (idx*5)+4, (idx*5)+5))
			idx++
		}
	}

	query += strings.Join(placeholders, ", ")

	return query, values, nil
}
