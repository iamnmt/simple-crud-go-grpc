package service

import (
	"fmt"
	"log"
	"context"
	"time"

	"ms/server/errors"
	"ms/server/types"
)

var msgDb = make([]*types.MsgDb, 0)

type CrudService struct {
}

func (cs *CrudService) Create(ctx context.Context, req *types.MsgCreate) error {
	msgDb = append(msgDb, &types.MsgDb{
		Id:        req.GetId(),
		Content:   req.GetContent(),
		Timestamp: time.Now().Second(),
	})

	log.Printf("A new Msg with id = %d is created", req.GetId())
	return nil
}

func (cs *CrudService) Read(ctx context.Context, req *types.MsgRead) (string, error) {
	for _, m := range msgDb {
		if m.Id == req.GetId() {
			return m.Content, nil
		}
	}

	log.Printf("Failed to read Msg with id = %d", req.GetId())
	return "", errors.CustomError{Msg: fmt.Sprintf("Can not read Msg with id = %d", req.GetId())}
}

func (cs *CrudService) Update(ctx context.Context, req *types.MsgUpdate) error {
	for _, m := range msgDb {
		if m.Id == req.GetId() {
			m.Content = req.GetContent()
			log.Printf("The content of Msg with id = %d is updated", req.GetId())

			return nil
		}
	}

	log.Printf("Failed to update Msg with id = %d", req.GetId())
	return errors.CustomError{Msg: fmt.Sprintf("Can not update Msg with id = %d", req.GetId())}
}

func (cs *CrudService) Delete(ctx context.Context, req *types.MsgDelete) error {
	idx := -1
	for i, m := range msgDb {
		if m.Id == req.GetId() {
			idx = i
			break
		}
	}

	if idx == -1 {
		log.Printf("Failed to delete Msg with id = %d", req.GetId())
		return errors.CustomError{Msg: fmt.Sprintf("Can not delete Msg with id = %d", req.GetId())}
	}

	msgDb = append(msgDb[:idx], msgDb[idx+1:]...)

	return nil
}
