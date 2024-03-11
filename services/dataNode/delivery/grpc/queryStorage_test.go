package grpc

import (
	"context"
	"syscall"
	"testing"

	dataNodeProto "github.com/hdkef/hadoop/proto/dataNode"
	"github.com/hdkef/hadoop/services/dataNode/config"
	"github.com/hdkef/hadoop/services/dataNode/usecase"
	"github.com/stretchr/testify/assert"
)

func Test_handler_QueryStorage(t *testing.T) {
	type fields struct {
		UnimplementedDataNodeServer dataNodeProto.UnimplementedDataNodeServer
		writeUC                     usecase.WriteUsecase
		cfg                         *config.Config
	}
	type args struct {
		in0 context.Context
		in1 *dataNodeProto.QueryStorageReq
	}

	path := "/"
	var stat syscall.Statfs_t

	syscall.Statfs(path, &stat)

	// Calculate total size in bytes
	total := stat.Blocks * uint64(stat.Bsize)
	// Calculate used size in bytes
	used := (stat.Blocks - stat.Bfree) * uint64(stat.Bsize)

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *dataNodeProto.QueryStorageRes
		wantErr bool
	}{
		{
			name: "should be ok",
			fields: fields{
				cfg: &config.Config{
					StorageRoot: path,
				},
			},
			want: &dataNodeProto.QueryStorageRes{
				TotalStorage:      total,
				ActualUsedStorage: used,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &handler{
				UnimplementedDataNodeServer: tt.fields.UnimplementedDataNodeServer,
				writeUC:                     tt.fields.writeUC,
				cfg:                         tt.fields.cfg,
			}
			got, err := g.QueryStorage(tt.args.in0, tt.args.in1)
			if tt.wantErr {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
