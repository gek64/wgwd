@echo off
chcp 65001

set base=%~dp0

if not "%1"=="am_admin" (powershell start -verb runas '%0' am_admin & exit /b)
cd /d %base%

set W_REMOTE_INTERFACE="xxx"
set W_WG_INTERFACE="xxx"
set W_ENDPOINT="xxx"
set W_USERNAME="xxx"
set W_PASSWORD="xxx"
set W_FILEPATH="xxx"
set W_ENCRYPTION_KEY="xxx"
set N_INTERVAL="5m"

wgwd get -remote_interface=%W_REMOTE_INTERFACE% -wg_interface=%W_WG_INTERFACE% -interval=%N_INTERVAL% webdav -endpoint=%W_ENDPOINT% -username=%W_USERNAME% -password=%W_PASSWORD% -filepath=%W_FILEPATH% -encryption_key=%W_ENCRYPTION_KEY%
