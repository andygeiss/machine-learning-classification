all: clean setup generate_api gather_and_organize_data evaluate_model predict

clean:
	@rm -rf data reports
	@mkdir -p data/external data/interim data/processed reports/figures

evaluate_model:
	@go build -ldflags "-w -s" -o evaluate_model.bin cmd/evaluate_model/main.go
	@upx -9 evaluate_model.bin > /dev/null

gather_and_organize_data:
	@go build -ldflags "-w -s" -o gather_and_organize_data.bin cmd/gather_and_organize_data/main.go
	@upx -9 gather_and_organize_data.bin > /dev/null

generate_api:
	@protoc --go_out=. --python_out=. internal/api/*.proto

predict:
	@go build -ldflags "-w -s" -o predict.bin cmd/predict/main.go
	@upx -9 predict.bin > /dev/null

setup:
	@sudo apt-get install -y protobuf-compiler
	@go install google.golang.org/protobuf/cmd/protoc-gen-go
