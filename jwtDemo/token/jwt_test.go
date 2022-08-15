package jwt

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v4"
	"testing"
	"time"
)

func TestJWTTokenGen_GenerateToken(t *testing.T) {
	privateKey := "-----BEGIN RSA PRIVATE KEY-----\nMIIEowIBAAKCAQEAkuBEzR3UtdLx/j6zEy802VgLREt0vqg/XxxqSFEXH6tn/ziL\nDGwjswododYTh5RR/HUtwwLX+ErgNxWtAe6J1lWo/LTePZ9YyOTA1rpizSbivKc4\nwTkhqB52aDUxBxpFqKaPf2Hr9nLmuNhHZGpYXYfbe0Ko5lFk+g1SoE3eCR2AKluf\nCdJi+bQb5utAEsU2AXW2FXFr/plI15Msv/lrJxKi9EUWqbAVVb17ufj7qeQAQnQY\nulPfsN4NmysF2AgLL3BRSO1AVi4hREZb0ykY2vr6QHFbCJicQXBWIn3I0pfS9pqm\njIXpsk9KNQj2b0bXF9RXGXvqauY5hPkPY1zq1wIDAQABAoIBAFMUlvZrrFmZVBdb\nwhHU/xD+M4lTIVuDci1Ltnbnn6vjWjVM96J0PRFWZIKZxHxJksK6ScNohV/GguPF\n2BIiq0dZNviTGq9CxocElmwgDCk9mL3b3Ue3GGmvqYyyoeQBoWuPo9YLPt0uC4J9\nnIyZdxiY87bacw2s9vBWvuk8gp8medbs1D5blid/KCkMeMOs4PqPCYxUKOi6srEf\n3MG2XpgyWRRXYRDFe4jzyZ8EgKE5TodN7EMk8UbopJbTjubDtK59lNhM7G7xam3+\nKcjt5wF5MBYBItjVsdjW0PP4uhEwv+pS5EQqGVjUUBiWbOzaNdGxqHiT4DZGrop+\nrvlwccECgYEA4g4H9lHkIwaVbreqLM4w72caz/oOf+XhBCPUZsZ4/qgd3gNDbFPL\nzBVK9567lXxjXJSPwqkPuU7q+6uQcdaPbqb/NR1b2x2dLTGP75E8/qRsaWr8HwDt\nDP+10svQsa7kVtPKQKyZKC9Ehkhpfz1ksBNL7NwDiseoC5aQkVAc0zcCgYEAplUg\n0YCHFR7C8ajma0CB3+RV22MI5DRDJvl2BBubncAJGQGaxKNemqca0Sd/HauNVW1L\nD52IYAEpA4G7z2KmgOGeq9tzMuErTKY2q2jBR+J1VeK+L96JSDUijGErBZmBk7DS\n0x11bemQ4Xz37jnWeimgd1+Bf1ctfVvqrMU9tWECgYBvWlALwz0pD8YueuSmG4+5\nbkWj616XHAriui+XzO6vKutDgMIGq6R5TIfYQGXQ9iwP39KNJjBrrPdNr66AlRuk\nezi91k94OTdCmuZ3MUkleqM4ro2xTAh5XcEYLHQKAsfSCZPYle0AAYC00Ri+p+Rg\nY7ED5UENbNpXkvx1yEvaswKBgEd0ZqbZqi4+isy6HFKS4bhHJUMcJmOyt+50WGin\nW5DHQKFHeS9kZ23Uv8Z+SzN50TuTuh05Na7YkL/66a3L2W+gfUOjALgo1ysUdIPq\neq3g8Ts2LzHtTdiBSxg3Xf4H1gyWd9tN2nafUigsQW5L1oGghFkD5GPKFba21/fk\nu//BAoGBAIVW+HNip/p2L4jjXz1XlI8ZcJ8xp1H63TSiclji8NLPYQFLCzCkvJq/\niWRsIcKnwOxb0UFhjKIyAx7xhloarFPdiBf+Icth51wBf1v8OhhVcY2Re/gu+pRv\nH+9qxpNjQZM6yT4w46VNlg750qE/30iq2EK138dEoucLfRhWRU3O\n-----END RSA PRIVATE KEY-----"

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
			want:    "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MTYyNDI2MjIsImlhdCI6MTUxNjIzOTAyMiwiaXNzIjoid3d3LmRhcnItZW4xLnRvcCIsInN1YiI6ImRhcnJfZW4xIn0.NR7IIOqPL09JGQ2ZA5ZM2TNreBVTtbgzt8-XdMUAOMMVf6fgESyXdqhX5Du-PaKNhLlWaI0JthNLIkrgPKcNbXoXMJqL3f4mlHzIq5KwnCmEQDRMXN-mBsbOz6qNIy1ny1RI1C5Wj5sV6-knL84VA8GYHJhNUnzOXe8diVic-cWaWHb1NJz6el45UB0mCOXu01HI7A3l7krxxGUw4MXznFMXqmc5mrWRKfqFmp8ZsnwrPOD9ZKU7xWm6rGseqT6haGS6Ci_bxH8WonKx236sw1hNCGFSvHHMJhQSfgtyFWYPRrUeVu6kT0q-0UCQU4BAr7NJqY5bhJ8rdYDX48FnAA",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			J := JWTTokenGen{
				privateKey: tt.fields.privateKey,
				issuer:     tt.fields.issuer,
				nowFunc:    tt.fields.nowFunc,
			}
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
