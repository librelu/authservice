package register

import (
	"reflect"
	"testing"
)

func Test_validateEmail(t *testing.T) {
	type args struct {
		email string
	}

	testCases := []struct {
		name               string
		args               args
		wantResult         bool
		expectedErr        error
		customErrorMatcher func(err error, expectedErr error) bool
	}{
		{
			name: "success example",
			args: args{
				email: "abcd@gmailyahoo",
			},
			wantResult:         true,
			expectedErr:        nil,
			customErrorMatcher: nil,
		},
		{
			name: "success example",
			args: args{
				email: "abcd@gmail.yahoo",
			},
			wantResult:         true,
			expectedErr:        nil,
			customErrorMatcher: nil,
		},
		{
			name: "success example",
			args: args{
				email: "abcd@gmail.yahoo.com",
			},
			wantResult:         true,
			expectedErr:        nil,
			customErrorMatcher: nil,
		},
		{
			name: "success example",
			args: args{
				email: "abcd@gmail.yahoo.com",
			},
			wantResult:         true,
			expectedErr:        nil,
			customErrorMatcher: nil,
		},
		{
			name: "success example",
			args: args{
				email: "abcd123123@gmail.yahoo.com",
			},
			wantResult:         true,
			expectedErr:        nil,
			customErrorMatcher: nil,
		},
		{
			name: "success example",
			args: args{
				email: "abcd@gmail-yahoo.com",
			},
			wantResult: true,
		},
		{
			name: "failed example",
			args: args{
				email: "ç$€§/az@gmail.com",
			},
			wantResult: false,
			customErrorMatcher: func(err error, expectedErr error) bool {
				return err.Error() == "incorrect email ex: 1234@domain.com current input:ç$€§/az@gmail.com"
			},
		},
		{
			name: "failed example",
			args: args{
				email: "ç$€§/azalsdkfjk34ji1ouds0f989.asdfkj3.324.234.234.234.",
			},
			wantResult: false,
			customErrorMatcher: func(err error, expectedErr error) bool {
				return err.Error() == "incorrect email ex: 1234@domain.com current input:ç$€§/azalsdkfjk34ji1ouds0f989.asdfkj3.324.234.234.234."
			},
		},
		{
			name: "failed example",
			args: args{
				email: "abcd@gmail_yahoo.com",
			},
			wantResult: false,
			customErrorMatcher: func(err error, expectedErr error) bool {
				return err.Error() == "incorrect email ex: 1234@domain.com current input:abcd@gmail_yahoo.com"
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := validateEmail(tt.args.email)
			if tt.customErrorMatcher != nil {
				if !tt.customErrorMatcher(err, tt.expectedErr) {
					t.Errorf("validateEmail() error = %v, expectedErr %v", err, tt.expectedErr)
				}
			} else if !reflect.DeepEqual(err, tt.expectedErr) {
				t.Errorf("validateEmail() error = %v, expectedErr %v", err, tt.expectedErr)
			}
			if gotResult != tt.wantResult {
				t.Errorf("validateEmail() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func Test_validatePassword(t *testing.T) {
	type args struct {
		password string
	}

	testCases := []struct {
		name               string
		args               args
		wantResult         bool
		expectedErr        error
		customErrorMatcher func(err error, expectedErr error) bool
	}{
		{
			name: "Successful password",
			args: args{
				password: "123kasdfuv4asd",
			},
			wantResult:  true,
			expectedErr: nil,
		},
		{
			name: "failed password",
			args: args{
				password: "123kasd",
			},
			wantResult: false,
			customErrorMatcher: func(err error, expectedErr error) bool {
				return err.Error() == "the password should contains at lease 8 charactor"
			},
		},
	}
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := validatePassword(tt.args.password)
			if tt.customErrorMatcher != nil {
				if !tt.customErrorMatcher(err, tt.expectedErr) {
					t.Errorf("validatePassword() error = %v, expectedErr %v", err, tt.expectedErr)
				}
			} else if !reflect.DeepEqual(err, tt.expectedErr) {
				t.Errorf("validatePassword() error = %v, expectedErr %v", err, tt.expectedErr)
			}
			if gotResult != tt.wantResult {
				t.Errorf("validatePassword() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
