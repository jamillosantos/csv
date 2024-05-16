package service

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

type Service struct {
}

func NewService() Service {
	return Service{}
}

type RunRequest struct {
	Output      io.Writer
	Input       io.Reader
	SkipHeaders bool
	Separator   string
	Columns     string
}

func (s *Service) Run(_ context.Context, req RunRequest) error {
	r := csv.NewReader(req.Input)
	r.Comma = rune(req.Separator[0])

	var columns map[int]struct{}
	if req.Columns != "" {
		columnsStrings := lo.Map(strings.Split(req.Columns, ","), func(s string, _ int) string {
			return strings.TrimSpace(s)
		})
		columns = make(map[int]struct{}, len(columnsStrings))
		for _, column := range columnsStrings {
			columnInt, err := strconv.Atoi(column)
			if err != nil {
				return fmt.Errorf("failed parsing column number: %w", err)
			}
			columns[columnInt] = struct{}{}
		}
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if req.SkipHeaders {
			req.SkipHeaders = false
			continue
		}
		columnI := 1
		for i, field := range record {
			if len(columns) > 0 {
				if _, ok := columns[i+1]; !ok {
					continue
				}
			}

			if columnI > 1 {
				_, _ = req.Output.Write([]byte(req.Separator))
			}
			containsSeparator := strings.Contains(field, req.Separator)
			if containsSeparator {
				_, _ = req.Output.Write([]byte(`"`))
			}
			_, _ = req.Output.Write([]byte(field))
			if containsSeparator {
				_, _ = req.Output.Write([]byte(`"`))
			}
			columnI++
		}
		_, _ = req.Output.Write([]byte("\n"))
	}
	return nil
}
