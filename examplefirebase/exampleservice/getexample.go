package exampleservice

import (
	"context"
	"fmt"

	connect_go "github.com/bufbuild/connect-go"

	"github.com/lcmaguire/protoc-gen-go-goo/examplefirebase/sample"
)

//const token = "eyJhbGciOiJSUzI1NiIsImtpZCI6ImQ3YjE5MTI0MGZjZmYzMDdkYzQ3NTg1OWEyYmUzNzgzZGMxYWY4OWYiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoibGlhbSBlbCBhdXN0cmFsaWFubyIsInBpY3R1cmUiOiJodHRwczovL2xoMy5nb29nbGV1c2VyY29udGVudC5jb20vYS0vQU9oMTRHaUFjRWV4dEg4UXlRalgySjJXSGJITFItZU45TEdSWHRpa0Z3Qmg9czk2LWMiLCJzdG9yZXMiOnsic3RvcmVOYW1lIjoidGVzdCIsImFkbWluIjp0cnVlfSwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL3BiY2MtdGVzdC1lbnYiLCJhdWQiOiJwYmNjLXRlc3QtZW52IiwiYXV0aF90aW1lIjoxNjY4Mzc3MTc0LCJ1c2VyX2lkIjoiSlVOQ3RUa1BJN1oyazUzeE9obTRuSEs3ekJvMiIsInN1YiI6IkpVTkN0VGtQSTdaMms1M3hPaG00bkhLN3pCbzIiLCJpYXQiOjE2NjgzNzcxNzQsImV4cCI6MTY2ODM4MDc3NCwiZW1haWwiOiJtYWduYWxkaW5vcmluaG9AZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZ29vZ2xlLmNvbSI6WyIxMTQ0MDUxNTYyNjI2MTg1NTc3NDgiXSwiZW1haWwiOlsibWFnbmFsZGlub3JpbmhvQGdtYWlsLmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6Imdvb2dsZS5jb20ifX0.CJRqdmQMuhY8OZqeRede3z06jUw1FmB9-VvKGAWBiDAiUUblKI6_u3Ut9oyRfuZiaMyJu9v31Bg4Qxy8rdw3EuDKeJEy4H4ZIv6SWmoJufaDHYVjeXowLelcOgYtJZD7tYAaNWMUb7bhZhLDH1OkxJqs6cjXqajs5Fj1fJVklQCEoDSr4Qw6_CdWH4WHUK-mqOuWPM4ixAApAqDNKIfOSkCRCkDSR8UtnEemCDH-sbF59E60UW7InvIf_MndhNiEfv8ZBCgU3apajOd9DnUGVryq15TBvL5gTyAELpLWfsjfXguS1Souom0joVumzWtjcR4h-p8r3NxLxypbGmZ0wg"

// GetExample implements tutorial.ExampleService.GetExample.
func (s *Service) GetExample(ctx context.Context, req *connect_go.Request[sample.GetExampleRequest]) (*connect_go.Response[sample.Example], error) {
	/*authUnstripped := req.Header().Get("Authorization")
	reqJWT := strings.Split(authUnstripped, "Bearer ")
	if len(reqJWT) < 2 {
		return nil, connect_go.NewError(connect_go.CodeUnauthenticated, errors.New("err"))
	}

	_, err := s.auth.VerifyIDToken(ctx, reqJWT[1])
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeUnauthenticated, err)
	} */

	/*docSnap, err := s.firestore.Doc(req.Msg.Name).Get(ctx)
	if err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}

	if docSnap == nil || docSnap.Data() == nil {
		return nil, connect_go.NewError(connect_go.CodeNotFound, err)
	}

	res := &sample.Example{}
	if err := docSnap.DataTo(res); err != nil {
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}*/

	res, err := s.db.Get(ctx, req.Msg.Name)
	if err != nil {
		fmt.Println("asi")
		fmt.Println(res)
		return nil, connect_go.NewError(connect_go.CodeInternal, err)
	}

	return connect_go.NewResponse(res), nil
}

// eyJhbGciOiJSUzI1NiIsImtpZCI6ImQ3YjE5MTI0MGZjZmYzMDdkYzQ3NTg1OWEyYmUzNzgzZGMxYWY4OWYiLCJ0eXAiOiJKV1QifQ.eyJuYW1lIjoibGlhbSBlbCBhdXN0cmFsaWFubyIsInBpY3R1cmUiOiJodHRwczovL2xoMy5nb29nbGV1c2VyY29udGVudC5jb20vYS0vQU9oMTRHaUFjRWV4dEg4UXlRalgySjJXSGJITFItZU45TEdSWHRpa0Z3Qmg9czk2LWMiLCJzdG9yZXMiOnsic3RvcmVOYW1lIjoidGVzdCIsImFkbWluIjp0cnVlfSwiaXNzIjoiaHR0cHM6Ly9zZWN1cmV0b2tlbi5nb29nbGUuY29tL3BiY2MtdGVzdC1lbnYiLCJhdWQiOiJwYmNjLXRlc3QtZW52IiwiYXV0aF90aW1lIjoxNjY4Mzc3MTc0LCJ1c2VyX2lkIjoiSlVOQ3RUa1BJN1oyazUzeE9obTRuSEs3ekJvMiIsInN1YiI6IkpVTkN0VGtQSTdaMms1M3hPaG00bkhLN3pCbzIiLCJpYXQiOjE2NjgzNzcxNzQsImV4cCI6MTY2ODM4MDc3NCwiZW1haWwiOiJtYWduYWxkaW5vcmluaG9AZ21haWwuY29tIiwiZW1haWxfdmVyaWZpZWQiOnRydWUsImZpcmViYXNlIjp7ImlkZW50aXRpZXMiOnsiZ29vZ2xlLmNvbSI6WyIxMTQ0MDUxNTYyNjI2MTg1NTc3NDgiXSwiZW1haWwiOlsibWFnbmFsZGlub3JpbmhvQGdtYWlsLmNvbSJdfSwic2lnbl9pbl9wcm92aWRlciI6Imdvb2dsZS5jb20ifX0.CJRqdmQMuhY8OZqeRede3z06jUw1FmB9-VvKGAWBiDAiUUblKI6_u3Ut9oyRfuZiaMyJu9v31Bg4Qxy8rdw3EuDKeJEy4H4ZIv6SWmoJufaDHYVjeXowLelcOgYtJZD7tYAaNWMUb7bhZhLDH1OkxJqs6cjXqajs5Fj1fJVklQCEoDSr4Qw6_CdWH4WHUK-mqOuWPM4ixAApAqDNKIfOSkCRCkDSR8UtnEemCDH-sbF59E60UW7InvIf_MndhNiEfv8ZBCgU3apajOd9DnUGVryq15TBvL5gTyAELpLWfsjfXguS1Souom0joVumzWtjcR4h-p8r3NxLxypbGmZ0wg
