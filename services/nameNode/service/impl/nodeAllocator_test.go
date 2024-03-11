package impl

import (
	"sort"
	"testing"

	"github.com/hdkef/hadoop/services/nameNode/entity"
	"github.com/stretchr/testify/assert"
)

func TestNodeAllocator_Sort(t *testing.T) {
	type args struct {
		nodeStorage []*entity.NodeStorage
	}

	nodeA := &entity.NodeStorage{}
	nodeA.SetNodeID("A")
	nodeA.SetTotalStorage(5000)
	nodeA.SetAllocated()
	nodeB := &entity.NodeStorage{}
	nodeB.SetNodeID("B")
	nodeB.SetTotalStorage(4000)
	nodeC := &entity.NodeStorage{}
	nodeC.SetNodeID("C")
	nodeC.SetTotalStorage(4000)
	nodeC.SetAllocated()
	nodeD := &entity.NodeStorage{}
	nodeD.SetNodeID("D")
	nodeD.SetTotalStorage(1000)
	nodeE := &entity.NodeStorage{}
	nodeE.SetNodeID("E")
	nodeE.SetTotalStorage(150)

	nodeStorage1 := []*entity.NodeStorage{
		nodeC,
		nodeD,
		nodeE,
		nodeA,
		nodeB,
	}

	tests := []struct {
		name string
		args args
		want []*entity.NodeStorage
	}{
		{
			name: "should be OK",
			args: args{
				nodeStorage: nodeStorage1,
			},
			want: []*entity.NodeStorage{
				nodeB,
				nodeD,
				nodeE,
				nodeA,
				nodeC,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			sort.Sort(ByTotalStorage(tt.args.nodeStorage))

			for i := 0; i < len(tt.want); i++ {
				assert.Equal(t, tt.want[i].GetNodeID(), tt.args.nodeStorage[i].GetNodeID(), "must be %s got %s", tt.want[i].GetNodeID(), tt.args.nodeStorage[i].GetNodeID())
			}
		})
	}
}

func TestNodeAllocator_AllocateTC1(t *testing.T) {
	type args struct {
		nodeStorage       []*entity.NodeStorage
		replicationTarget uint32
		blockSplitTarget  uint32
		fileSize          uint64
	}

	nodeA := &entity.NodeStorage{}
	nodeA.SetNodeID("A")
	nodeA.SetTotalStorage(5000)
	nodeB := &entity.NodeStorage{}
	nodeB.SetNodeID("B")
	nodeB.SetTotalStorage(4000)
	nodeC := &entity.NodeStorage{}
	nodeC.SetNodeID("C")
	nodeC.SetTotalStorage(4000)
	nodeD := &entity.NodeStorage{}
	nodeD.SetNodeID("D")
	nodeD.SetTotalStorage(1000)
	nodeE := &entity.NodeStorage{}
	nodeE.SetNodeID("E")
	nodeE.SetTotalStorage(300)

	nodeStorage1 := []*entity.NodeStorage{
		nodeC,
		nodeD,
		nodeE,
		nodeA,
		nodeB,
	}

	wantNodeA := &entity.NodeStorage{}
	wantNodeA.SetNodeID("A")
	wantNodeA.SetTotalStorage(5000)
	wantNodeA.SetLeaseUsedStorage(820)
	wantNodeB := &entity.NodeStorage{}
	wantNodeB.SetNodeID("B")
	wantNodeB.SetLeaseUsedStorage(612)
	wantNodeB.SetTotalStorage(4000)
	wantNodeC := &entity.NodeStorage{}
	wantNodeC.SetNodeID("C")
	wantNodeC.SetTotalStorage(4000)
	wantNodeC.SetLeaseUsedStorage(820)
	wantNodeD := &entity.NodeStorage{}
	wantNodeD.SetNodeID("D")
	wantNodeD.SetTotalStorage(1000)
	wantNodeD.SetLeaseUsedStorage(616)
	wantNodeE := &entity.NodeStorage{}
	wantNodeE.SetNodeID("E")
	wantNodeE.SetTotalStorage(300)
	wantNodeE.SetLeaseUsedStorage(204)

	wantNode := []*entity.NodeStorage{
		wantNodeE,
		wantNodeD,
		wantNodeA,
		wantNodeC,
		wantNodeB,
	}

	wantBlockTarget := []*entity.BlockTarget{
		{
			Size: 204,
			NodeIDs: []string{
				"A", "C", "B",
			},
		},
		{
			Size: 204,
			NodeIDs: []string{
				"D", "E", "A",
			},
		},
		{
			Size: 204,
			NodeIDs: []string{
				"C", "B", "D",
			},
		},
		{
			Size: 204,
			NodeIDs: []string{
				"A", "C", "B",
			},
		},
		{
			Size: 208,
			NodeIDs: []string{
				"D", "A", "C",
			},
		},
	}

	tests := []struct {
		name    string
		args    args
		want    []*entity.BlockTarget
		want1   []*entity.NodeStorage
		wantErr bool
	}{
		{
			name: "should be OK",
			args: args{
				replicationTarget: 3,
				blockSplitTarget:  5,
				fileSize:          1024,
				nodeStorage:       nodeStorage1,
			},
			want1: wantNode,
			want:  wantBlockTarget,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NodeAllocator{}
			got, got1, err := n.Allocate(tt.args.nodeStorage, tt.args.replicationTarget, tt.args.blockSplitTarget, tt.args.fileSize)

			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			// assert block target
			for i := 0; i < len(tt.want); i++ {
				assert.Equal(t, tt.want[i].NodeIDs, got[i].NodeIDs, "must be %v got %v", tt.want[i].NodeIDs, got[i].NodeIDs)
				assert.Equal(t, tt.want[i].Size, got[i].Size, "must be %d got %d", tt.want[i].Size, got[i].Size)
			}

			// assert node storage updated
			for i := 0; i < len(tt.want1); i++ {
				assert.Equal(t, tt.want1[i].GetFreeSpace(), got1[i].GetFreeSpace(), "got %d must be %d", got1[i].GetFreeSpace(), tt.want1[i].GetFreeSpace())
				assert.Equal(t, tt.want1[i].GetNodeID(), got1[i].GetNodeID(), "must be %s got %s", tt.want1[i].GetNodeID(), got1[i].GetNodeID())
			}
		})
	}
}

func TestNodeAllocator_AllocateTC2(t *testing.T) {
	type args struct {
		nodeStorage       []*entity.NodeStorage
		replicationTarget uint32
		blockSplitTarget  uint32
		fileSize          uint64
	}

	nodeA := &entity.NodeStorage{}
	nodeA.SetNodeID("A")
	nodeA.SetTotalStorage(5000)
	nodeB := &entity.NodeStorage{}
	nodeB.SetNodeID("B")
	nodeB.SetTotalStorage(4000)
	nodeC := &entity.NodeStorage{}
	nodeC.SetNodeID("C")
	nodeC.SetTotalStorage(4000)
	nodeD := &entity.NodeStorage{}
	nodeD.SetNodeID("D")
	nodeD.SetTotalStorage(1000)
	nodeE := &entity.NodeStorage{}
	nodeE.SetNodeID("E")
	nodeE.SetTotalStorage(150)

	nodeStorage1 := []*entity.NodeStorage{
		nodeC,
		nodeD,
		nodeE,
		nodeA,
		nodeB,
	}

	wantNodeA := &entity.NodeStorage{}
	wantNodeA.SetNodeID("A")
	wantNodeA.SetTotalStorage(5000)
	wantNodeA.SetLeaseUsedStorage(820)
	wantNodeB := &entity.NodeStorage{}
	wantNodeB.SetNodeID("B")
	wantNodeB.SetLeaseUsedStorage(820)
	wantNodeB.SetTotalStorage(4000)
	wantNodeC := &entity.NodeStorage{}
	wantNodeC.SetNodeID("C")
	wantNodeC.SetTotalStorage(4000)
	wantNodeC.SetLeaseUsedStorage(820)
	wantNodeD := &entity.NodeStorage{}
	wantNodeD.SetNodeID("D")
	wantNodeD.SetTotalStorage(1000)
	wantNodeD.SetLeaseUsedStorage(612)
	wantNodeE := &entity.NodeStorage{}
	wantNodeE.SetNodeID("E")
	wantNodeE.SetTotalStorage(150)

	wantNode := []*entity.NodeStorage{
		wantNodeE,
		wantNodeA,
		wantNodeB,
		wantNodeC,
		wantNodeD,
	}

	wantBlockTarget := []*entity.BlockTarget{
		{
			Size: 204,
			NodeIDs: []string{
				"A", "C", "B",
			},
		},
		{
			Size: 204,
			NodeIDs: []string{
				"D", "A", "C",
			},
		},
		{
			Size: 204,
			NodeIDs: []string{
				"B", "D", "A",
			},
		},
		{
			Size: 204,
			NodeIDs: []string{
				"B", "C", "D",
			},
		},
		{
			Size: 208,
			NodeIDs: []string{
				"A", "B", "C",
			},
		},
	}

	tests := []struct {
		name    string
		args    args
		want    []*entity.BlockTarget
		want1   []*entity.NodeStorage
		wantErr bool
	}{
		{
			name: "should be OK with one node storage unavailable",
			args: args{
				replicationTarget: 3,
				blockSplitTarget:  5,
				fileSize:          1024,
				nodeStorage:       nodeStorage1,
			},
			want1: wantNode,
			want:  wantBlockTarget,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NodeAllocator{}
			got, got1, err := n.Allocate(tt.args.nodeStorage, tt.args.replicationTarget, tt.args.blockSplitTarget, tt.args.fileSize)

			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			// assert block target
			for i := 0; i < len(tt.want); i++ {
				assert.Equal(t, tt.want[i].NodeIDs, got[i].NodeIDs, "must be %v got %v", tt.want[i].NodeIDs, got[i].NodeIDs)
				assert.Equal(t, tt.want[i].Size, got[i].Size, "must be %d got %d", tt.want[i].Size, got[i].Size)
			}

			// assert node storage updated
			for i := 0; i < len(tt.want1); i++ {
				assert.Equal(t, tt.want1[i].GetFreeSpace(), got1[i].GetFreeSpace(), "must be %d got %d", tt.want1[i].GetFreeSpace(), got1[i].GetFreeSpace())
				assert.Equal(t, tt.want1[i].GetNodeID(), got1[i].GetNodeID(), "must be %s got %s", tt.want1[i].GetNodeID(), got1[i].GetNodeID())
			}
		})
	}
}

func TestNodeAllocator_AllocateTC3(t *testing.T) {
	type args struct {
		nodeStorage       []*entity.NodeStorage
		replicationTarget uint32
		blockSplitTarget  uint32
		fileSize          uint64
	}

	nodeA := &entity.NodeStorage{}
	nodeA.SetNodeID("A")
	nodeA.SetTotalStorage(5000)
	nodeB := &entity.NodeStorage{}
	nodeB.SetNodeID("B")
	nodeB.SetTotalStorage(1000)
	nodeC := &entity.NodeStorage{}
	nodeC.SetNodeID("C")
	nodeC.SetTotalStorage(210)
	nodeD := &entity.NodeStorage{}
	nodeD.SetNodeID("D")
	nodeD.SetTotalStorage(1000)
	nodeE := &entity.NodeStorage{}
	nodeE.SetNodeID("E")
	nodeE.SetTotalStorage(150)

	nodeStorage1 := []*entity.NodeStorage{
		nodeC,
		nodeD,
		nodeE,
		nodeA,
		nodeB,
	}

	tests := []struct {
		name    string
		args    args
		want    []*entity.BlockTarget
		want1   []*entity.NodeStorage
		wantErr bool
	}{
		{
			name: "err not available space node",
			args: args{
				replicationTarget: 3,
				blockSplitTarget:  5,
				fileSize:          1024,
				nodeStorage:       nodeStorage1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NodeAllocator{}
			got, got1, err := n.Allocate(tt.args.nodeStorage, tt.args.replicationTarget, tt.args.blockSplitTarget, tt.args.fileSize)

			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			// assert block target
			for i := 0; i < len(tt.want); i++ {
				assert.Equal(t, tt.want[i].NodeIDs, got[i].NodeIDs, "must be %v got %v", tt.want[i].NodeIDs, got[i].NodeIDs)
				assert.Equal(t, tt.want[i].Size, got[i].Size, "must be %d got %d", tt.want[i].Size, got[i].Size)
			}

			// assert node storage updated
			for i := 0; i < len(tt.want1); i++ {
				assert.Equal(t, tt.want1[i].GetFreeSpace(), got1[i].GetFreeSpace(), "must be %d got %d", tt.want1[i].GetFreeSpace(), got1[i].GetFreeSpace())
				assert.Equal(t, tt.want1[i].GetNodeID(), got1[i].GetNodeID(), "must be %s got %s", tt.want1[i].GetNodeID(), got1[i].GetNodeID())
			}
		})
	}
}

func TestNodeAllocator_AllocateTC4(t *testing.T) {
	type args struct {
		nodeStorage       []*entity.NodeStorage
		replicationTarget uint32
		blockSplitTarget  uint32
		fileSize          uint64
	}

	nodeA := &entity.NodeStorage{}
	nodeA.SetNodeID("A")
	nodeA.SetTotalStorage(1024)
	nodeB := &entity.NodeStorage{}
	nodeB.SetNodeID("B")
	nodeB.SetTotalStorage(1024)
	nodeC := &entity.NodeStorage{}
	nodeC.SetNodeID("C")
	nodeC.SetTotalStorage(1024)
	nodeD := &entity.NodeStorage{}
	nodeD.SetNodeID("D")
	nodeD.SetTotalStorage(50)
	nodeE := &entity.NodeStorage{}
	nodeE.SetNodeID("E")
	nodeE.SetTotalStorage(150)

	nodeStorage1 := []*entity.NodeStorage{
		nodeC,
		nodeD,
		nodeE,
		nodeA,
		nodeB,
	}

	wantNodeA := &entity.NodeStorage{}
	wantNodeA.SetNodeID("A")
	wantNodeA.SetTotalStorage(1024)
	wantNodeA.SetLeaseUsedStorage(1024)
	wantNodeB := &entity.NodeStorage{}
	wantNodeB.SetNodeID("B")
	wantNodeB.SetLeaseUsedStorage(1024)
	wantNodeB.SetTotalStorage(1024)
	wantNodeC := &entity.NodeStorage{}
	wantNodeC.SetNodeID("C")
	wantNodeC.SetTotalStorage(1024)
	wantNodeC.SetLeaseUsedStorage(1024)
	wantNodeD := &entity.NodeStorage{}
	wantNodeD.SetNodeID("D")
	wantNodeD.SetTotalStorage(50)
	wantNodeE := &entity.NodeStorage{}
	wantNodeE.SetNodeID("E")
	wantNodeE.SetTotalStorage(150)

	wantNode := []*entity.NodeStorage{
		wantNodeE,
		wantNodeD,
		wantNodeC,
		wantNodeA,
		wantNodeB,
	}

	wantBlockTarget := []*entity.BlockTarget{
		{
			Size: 204,
			NodeIDs: []string{
				"C", "A", "B",
			},
		},
		{
			Size: 204,
			NodeIDs: []string{
				"C", "A", "B",
			},
		},
		{
			Size: 204,
			NodeIDs: []string{
				"C", "A", "B",
			},
		},
		{
			Size: 204,
			NodeIDs: []string{
				"C", "A", "B",
			},
		},
		{
			Size: 208,
			NodeIDs: []string{
				"C", "A", "B",
			},
		},
	}

	tests := []struct {
		name    string
		args    args
		want    []*entity.BlockTarget
		want1   []*entity.NodeStorage
		wantErr bool
	}{
		{
			name: "should be OK with tight available space",
			args: args{
				replicationTarget: 3,
				blockSplitTarget:  5,
				fileSize:          1024,
				nodeStorage:       nodeStorage1,
			},
			want1: wantNode,
			want:  wantBlockTarget,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &NodeAllocator{}
			got, got1, err := n.Allocate(tt.args.nodeStorage, tt.args.replicationTarget, tt.args.blockSplitTarget, tt.args.fileSize)

			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
			}

			// assert block target
			for i := 0; i < len(tt.want); i++ {
				assert.Equal(t, tt.want[i].NodeIDs, got[i].NodeIDs, "must be %v got %v", tt.want[i].NodeIDs, got[i].NodeIDs)
				assert.Equal(t, tt.want[i].Size, got[i].Size, "must be %d got %d", tt.want[i].Size, got[i].Size)
			}

			// assert node storage updated
			for i := 0; i < len(tt.want1); i++ {
				assert.Equal(t, tt.want1[i].GetFreeSpace(), got1[i].GetFreeSpace(), "must be %d got %d", tt.want1[i].GetFreeSpace(), got1[i].GetFreeSpace())
				assert.Equal(t, tt.want1[i].GetNodeID(), got1[i].GetNodeID(), "must be %s got %s", tt.want1[i].GetNodeID(), got1[i].GetNodeID())
			}
		})
	}
}
