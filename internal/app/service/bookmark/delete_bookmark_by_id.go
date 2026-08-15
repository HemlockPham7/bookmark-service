package bookmark

import "context"

func (s *bookmarkService) DeleteBookmarkByID(ctx context.Context, userID, bookmarkID string) error {
	return s.bookmarkRepository.DeleteBookmarkByID(ctx, userID, bookmarkID)
}
