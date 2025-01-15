package service

import "testing"

func Test_convertMenuToAPI(t *testing.T) {
	type args struct {
		menuName string
	}
	tests := []struct {
		name       string
		args       args
		wantPath   string
		wantMethod string
	}{
		{
			name: "test1",
			args: args{
				menuName: "system:role:get:menus",
			},
			wantPath:   "/api/system/role/:id/menus",
			wantMethod: "GET",
		},
		{
			name: "test2",
			args: args{
				menuName: "system:dict:list",
			},
			wantPath:   "/api/system/dict",
			wantMethod: "GET",
		},
		{
			name: "test3",
			args: args{
				menuName: "system:dict:detail",
			},
			wantPath:   "/api/system/dict/:id",
			wantMethod: "GET",
		},
		{
			name: "test4",
			args: args{
				menuName: "system:dict:delete",
			},
			wantPath:   "/api/system/dict/:ids",
			wantMethod: "DELETE",
		},
		{
			name: "test5",
			args: args{
				menuName: "permission:menu:tree",
			},
			wantPath:   "/api/permission/menu/tree",
			wantMethod: "GET",
		},
		{
			name: "test6",
			args: args{
				menuName: "permission:role:set:permissions",
			},
			wantPath:   "/api/permission/role/:id/permissions",
			wantMethod: "PUT",
		},
		{
			name: "test7",
			args: args{
				menuName: "system:role:get:menus",
			},
			wantPath:   "/api/system/role/:id/menus",
			wantMethod: "GET",
		},
		{
			name: "test8",
			args: args{
				menuName: "system:role:update",
			},
			wantPath:   "/api/system/role/:id",
			wantMethod: "PUT",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPath, gotMethod := convertMenuToAPI(tt.args.menuName)
			if gotPath != tt.wantPath {
				t.Errorf("convertMenuToAPI() gotPath = %v, want %v", gotPath, tt.wantPath)
			}
			if gotMethod != tt.wantMethod {
				t.Errorf("convertMenuToAPI() gotMethod = %v, want %v", gotMethod, tt.wantMethod)
			}
		})
	}
}
