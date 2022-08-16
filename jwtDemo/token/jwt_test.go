package token

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v4"
	"testing"
	"time"
)

/*
// json key 先后顺序不同生成的token会一样

	{
	  "iss": "www.darr-en1.top",
	  "sub": "darr_en1",
	  "exp": 1516242622,
	  "iat": 1516239022
	}
*/
const privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEoQIBAAKCAQBcgJ7S8Lx0XkxwW5XYb38MxCXgaCXug/WUU6fU8g5FHY+wHcSY
Hl4L+T7oRsE0mtDDNAWyPAu9nlwsEzh4nzpqCN/LPqd7Xfs7fyjJglM2rvoX19YH
mNuXtgKr67zrUE9QlOQ1ro9+HSSHSf/eE8RKNxSwC1Al8oyxXUeaqRR+2gPLP+8L
e4nb9xyUgrdVi5Kj1DA6J5F9/bpmkVLMvMid2Fpk9G19Xq3y1fgMLclqHpBipu8D
wvwayaPPM7eMmgUthLXS/JA5gWq7AyMruFehtOUNvY9tWVgwNaY966FmWxZeklGZ
2Vp5WGfoz1CKqaF3LbvwkqPKAjoAqfw+CXaJAgMBAAECggEABGR3Jj0YNc5kgtFW
hDiHJ+wIgGdxpanOd4Sha5My6kVFFW/UbgTDIn5rZkw10HGpiBtoSdZgGFj0t4MS
I2gsNezF87i96zXDQEgBf9AYc3hLl+Y/24x+oO2bot5G0kW1/nWTgJkRZ0OrAGGb
LMhzgbKr4NAotiSWWbW8OLPerUdOOUMHVeaFOrEGk1E3yC7+/2LC+X7W4eUAdzUM
OSWQoZxNw6G0Mw+8OAvVWNHrjOCHxhkCG+nNdUtaOe2nLzY8+MOXXIEiQ/Vpf71n
icb/PVHUqvzY55wrSh1FbCFIj4SmbAgM6SwcBxOBI3GOAe7PBdy03WwEZVul21mr
dl4QoQKBgQCvczso8oFYJVhd5KwNZWhOyeGUAG8+lTaD09UUNtwQZLf8vPF4iiG2
VYOGQKT284P0CJFKvjL/Kvxbyar42B6XaKOXAIXGbIvQFlsQwwJe3MRef7pcwU9j
7d+nJPLn0s95UNBv0fIm6lDHIEPmc3u3uUolJ4wcgSoVJLg+jILh8wKBgQCG+H3Q
jPqYRuMVMtEk4xOJvimyQCbaasMSCBTOMwQcUmLimOtHenuLtFnl5UK0Siz745AH
5dEG4epG7dpB7KHpktFQ1GXHaEwGJbvfzYUwWj1cCkSwAIQfLTsNdd6sWTRBhagj
fXRDv8R8BDW+CkeAD9gtUIVkStCj9FU2H01okwKBgAbpjU0hsMLeRcr9NWIZurBP
99ky4y8eBdXPxLdVKfpjXXRRpVnQZ1+dot7pMRahpXM52y+Mqsmu0d5z5Y6ERUBF
2Rwb4ylcIW4DU3cnl7JRFFN8yMTawNv5BTS6me9UOORYsL2XEelClggiV0p/BkLI
xAkOs9BzrkrC9ZdRRd7NAoGAdnMfn2YKCpezdIgrac7q96hu2WNZdbfaSDnVYHz/
L4fSBoEjgpBoWlvFxW9lEepC/jBg172Fqx+axfK5578u/Vh+4Av24oCSr5ZwAeJQ
7/VSAjN0wm1Bhkbgm1iFOgXzSUluHO9dmJN0cU4UiBpnU2kUU/hliHXsEuc73Qqt
mYcCgYB6CmMIOGg3N20ipc4UX0qwoUli3Aw+3772C6C6boFm4HQRPY4a5eYCy8s2
Aiuu4h+NdMi+5C/npnZIHZv/kJLCFs92iLNB1ZCN7svsArnySk+wpAQ49H/VOBPg
D4n8S+5skX+sio2kGrqBpEgWFvOeDonK+oyGEV7Nc31dPPYU0A==
-----END RSA PRIVATE KEY-----`

func TestJWTTokenGen_GenerateToken(t *testing.T) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	if err != nil {
		t.Fatalf("cannot parse private key:%v", err)
	}

	type fields struct {
		privateKey *rsa.PrivateKey
		issuer     string
		nowFunc    func() time.Time
	}
	type args struct {
		code   string
		expire time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "GenerateToken",
			fields: fields{
				privateKey: key,
				issuer:     "www.darr-en1.top",
				nowFunc: func() time.Time {
					return time.Unix(1516239022, 0)
				},
			},
			args: args{
				code:   "darr_en1",
				expire: time.Hour,
			},
			want:    "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ3d3cuZGFyci1lbjEudG9wIiwic3ViIjoiZGFycl9lbjEiLCJleHAiOjE1MTYyNDI2MjIsImlhdCI6MTUxNjIzOTAyMn0.HInvEyJxa6Vn9WbPdMzb2A2tD_IyE4cH3cQOLmJo-maDpHUa_T8oHe3h1VDa2v9cuN88fiHpVT6eBJcLxKKilGiyQclROZDRVYqEppaRUwIVZbWCp1HI4jBUa8mMnk3pDi9ES5ijfFtqjBDsCFzuZu9A1cFllZ-zOealCyWzoEJob_PDyWc8teA0vuosGhmQdEM-tSXXpMLbe4X36fN92QqyUJUzodRTiu-AMF1b4WY5D2sFSmcgiydVElPDYmnIweHgEQpSwBBw404xF9IikU73GhDlhaTGHcfgQayPhtpyDPKFZh92OfB2JOsKfG3CSrxxMvXtd7_DQUXXajI-oA",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			J := NewJWTTokenGen(tt.fields.issuer, tt.fields.privateKey, tt.fields.nowFunc)

			got, err := J.GenerateToken(tt.args.code, tt.args.expire)
			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GenerateToken() got = %v, want %v", got, tt.want)
			}
		})
	}
}
