package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func Test_getTopMenuItems(t *testing.T) {
	type args struct {
		entries []Entry
		count   int
	}
	tests := []struct {
		name    string
		args    args
		want    []MenuItem
		wantErr bool
	}{
		{
			name: "valid use",
			args: args{
				entries: []Entry{
					{EaterID: 1, FoodMenuID: 1},
					{EaterID: 1, FoodMenuID: 2},
					{EaterID: 2, FoodMenuID: 1},
					{EaterID: 2, FoodMenuID: 3},
					{EaterID: 3, FoodMenuID: 1},
					{EaterID: 3, FoodMenuID: 4},
				},
				count: 3,
			},
			want: []MenuItem{
				{FoodMenuID: 1, Count: 3},
				{FoodMenuID: 2, Count: 1},
				{FoodMenuID: 3, Count: 1},
			},
			wantErr: false,
		},
		{
			name: "duplicate combination",
			args: args{
				entries: []Entry{
					{EaterID: 1, FoodMenuID: 1},
					{EaterID: 1, FoodMenuID: 1},
					{EaterID: 1, FoodMenuID: 2},
					{EaterID: 2, FoodMenuID: 1},
					{EaterID: 2, FoodMenuID: 3},
					{EaterID: 3, FoodMenuID: 1},
					{EaterID: 3, FoodMenuID: 4},
				},
				count: 3,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getTopMenuItems(tt.args.entries, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("getTopMenuItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}

func Test_getTopNMenuItems(t *testing.T) {
	type args struct {
		items []MenuItem
		n     int
	}
	tests := []struct {
		name string
		args args
		want []MenuItem
	}{
		{
			name: "valid use",
			args: args{
				items: []MenuItem{
					{FoodMenuID: 1, Count: 17},
					{FoodMenuID: 4, Count: 2},
					{FoodMenuID: 3, Count: 11},
					{FoodMenuID: 2, Count: 21},
					{FoodMenuID: 5, Count: 5},
				},
				n: 3,
			},
			want: []MenuItem{
				{FoodMenuID: 2, Count: 21},
				{FoodMenuID: 1, Count: 17},
				{FoodMenuID: 3, Count: 11},
			},
		},
		{
			name: "valid use with n greater than length of slice",
			args: args{
				items: []MenuItem{
					{FoodMenuID: 1, Count: 17},
					{FoodMenuID: 3, Count: 11},
					{FoodMenuID: 2, Count: 21},
				},
				n: 4,
			},
			want: []MenuItem{
				{FoodMenuID: 2, Count: 21},
				{FoodMenuID: 1, Count: 17},
				{FoodMenuID: 3, Count: 11},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getTopNMenuItems(tt.args.items, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getTopNMenuItems() = %v, want %v", got, tt.want)
			}
		})
	}
}
