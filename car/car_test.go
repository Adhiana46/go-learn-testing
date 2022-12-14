package car

import "testing"

func TestDefaultEngine_MaxSpeed(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "default engine should have 50 max speed",
			want: 50,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := DefaultEngine{}
			if got := e.MaxSpeed(); got != tt.want {
				t.Errorf("MaxSpeed() = %d, want %d", got, tt.want)
			}
		})
	}
}

type FakeEngine struct{}

func (e FakeEngine) MaxSpeed() int {
	return 5
}

func TestCar_Speed(t *testing.T) {
	type fields struct {
		Speeder Speeder
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "speed should be 50, when use Default Engine",
			fields: fields{
				Speeder: &DefaultEngine{},
			},
			want: 50,
		},
		{
			name: "speed should be 80, when use Turbo Engine",
			fields: fields{
				Speeder: &TurboEngine{},
			},
			want: 80,
		},
		{
			name: "speed should be 20, when maxspeed is less than 10",
			fields: fields{
				Speeder: &FakeEngine{},
			},
			want: 20,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Car{
				Speeder: tt.fields.Speeder,
			}
			if got := c.Speed(); got != tt.want {
				t.Errorf("Speed() = %d, want %d", got, tt.want)
			}
		})
	}
}
