
package checkpoint_meta

import (
	"context"
    
    "meta/pb/checkpoint_meta"
    
)

type MetaServiceServer struct{}

func (server *MetaServiceServer) GetLatestVersion(context context.Context, request *checkpoint_meta.GetLatestVersionRequest) (response *checkpoint_meta.GetLatestVersionRequest, err error) {
	panic("implement me")
    return 
}

func (server *MetaServiceServer) UpdatePartition(context context.Context, request *checkpoint_meta.UpdatePartitionRequest) (response *checkpoint_meta.UpdatePartitionResponse, err error) {
	panic("implement me")
    return 
}

