package middleware

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"logger/app/logging"
	"pkg/errors"
	"pkg/errors/pbError"
	"pkg/panicRecover"
)

func LoggingError(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ any, err error) {

	defer panicRecover.PanicRecover(func(er error) {
		logging.GetLogger().Panic(er)
		err = convertCustomErrToGrpcErr(er)
		return
	})

	resp, err := handler(ctx, req)

	if err != nil {

		// Логгируем ошибку в соответствии с ее типом
		logging.DefaultErrorLoggerFunc(err)

		// Конвертируем ошибку в gRPC ошибку
		return resp, convertCustomErrToGrpcErr(err)
	}

	return resp, err
}

func convertCustomErrToGrpcErr(defErr error) error {

	err, ok := defErr.(errors.CustomError)
	if !ok {
		return defErr
	}

	details := &pbError.CustomError{
		ErrorType:   uint32(errors.GetType(err)),
		HumanText:   err.HumanText,
		DevelopText: err.DevelopText,
		Err:         err.Error(),
		Path:        err.Path,
		Context:     err.Context,
	}

	st := status.Newf(
		codes.Code(err.ErrorType),
		err.HumanText,
	)

	st, _ = st.WithDetails(details)

	return st.Err()
}
