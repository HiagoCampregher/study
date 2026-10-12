package service

import (
	"context"
	"io"
	"study/grpc/internal/database"
	"study/grpc/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (s *CategoryService) CreateCategory(ctx context.Context, req *pb.CategoryRequest) (*pb.Category, error) {
	category, err := s.CategoryDB.Create(req.Name, req.Description)
	if err != nil {
		return nil, err
	}

	CategoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return CategoryResponse, nil
}

func (s *CategoryService) ListCategories(ctx context.Context, req *pb.Blank) (*pb.CategoryList, error) {
	categories, err := s.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	categoryList := &pb.CategoryList{
		Categories: []*pb.Category{},
	}

	for _, category := range categories {
		categoryList.Categories = append(categoryList.Categories, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return categoryList, nil
}

func (s *CategoryService) GetCategory(ctx context.Context, req *pb.CategoryGetRequest) (*pb.Category, error) {
	category, err := s.CategoryDB.Find(req.Id)
	if err != nil {
		return nil, err
	}

	categoryResponse := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return categoryResponse, nil
}

func (s *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {

	categories := &pb.CategoryList{}
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}

		if err != nil {
			return err
		}

		categoryResult, err := s.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		categoryResponse := &pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		}

		categories.Categories = append(categories.Categories, categoryResponse)
	}
}

func (s *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		categoryResult, err := s.CategoryDB.Create(category.Name, category.Description)
		if err != nil {
			return err
		}

		categoryResponse := &pb.Category{
			Id:          categoryResult.ID,
			Name:        categoryResult.Name,
			Description: categoryResult.Description,
		}

		if err := stream.Send(categoryResponse); err != nil {
			return err
		}
	}
}
