package classifications

import "context"

func (s *ServiceImpl) FindAllByEventId(ctx context.Context, eventId string) (*FindAllResponse, error) {
	classifications, err := s.classRepo.FindAllByEventId(ctx, eventId)

	if err != nil {
		return nil, err
	}

	return &FindAllResponse{
		StatusCode: 200,
		Message:    "Classifications found successfully",
		Data:       classifications,
	}, nil
}
