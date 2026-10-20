package main

import (
	"database/sql"
	"net"

	"github.com/johnataferreira/gRPC/internal/database"
	"github.com/johnataferreira/gRPC/internal/pb"
	"github.com/johnataferreira/gRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	// _"github.com/mattn/go-sqlite3"
	_ "modernc.org/sqlite"
)

func main() {
	//db, err := sql.Open("sqlite", "./data.db")
	db, err := sql.Open("sqlite", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}