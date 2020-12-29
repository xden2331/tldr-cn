package main

import (
	"flag"
	"os"
	"reflect"
	"testing"
)

func Test_readJSONData(t *testing.T) {
	tests := []struct {
		name    string
		want    []command
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "testJSONFile.json",
			want: []command{
				command{
					"tar",
					[]string{"压缩", "tar"},
					"通常用来压缩/解压文件",
					[]example{
						example{
							"压缩一系列文件到同一压缩文件",
							"tar cf <filename.tar> file1 file2 file3"},
					}},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readJSONData(jsonpath)
			if (err != nil) != tt.wantErr {
				t.Errorf("readJSONData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readJSONData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getArgs(t *testing.T) {
	tests := []struct {
		name    string
		want    arguments
		wantErr bool
		osArgs  []string
	}{
		// TODO: Add test cases.
		{
			name:    "Set List",
			want:    arguments{listAll: true, query: ""},
			wantErr: false,
			osArgs:  []string{"cmd", "--list"},
		},
		{
			name:    "Set query",
			want:    arguments{listAll: false, query: "tar"},
			wantErr: false,
			osArgs:  []string{"cmd", "tar"},
		},
		{
			name:    "Set List and query together",
			want:    arguments{},
			wantErr: true,
			osArgs:  []string{"cmd", "--list", "tar"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Store the original args
			actualArgs := os.Args
			// Reset args
			defer func() {
				os.Args = actualArgs
				flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			}()

			os.Args = tt.osArgs
			got, err := getArgs()
			if (err != nil) != tt.wantErr {
				t.Errorf("getArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fetchInformation(t *testing.T) {
	type args struct {
		commands []command
		args     arguments
	}
	commands := []command{
		command{
			"tar",
			[]string{"压缩", "tar"},
			"通常用来压缩/解压文件",
			[]example{
				example{
					"压缩一系列文件到同一压缩文件",
					"tar cf <filename.tar> file1 file2 file3"},
			}},
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "List all command",
			args: args{
				commands: commands,
				args: arguments{
					listAll: true,
					query:   "",
				},
			},
			want:    "Supported commands:\n-tar\n\t通常用来压缩/解压文件\n",
			wantErr: false,
		},
		{
			name: "Show tar command",
			args: args{
				commands: commands,
				args: arguments{
					listAll: false,
					query:   "tar",
				},
			},
			want:    "tar\n通常用来压缩/解压文件\n- 压缩一系列文件到同一压缩文件\n\ttar cf <filename.tar> file1 file2 file3\n",
			wantErr: false,
		},
		{
			name: "Not supported command",
			args: args{
				commands: commands,
				args: arguments{
					listAll: false,
					query:   "zip",
				},
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fetchInformation(tt.args.commands, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("fetchInformation() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("fetchInformation() = %v, want %v", got, tt.want)
			}
		})
	}
}
