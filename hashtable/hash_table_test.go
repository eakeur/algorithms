package hashtable

import (
	"reflect"
	"testing"
)

func Test_hashTable(t *testing.T) {
	type args struct {
		arr map[string]string
		key string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{
				arr: map[string]string{
					"name":        "Igor",
					"age":         "98",
					"surname":     "karkarov",
					"ake":         "18",
					"religion":    "hakmen-ra",
					"since":       "29/10/22",
					"nationality": "Briatmoruntev",
				},
				key: "age",
			},
			want: "98",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tb := NewTable(21)

			for k, v := range tt.args.arr {
				tb.set(k, v)
			}

			if got := tb.get(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hashTable.get() = %v, want %v", got, tt.want)
			}

			tb.delete(tt.args.key)

			if got := tb.get(tt.args.key); !reflect.DeepEqual(got, nil) {
				t.Errorf("hashTable.get() = %v, want %v", got, nil)
			}
		})
	}
}

func Benchmark_hashTable(b *testing.B) {
	tb := NewTable(5)

	for k, v := range map[string]string{
		"name":        "Igor",
		"age":         "98",
		"surname":     "karkarov",
		"ake":         "18",
		"religion":    "hakmen-ra",
		"since":       "29/10/22",
		"nationality": "Briatmoruntev",
		"graduation":  "MBA in Software Architecture",
		"a":           "b",
		"c":           "d",
		"e":           "f",
		"g":           "h",
	} {
		tb.set(k, v)
	}

	b.Run("read_key_1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tb.get("age")
		}
	})

	b.Run("read_key_2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tb.get("a")
		}
	})

	b.Run("insert_key", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tb.set("surname", i)
		}
	})

	b.Run("delete_key", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tb.delete("surname")
		}
	})

}
