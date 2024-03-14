package postgres

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/stretchr/testify/assert"
)

func TestINodeRepo_queryInsert(t *testing.T) {
	type fields struct {
		db *sql.DB
	}
	type args struct {
		inode *entity.INode
	}

	blockID1 := uuid.New()
	blockID2 := uuid.New()

	nodeID1 := "A"
	nodeID2 := "B"

	blockTarget := []*entity.BlockTarget{
		{
			ID:      blockID1,
			Size:    1,
			NodeIDs: []string{nodeID1, nodeID2},
		},
		{
			ID:      blockID2,
			Size:    1,
			NodeIDs: []string{nodeID1, nodeID2},
		},
	}

	iNodeID := uuid.New()
	iNode := &entity.INode{}
	iNode.SetID(iNodeID)
	iNode.SetBlocks(blockTarget)
	iNode.SetAllBlockIds([]uuid.UUID{blockID1, blockID2})

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   []interface{}
		wantErr bool
	}{
		{
			name: "should be ok",
			args: args{
				inode: iNode,
			},
			want: `INSERT INTO i_nodes_blocks (i_node_id,blocks_id,blocks_index,node_id,size) VALUES (($1, $2, $3, $4, $5), ($6, $7, $8, $9, $10), ($11, $12, $13, $14, $15), ($16, $17, $18, $19, $20))`,
			want1: []interface{}{
				iNodeID.String(),
				blockID1.String(),
				0,
				nodeID1,
				uint64(1),
				iNodeID.String(),
				blockID1.String(),
				0,
				nodeID2,
				uint64(1),
				iNodeID.String(),
				blockID2.String(),
				1,
				nodeID1,
				uint64(1),
				iNodeID.String(),
				blockID2.String(),
				1,
				nodeID2,
				uint64(1),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &INodeRepo{
				db: tt.fields.db,
			}
			got, got1, err := i.queryInsert(tt.args.inode)
			if (err != nil) != tt.wantErr {
				t.Errorf("INodeRepo.queryInsert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.want1, got1)
		})
	}
}

func TestGet(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	// Create a new instance of your repository
	iNodeRepo := &INodeRepo{db: db}
	blockId1 := uuid.New()
	blockId2 := uuid.New()
	nodeID1 := "A"
	nodeID2 := "B"
	nodeID3 := "C"

	// Prepare expected query and rows
	inodeID := uuid.New()
	rows := sqlmock.NewRows([]string{"blocks_id", "node_id", "blocks_index", "size"}).
		AddRow(blockId1.String(), nodeID1, 0, uint64(1)).
		AddRow(blockId1.String(), nodeID2, 0, uint64(1)).
		AddRow(blockId2.String(), nodeID1, 1, uint64(1)).
		AddRow(blockId2.String(), nodeID2, 1, uint64(1)).
		AddRow(blockId2.String(), nodeID3, 1, uint64(1))

	// Expectation for the query
	mock.ExpectQuery(regexp.QuoteMeta("SELECT FROM i_nodes_blocks (blocks_id,node_id,blocks_index,size) WHERE i_node_id = $1")).WithArgs(inodeID.String()).WillReturnRows(rows)

	// Call the Get function
	et, err := iNodeRepo.Get(context.Background(), inodeID, nil)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Assert the returned entity
	expectedBlocks := []*entity.BlockTarget{
		{ID: blockId1, NodeIDs: []string{nodeID1, nodeID2}, Size: uint64(1)},
		{ID: blockId2, NodeIDs: []string{nodeID1, nodeID2, nodeID3}, Size: uint64(1)},
	}
	assert.Equal(t, expectedBlocks, et.GetBlocks())

	// Check if all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestCreate(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer db.Close()

	// Create a new instance of your repository
	iNodeRepo := &INodeRepo{db: db}

	// Prepare a sample INode entity
	blockID1 := uuid.New()
	blockID2 := uuid.New()

	nodeID1 := "A"
	nodeID2 := "B"

	blockTarget := []*entity.BlockTarget{
		{
			ID:      blockID1,
			NodeIDs: []string{nodeID1, nodeID2},
			Size:    1,
		},
		{
			ID:      blockID2,
			NodeIDs: []string{nodeID1, nodeID2},
			Size:    1,
		},
	}

	iNodeID := uuid.New()
	iNode := &entity.INode{}
	iNode.SetID(iNodeID)
	iNode.SetBlocks(blockTarget)
	iNode.SetAllBlockIds([]uuid.UUID{blockID1, blockID2})

	// Prepare expected query and values
	expectedQuery := `INSERT INTO i_nodes_blocks (i_node_id,blocks_id,blocks_index,node_id,size) VALUES (($1, $2, $3, $4, $5), ($6, $7, $8, $9, $10), ($11, $12, $13, $14, $15), ($16, $17, $18, $19, $20))`
	expectedValues := []driver.Value{
		iNodeID.String(),
		blockID1.String(),
		0,
		nodeID1,
		uint64(1),
		iNodeID.String(),
		blockID1.String(),
		0,
		nodeID2,
		uint64(1),
		iNodeID.String(),
		blockID2.String(),
		1,
		nodeID1,
		uint64(1),
		iNodeID.String(),
		blockID2.String(),
		1,
		nodeID2,
		uint64(1),
	}

	// Expectation for the query
	mock.ExpectExec(regexp.QuoteMeta(expectedQuery)).WithArgs(expectedValues...).WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the Create function
	err = iNodeRepo.Create(context.Background(), iNode, nil)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Check if all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
