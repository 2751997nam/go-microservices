package main

import (
	"context"
	"log-service/cmd/logs"
	"log-service/data"
)

type LogServer struct {
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogReponse, error) {
	input := req.GetLogEntry()
	//write the log
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogReponse{
			Result: "failed",
		}

		return res, err
	}

	//return response
	res := &logs.LogReponse{
		Result: "logged!",
	}

	return res, nil
}
