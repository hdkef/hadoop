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

const queryGetByINodeID = "SELECT FROM i_nodes_blocks (blocks_id,node_id,blocks_index) WHERE i_node_id = $1"

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
	}

	// else
	_, err = i.db.Exec(q, val)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements repository.INodeRepo.
func (i *INodeRepo) Delete(ctx context.Context, inodeID uuid.UUID, tx *pkgRepoTr.Transactionable) error {
	panic("unimplemented")
}

// Get implements repository.INodeRepo.
func (i *INodeRepo) Get(ctx context.Context, inodeID uuid.UUID, tx *pkgRepoTr.Transactionable) (et *entity.INode, err error) {

	et = &entity.INode{}
	et.SetID(inodeID)

	var rows *sql.Rows

	// if use tx
	if tx != nil {
		stmt, err := tx.Tx.PrepareContext(ctx, queryGetByINodeID)
		if err != nil {
			return nil, err
		}
		defer stmt.Close()

		rows, err = stmt.QueryContext(ctx, inodeID)
		if err != nil {
			return nil, err
		}
	} else {
		// else
		rows, err = i.db.QueryContext(ctx, queryGetByINodeID, inodeID)
		if err != nil {
			return nil, err
		}
	}

	if rows == nil {
		err = errors.New("rows not exist")
		return
	}
	defer rows.Close()

	blocksIDNodeIds := make(map[uuid.UUID][]string)
	blockIDsIndex := make(map[uint32]uuid.UUID)

	for rows.Next() {
		var blockID uuid.UUID
		NodeID := ""
		Index := uint32(0)

		err := rows.Scan(&blockID, &NodeID, &Index)
		if err != nil {
			return nil, err
		}

		blocksIDNodeIds[blockID] = append(blocksIDNodeIds[blockID], NodeID)
		blockIDsIndex[Index] = blockID
	}

	allBlockIds := []uuid.UUID{}
	blocks := []*entity.BlockTarget{}

	for i := 0; i < len(blockIDsIndex); i++ {

		blockID := blockIDsIndex[uint32(i)]
		nodeIds := blocksIDNodeIds[blockID]

		allBlockIds = append(allBlockIds, blockID)

		blocks = append(blocks, &entity.BlockTarget{
			ID:      blockID,
			NodeIDs: nodeIds,
		})
	}

	et.SetAllBlockIds(allBlockIds)
	et.SetBlocks(blocks)

	return

}

func NewINodeRepo(db *sql.DB) repository.INodeRepo {
	return &INodeRepo{
		db: db,
	}
}

func (i *INodeRepo) queryInsert(inode *entity.INode) (string, []interface{}, error) {

	if len(inode.GetAllBlockIds()) != len(inode.GetBlocks()) {
		return "", nil, errors.New("unmatch size of blocks & blocks id")
	}

	query := "INSERT INTO i_nodes_blocks (i_node_id,blocks_id,blocks_index,node_id) VALUES "

	var placeholders []string
	var values []interface{}

	for i := range inode.GetAllBlockIds() {
		blocks := inode.GetBlocks()[i]
		for _, nodeID := range blocks.NodeIDs {
			values = append(values, inode.GetID().String(), blocks.ID.String(), i, nodeID)
			placeholders = append(placeholders, fmt.Sprintf("($%d, $%d, $%d, $%d)", (i*4)+1, (i*4)+2, (i*4)+3, (i*4)+4))
		}
	}

	query += fmt.Sprintf("(%s)", strings.Join(placeholders, ", "))

	return query, values, nil
}
