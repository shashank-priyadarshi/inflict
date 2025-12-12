BUF ?= buf
SQLC ?= sqlc
TMP_DIR ?= tmp

ROOT_DIR := $(CURDIR)

ifeq ($(OS),Windows_NT)
	# 1. FORCE make to use cmd.exe on Windows. 
	# This prevents it from using w64devkit's 'sh.exe' which breaks "if exist" commands.
	SHELL := cmd.exe
	
	# Windows Commands
	MKDIR_CMD = if not exist $(TMP_DIR) mkdir $(TMP_DIR)
	RM_CMD = if exist $(TMP_DIR) rmdir /s /q $(TMP_DIR)
else
	# Linux Commands
	MKDIR_CMD = mkdir -p $(TMP_DIR)
	RM_CMD = rm -rf $(TMP_DIR)
endif

.PHONY: codegen gen-enums db clean create protos check-drift

codegen: clean create protos gen-enums db check-drift

clean:
	$(RM_CMD)

create:
	$(MKDIR_CMD)

protos:
	$(BUF) generate --config ./protos/buf.yaml --template ./protos/buf.gen.yaml
	$(BUF) build --config ./protos/buf.yaml -o $(TMP_DIR)/proto.descriptor.bin

gen-enums:
	go -C tools/enumgen run . -descriptor "$(ROOT_DIR)/$(TMP_DIR)/proto.descriptor.bin" -out "$(ROOT_DIR)/api/schema/db/v1/migrations/_generated_enums.sql"

db:
	$(SQLC) generate -f ./api/schema/db/v1/sqlc.yaml

check-drift:
	go -C tools/checkdrift run . -enums "$(ROOT_DIR)/api/schema/db/v1/migrations/_generated_enums.sql" -descriptor "$(ROOT_DIR)/$(TMP_DIR)/proto.descriptor.bin"
	go -C api test -v ./internal/mapper/...