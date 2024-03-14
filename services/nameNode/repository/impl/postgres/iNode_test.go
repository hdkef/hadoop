package postgres

import (
	"database/sql"
	"testing"

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
			NodeIDs: []string{nodeID1, nodeID2},
		},
		{
			ID:      blockID2,
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
			want: `INSERT INTO i_nodes_blocks (i_node_id,blocks_id,blocks_index,node_id) VALUES (($1, $2, $3, $4), ($1, $2, $3, $4), ($5, $6, $7, $8), ($5, $6, $7, $8))`,
			want1: []interface{}{
				iNodeID.String(),
				blockID1.String(),
				0,
				nodeID1,
				iNodeID.String(),
				blockID1.String(),
				0,
				nodeID2,
				iNodeID.String(),
				blockID2.String(),
				1,
				nodeID1,
				iNodeID.String(),
				blockID2.String(),
				1,
				nodeID2,
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
