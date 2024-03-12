package grpc

import (
	"context"

	"github.com/google/uuid"
	nameNodeProto "github.com/hdkef/hadoop/proto/nameNode"
)

// UpdateTransaction implements nameNode.NameNodeServer.
func (h *handler) CommitTransactions(ctx context.Context, req *nameNodeProto.CommitTransactionsReq) (*nameNodeProto.CommitTransactionsRes, error) {

	// retrieve request
	statusSuccess := false
	trID, err := uuid.FromBytes(req.GetTransactionID())
	if err != nil {
		return nil, err
	}

	if req.GetStatus() == nameNodeProto.CommitTransactionsReq_SUCCESS {
		statusSuccess = true
	}

	if statusSuccess {
		// if success, execute commit
		err = h.writeUC.CommitTransactions(ctx, trID)
		if err != nil {
			return nil, err
		}
		return &nameNodeProto.CommitTransactionsRes{}, nil
	} else {
		// else execute rollback

		err = h.writeUC.RollbackTransactions(ctx, trID)
		if err != nil {
			return nil, err
		}
		return &nameNodeProto.CommitTransactionsRes{}, nil
	}
}
