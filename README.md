# SSH Terminal

一个基于Web的SSH终端应用，允许用户通过浏览器连接和管理SSH服务器。

## 系统架构

```mermaid
graph TD
    A[用户浏览器] -->|HTTP/WebSocket| B[ssh_web 前端应用]
    B -->|WebSocket| C[ssh_server 后端服务]
    C -->|SSH协议| D[远程SSH服务器]
    
    subgraph 前端
    B -->|Vue.js| E[Terminal组件]
    E -->|Xterm.js| F[终端界面]
    end
    
    subgraph 后端
    C -->|WebSocket Handler| G[终端会话管理]
    G -->|SSH Client| H[SSH连接处理]
    end
```

## 项目结构

项目由两部分组成：

- **ssh_server**: Go语言编写的后端服务，负责处理SSH连接和会话管理
- **ssh_web**: Vue.js编写的前端应用，提供基于浏览器的终端界面

## 功能特点

- 通过浏览器连接SSH服务器
- 实时终端交互体验
- 支持终端窗口大小调整
- 基于WebSocket的实时通信

## 技术栈

### 后端
- Go
- WebSocket
- SSH协议

### 前端
- Vue.js 3
- TypeScript
- Xterm.js
- WebSocket

## 安装要求

- Go 1.18+
- Node.js 16+
- npm 8+

## 快速开始

### 手动启动

1. 启动后端服务器:

```bash
cd ssh_server
go run main.go
```

2. 启动前端应用:

```bash
cd ssh_web
npm install
npm run dev
```

3. 在浏览器中访问 http://localhost:5174

### 使用脚本启动

使用PowerShell一键启动前后端服务:

```powershell
Start-Process -FilePath "powershell" -ArgumentList "-Command cd $PWD\ssh_server; go run main.go" -WindowStyle Normal; Start-Sleep -Seconds 3; Start-Process -FilePath "powershell" -ArgumentList "-Command cd $PWD\ssh_web; npm run dev" -WindowStyle Normal
```

或者使用提供的批处理文件:

```bash
.\start-project.cmd
```

## 配置

SSH连接信息在 `ssh_server/settings.yaml` 文件中配置:

```yaml
system:
  ip: 127.0.0.1  # 后端服务器IP
  port: 8080     # 后端服务器端口
ssh:
  destIP: 192.168.179.135  # 目标SSH服务器IP
  destPort: 22             # 目标SSH服务器端口
  user: root               # SSH用户名
  pwd: ********            # SSH密码
```

## 使用说明

1. 启动应用后，浏览器会自动连接到配置的SSH服务器
2. 终端界面支持所有标准SSH命令
3. 可以通过调整浏览器窗口大小来改变终端大小

## 开发

### 后端开发

```bash
cd ssh_server
go run main.go
```

### 前端开发

```bash
cd ssh_web
npm install
npm run dev
```

## 构建部署

### 前端构建

```bash
cd ssh_web
npm run build
```

构建后的文件将位于 `ssh_web/dist` 目录，可以部署到任何Web服务器。

### 后端构建

```bash
cd ssh_server
go build -o ssh_server
```

