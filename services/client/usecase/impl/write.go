package impl

import (
	"context"
	"sync"

	pkgEt "github.com/hdkef/hadoop/pkg/entity"
	"github.com/hdkef/hadoop/pkg/helper"
	"github.com/hdkef/hadoop/services/client/entity"
	"github.com/hdkef/hadoop/services/client/service"
	"github.com/hdkef/hadoop/services/client/usecase"
	"golang.org/x/sync/errgroup"
)

type WriteUsecaseImpl struct {
	dataNodeService service.DataNodeService
}

func NewWriteUsecase(dataNodeService service.DataNodeService) usecase.WriteUsecase {
	return &WriteUsecaseImpl{
		dataNodeService: dataNodeService,
	}
}

// Write implements usecase.WriteUsecase.
func (w *WriteUsecaseImpl) Create(ctx context.Context, dto *entity.CreateDto, chProgress chan entity.CreateStreamRes) {

	// hash file

	// TODO req write to nameNode
	queryResult := entity.QueryNodeTarget{}

	totalBlock := queryResult.GetTotalBlock()

	// tokenizer
	blocksData := dto.Tokenize(totalBlock)
	mtx := &sync.Mutex{}
	progress := float32(0)

	errGroup := errgroup.Group{}

	for i := 0; i < totalBlock; i++ {

		errGroup.Go(func() error {

			// compress each block
			compressed, err := helper.Compress(blocksData[i])
			if err != nil {
				return err
			}

			blockID := queryResult.GetBlockID(i)
			nodeTarget := queryResult.GetNodeTarget(blockID)

			// execute replication for each blocks
			replicaDto := entity.ReplicateNextNodeDto{}
			replicaDto.SetINodeID(queryResult.GetINodeID())
			replicaDto.SetBlockID(blockID)
			replicaDto.SetBlocksData(compressed)
			replicaDto.SetCurrentReplicated(0)

			nextNode := &pkgEt.NodeInfo{}
			nextNode.SetNodeID(nodeTarget[0].GetNodeID())
			nextNode.SetAddress(nodeTarget[0].GetNodeAddress())
			nextNode.SetGRPCPort(nodeTarget[0].GetNodeGrpcPort())
			nextNode.SetReplicationStatus(pkgEt.REPLICATION_STATUS_PENDING)

			replicaDto.SetNextNode(nextNode)

			targetNode := []*pkgEt.NodeInfo{}

			for _, v := range nodeTarget {

				node := pkgEt.NodeInfo{}
				node.SetAddress(v.GetNodeAddress())
				node.SetGRPCPort(v.GetNodeGrpcPort())
				node.SetNodeID(v.GetNodeID())
				node.SetReplicationStatus(pkgEt.REPLICATION_STATUS_PENDING)

				targetNode = append(targetNode, &node)
			}

			replicaDto.SetReplicationNodeTarget(targetNode)

			err = w.dataNodeService.ReplicateNextNode(ctx, &replicaDto)

			// if success, increment progress and send progress info
			mtx.Lock()
			defer mtx.Unlock()
			progress += 100.0 / float32(totalBlock)
			p := entity.CreateStreamRes{}
			p.SetProgress(uint8(progress))
			chProgress <- p

			if err != nil {
				return err
			}

			return nil
		})

	}

	err := errGroup.Wait()
	defer close(chProgress)
	if err != nil {
		p := entity.CreateStreamRes{}
		p.SetError(err)
		chProgress <- p
		return
	}

	// TODO
	// commit result

}
