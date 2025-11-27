package jti

import "context"

func (j *JTIUseCase) RevokeJTI(ctx context.Context, jti string) error {
	return j.jtiRepo.DeleteJTI(ctx, jti)
}
