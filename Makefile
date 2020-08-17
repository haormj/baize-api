# variable
binaryName=baize-api
versionPath=github.com/haormj/version
version=v0.1.0
outputPath=_output
workingDirectory=`pwd`
ldLibraryPath=.:${LD_LIBRARY_PATH}:/usr/local/libtorch/lib:/usr/local/lib64

all: build

proto:
	protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. idl/dogsvscats/dogs_vs_cats.proto

build: proto
	@buildTime=`date "+%Y-%m-%d %H:%M:%S"`; \
	CGO_CXXFLAGS="-g -O2 -D_GLIBCXX_USE_CXX11_ABI=1 \
	              -I /usr/local/libtorch/include \
	              -I /usr/local/libtorch/include/torch/csrc/api/include \
	              -I /usr/local/include/opencv4" \
	CGO_LDFLAGS="-g -O2 \
	             -L/usr/local/libtorch/lib -L/usr/local/lib64 -lc10 -ltorch -lopencv_calib3d -lopencv_core \
	             -lopencv_dnn -lopencv_features2d -lopencv_flann -lopencv_gapi -lopencv_highgui \
	             -lopencv_imgcodecs -lopencv_imgproc -lopencv_ml -lopencv_objdetect -lopencv_photo \
	             -lopencv_stitching -lopencv_video -lopencv_videoio \
	             -lhaormj_vision -lhaormj_http -lcurl" \
	go build -ldflags "-X '${versionPath}.Version=${version}' \
	                   -X '${versionPath}.BuildTime=$$buildTime' \
	                   -X '${versionPath}.GoVersion=`go version`' \
	                   -X '${versionPath}.GitCommit=`git rev-parse --short HEAD`'" -o ${outputPath}/${binaryName}; \

run: build
	LD_LIBRARY_PATH=${ldLibraryPath} ./_output/baize-api

clean:
	rm -rf _output
	rm -rf logs

.ONESHELL:
systemd: 
	cat > /etc/systemd/system/micro-api.service << EOF
	[Unit]
	Description=micro api
	[Service]
	LimitNOFILE=40960
	LimitNPROC=40960
	Type=simple
	ExecStart=/path/to/micro api --handler=rpc --enable_cors false
	Restart=always
	RestartSec=2s
	MemoryLimit=1G
	[Install]
	WantedBy=multi-user.target
	EOF
	systemctl daemon-reload
	systemctl enable micro-api
	systemctl restart micro-api

	cat > /etc/systemd/system/${binaryName}.service << EOF
	[Unit]
	Description=${binaryName}
	After=micro-api.service
	Requires=micro-api.service
	[Service]
	LimitNOFILE=40960
	LimitNPROC=40960
	Type=simple
	Environment=LD_LIBRARY_PATH=${ldLibraryPath}
	WorkingDirectory=${workingDirectory}
	ExecStart=${workingDirectory}/${outputPath}/${binaryName}
	Restart=always
	RestartSec=2s
	MemoryLimit=1G
	[Install]
	WantedBy=multi-user.target
	EOF
	systemctl daemon-reload
	systemctl enable ${binaryName}
	systemctl restart ${binaryName}

systemd-status:
	systemctl status micro-api
	systemctl status ${binaryName}

systemd-stop:
	systemctl stop micro-api
	systemctl stop ${binaryName}

.PHONY: all proto build run systemd clean
