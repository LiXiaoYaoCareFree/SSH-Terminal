@echo off
echo 正在启动SSH终端项目...

:: 启动后端服务器（在新窗口中）
echo 启动后端服务器...
start "SSH后端服务器" cmd /c "cd %~dp0ssh_server && go run main.go"

:: 等待后端启动
echo 等待后端服务启动...
timeout /t 3 /nobreak > nul

:: 启动前端应用（在新窗口中）
echo 启动前端应用...
start "SSH前端应用" cmd /c "cd %~dp0ssh_web && npm run dev"

echo.
echo 项目已启动!
echo - 后端服务运行在: http://127.0.0.1:8080
echo - 前端应用运行在: http://localhost:5174
echo. 