package car

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

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

type MockEngine struct {
	mock.Mock
}

func (m MockEngine) MaxSpeed() int {
	args := m.Called()
	return args.Get(0).(int)
}

func TestCar_Speed_WithMock(t *testing.T) {
	t.Run("Called 1 Times", func(t *testing.T) {
		mock := new(MockEngine)

		car := NewCar(mock)

		mock.On("MaxSpeed").Return(9).Times(1)

		speed := car.Speed()
		assert.Equal(t, 20, speed)

		// untuk memastikan bahwa function MaxSpeed hanya dipanggil satu kali
		// ketika memanggil function Speed di car
		mock.AssertExpectations(t)
	})

	t.Run("Called 3 Times", func(t *testing.T) {
		mock := new(MockEngine)

		car := NewCar(mock)

		mock.On("MaxSpeed").Return(70).Times(3)

		speed := car.Speed()
		assert.Equal(t, 70, speed)

		// untuk memastikan bahwa function MaxSpeed hanya dipanggil satu kali
		// ketika memanggil function Speed di car
		mock.AssertExpectations(t)
	})
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
