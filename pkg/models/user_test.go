package models

import (
	"encoding/json"
	"testing"
)

var fullUserJSON = `{
	"id": 1,
	"name": "Leanne Graham",
	"username": "Bret",
	"email": "Sincere@april.biz",
	"address": {
		"street": "Kulas Light",
		"suite": "Apt. 556",
		"city": "Gwenborough",
		"zipcode": "92998-3874",
		"geo": {
			"lat": "-37.3159",
			"lng": "81.1496"
		}
	},
	"phone": "1-770-736-8031 x56442",
	"website": "hildegard.org",
	"company": {
		"name": "Romaguera-Crona",
		"catchPhrase": "Multi-layered client-server neural-net",
		"bs": "harness real-time e-markets"
	}
}`

var fullUser = User{
	ID:       1,
	Name:     "Leanne Graham",
	Username: "Bret",
	Email:    "Sincere@april.biz",
	Address: Address{
		Street:  "Kulas Light",
		Suite:   "Apt. 556",
		City:    "Gwenborough",
		Zipcode: "92998-3874",
		Geo: Geo{
			Lat: "-37.3159",
			Lng: "81.1496",
		},
	},
	Phone:   "1-770-736-8031 x56442",
	Website: "hildegard.org",
	Company: Company{
		Name:        "Romaguera-Crona",
		CatchPhrase: "Multi-layered client-server neural-net",
		Bs:          "harness real-time e-markets",
	},
}

func TestUser_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    User
		wantErr bool
	}{
		{
			name:  "full object",
			input: fullUserJSON,
			want:  fullUser,
		},
		{
			name:  "missing fields default to zero values",
			input: `{}`,
			want:  User{},
		},
		{
			name:  "partial address, no geo",
			input: `{"id":2,"name":"Jane","address":{"street":"Main St","city":"Springfield"}}`,
			want: User{
				ID:   2,
				Name: "Jane",
				Address: Address{
					Street: "Main St",
					City:   "Springfield",
				},
			},
		},
		{
			name:  "extra fields ignored",
			input: `{"id":1,"name":"Test","extra":"ignored"}`,
			want:  User{ID: 1, Name: "Test"},
		},
		{
			name:    "invalid JSON",
			input:   `not json`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got User
			err := json.Unmarshal([]byte(tt.input), &got)
			if tt.wantErr {
				if err == nil {
					t.Error("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestUser_MarshalJSON(t *testing.T) {
	b, err := json.Marshal(fullUser)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got map[string]interface{}
	if err := json.Unmarshal(b, &got); err != nil {
		t.Fatalf("re-unmarshal: %v", err)
	}

	if got["id"].(float64) != 1 {
		t.Errorf("id = %v, want 1", got["id"])
	}
	if got["name"].(string) != "Leanne Graham" {
		t.Errorf("name = %v, want 'Leanne Graham'", got["name"])
	}
	if got["email"].(string) != "Sincere@april.biz" {
		t.Errorf("email = %v, want 'Sincere@april.biz'", got["email"])
	}

	addr := got["address"].(map[string]interface{})
	if addr["street"].(string) != "Kulas Light" {
		t.Errorf("address.street = %v, want 'Kulas Light'", addr["street"])
	}
	geo := addr["geo"].(map[string]interface{})
	if geo["lat"].(string) != "-37.3159" {
		t.Errorf("address.geo.lat = %v, want '-37.3159'", geo["lat"])
	}
	if geo["lng"].(string) != "81.1496" {
		t.Errorf("address.geo.lng = %v, want '81.1496'", geo["lng"])
	}

	company := got["company"].(map[string]interface{})
	if company["name"].(string) != "Romaguera-Crona" {
		t.Errorf("company.name = %v, want 'Romaguera-Crona'", company["name"])
	}
	if company["catchPhrase"].(string) != "Multi-layered client-server neural-net" {
		t.Errorf("company.catchPhrase = %v", company["catchPhrase"])
	}
}

func TestUser_RoundTrip(t *testing.T) {
	b, err := json.Marshal(fullUser)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var restored User
	if err := json.Unmarshal(b, &restored); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if restored != fullUser {
		t.Errorf("round-trip: got %+v, want %+v", restored, fullUser)
	}
}

func TestGeo_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  Geo
	}{
		{"positive coordinates", `{"lat":"37.4220","lng":"-122.0841"}`, Geo{Lat: "37.4220", Lng: "-122.0841"}},
		{"negative coordinates", `{"lat":"-37.3159","lng":"81.1496"}`, Geo{Lat: "-37.3159", Lng: "81.1496"}},
		{"zero coordinates", `{"lat":"0","lng":"0"}`, Geo{Lat: "0", Lng: "0"}},
		{"empty strings", `{"lat":"","lng":""}`, Geo{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got Geo
			if err := json.Unmarshal([]byte(tt.input), &got); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tt.want {
				t.Errorf("got %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestAddress_RoundTrip(t *testing.T) {
	original := Address{
		Street:  "123 Main St",
		Suite:   "Suite 100",
		City:    "Springfield",
		Zipcode: "12345",
		Geo:     Geo{Lat: "40.7128", Lng: "-74.0060"},
	}
	b, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var restored Address
	if err := json.Unmarshal(b, &restored); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if restored != original {
		t.Errorf("round-trip: got %+v, want %+v", restored, original)
	}
}

func TestCompany_RoundTrip(t *testing.T) {
	original := Company{
		Name:        "Acme Corp",
		CatchPhrase: "We do it all",
		Bs:          "synergize paradigms",
	}
	b, err := json.Marshal(original)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	var restored Company
	if err := json.Unmarshal(b, &restored); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if restored != original {
		t.Errorf("round-trip: got %+v, want %+v", restored, original)
	}
}
