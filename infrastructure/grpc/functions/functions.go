package functions

import (
	"context"
	"go-challenger/adapter/controller"
	"go-challenger/core/usecase"
	"go-challenger/core/usecase/input"

	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcFunctions struct {
	delC    controller.DeleteController
	createC controller.SaveController
	createAllC controller.SaveAllController
	updateC controller.UpdateController
	getC    controller.GetController
}

func NewGrpcFunctions(
	delUc usecase.DeleteByIdUseCaseInterface,
	createUc  usecase.SaveUseCaseInterface,
	saveAllUc  usecase.SaveAllUseCaseInterface,
	updateUc usecase.UpdateUseCaseInterface,
	getUc usecase.FindByIdUseCaseInterface,
) CRUDServiceServer {
	return &GrpcFunctions{
		delC:    controller.NewDeleteController(delUc),
		createC: controller.NewSaveController(createUc),
		createAllC: controller.NewSaveAllController(saveAllUc),
		updateC: controller.NewSaveUpdateController(updateUc),
		getC:    controller.NewGetController(getUc),
	}
}

func (f *GrpcFunctions) Create(ctx context.Context, req *CreateRequest) (*IdResponse, error) {
	i := input.TaskInput{Id: "", Name: req.Name, Status: req.Status}
	id, code, err := f.createC.Execute(ctx, &i)
	if err != nil {
		return nil, status.Errorf(code, err.Error())
	}
	return &IdResponse{Id: id}, err
}

func (f *GrpcFunctions) Read(ctx context.Context, req *IdRequest) (*TaskResponse, error) {
	out, code, err := f.getC.Execute(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(code, err.Error())
	}
	return &TaskResponse{Id: out.Id, Status: out.Status, Name: out.Name}, err
}

func (f *GrpcFunctions) Update(ctx context.Context, req *UpdateRequest) (*TaskResponse, error) {
	i := input.TaskInput{Id: req.Id, Name: req.Name, Status: req.Status}
	out, code, err := f.updateC.Execute(ctx, &i)
	if err != nil {
		return nil, status.Errorf(code, err.Error())
	}
	return &TaskResponse{Id: out.Id, Name: out.Name, Status: out.Status}, err
}
func (f *GrpcFunctions) Delete(ctx context.Context, req *IdRequest) (*emptypb.Empty, error) {
	code, err := f.delC.Execute(ctx, req.Id)
	if err != nil {
		return nil, status.Errorf(code, err.Error())
	}
	return &emptypb.Empty{}, err
}

func (f *GrpcFunctions) BatchCreate(ctx context.Context,req *BatchRequest) (*emptypb.Empty, error){
	var list []input.TaskInput
	for _, item := range req.Items{
		list = append(list, input.TaskInput{Name: item.Name, Status: item.Status})
	}
	code, err := f.createAllC.Execute(ctx, &list)
	if err != nil {
		return nil, status.Errorf(code, err.Error())
	}
	return &emptypb.Empty{}, err
}

func (f *GrpcFunctions) mustEmbedUnimplementedCRUDServiceServer() {
}
