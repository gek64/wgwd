@echo off
chcp 65001

set base=%~dp0

if not "%1"=="am_admin" (powershell start -verb runas '%0' am_admin & exit /b)
cd /d %base%

set W_REMOTE_INTERFACE="xxx"
set W_WG_INTERFACE="xxx"
set W_ENDPOINT="xxx"
set W_ACCESS_KEY_ID="xxx"
set W_SECRET_ACCESS_KEY="xxx"
set W_BUCKET="xxx"
set W_OBJECT_PATH="xxx"
set W_ENCRYPTION_KEY="xxx"
set N_INTERVAL="5m"

wgwd get -remote_interface=%W_REMOTE_INTERFACE% -wg_interface=%W_WG_INTERFACE% -interval=%N_INTERVAL% s3 -endpoint=%W_ENDPOINT% -path_style -access_key_id=%W_ACCESS_KEY_ID% -secret_access_key=%W_SECRET_ACCESS_KEY% -bucket=%W_BUCKET% -object_path=%W_OBJECT_PATH% -encryption_key=%W_ENCRYPTION_KEY%
