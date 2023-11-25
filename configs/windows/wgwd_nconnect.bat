@echo off
chcp 65001

set base=%~dp0

if not "%1"=="am_admin" (powershell start -verb runas '%0' am_admin & exit /b)
cd /d %base%

set W_REMOTE_INTERFACE="xxx"
set W_WG_INTERFACE="xxx"
set W_ID="xxx"
set W_ENDPOINT="https://xxx"
set N_INTERVAL="5m"

wgwd get -remote_interface=%W_REMOTE_INTERFACE% -wg_interface=%W_WG_INTERFACE% -interval=%N_INTERVAL% nconnect -id=%W_ID% -endpoint=%W_ENDPOINT% -allow_insecure
