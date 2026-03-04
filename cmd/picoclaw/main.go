// =============================================================================
// PicoClaw - 超轻量级个人AI助手
// =============================================================================
// 这是一个用Go语言编写的超轻量级个人AI助手项目。
// 项目灵感来源于 nanobot (https://github.com/HKUDS/nanobot)，
// 通过AI驱动的自我引导过程，从零开始用Go重构。
//
// 主要特点：
// - 超轻量级：<10MB内存占用，比OpenClaw小99%
// - 极低成本：可以在10美元的硬件上运行
// - 快速启动：1秒内启动，即使在0.6GHz单核上
// - 真正可移植：单一二进制文件，支持RISC-V、ARM和x86架构
//
// 许可证：MIT
// 版权所有 (c) 2026 PicoClaw 贡献者
//
// 致谢：
// - 感谢 nanobot 项目提供的灵感
// - 感谢所有为PicoClaw贡献代码的开发者
// =============================================================================

package main

// =============================================================================
// 导入 (Imports)
// =============================================================================
// Go语言的标准库和第三方库
import (
	"fmt" // 用于格式化输出，如打印版本信息
	"os"  // 用于操作系统相关操作，如退出程序

	// spf13/cobra: 一个强大的CLI库，用于构建命令行应用程序
	// PicoClaw使用Cobra来组织所有的子命令（如auth, cron, skills等）
	"github.com/spf13/cobra"

	// 内部包导入 - 这些是PicoClaw项目的核心模块
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal"         // 内部帮助函数和工具
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/agent"   // AI代理管理模块
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/auth"    // 认证模块
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/cron"    // 定时任务模块
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/gateway" // 网关模块
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/migrate" // 数据迁移模块
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/onboard" // 入门引导模块
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/skills"  // 技能模块
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/status"  // 状态查看模块
	"github.com/sipeed/picoclaw/cmd/picoclaw/internal/version" // 版本信息模块
)

// =============================================================================
// NewPicoclawCommand - 创建主命令
// =============================================================================
// 这是PicoClaw CLI的核心函数，它创建并配置主命令对象。
// Cobra使用"命令"的概念来组织CLI功能，每个子命令都是一个独立的操作。
//
// 工作原理：
// 1. 创建一个新的cobra.Command对象
// 2. 设置命令的基本属性（名称、简短描述、示例）
// 3. 添加所有子命令到主命令中
// 4. 返回配置好的命令对象，供main函数执行
//
// 子命令说明：
// - onboard: 首次使用引导，帮助用户配置PicoClaw
// - agent: 管理AI代理（创建、配置、删除等）
// - auth: 处理认证相关功能（登录、登出、状态）
// - gateway: 管理消息网关（启动、停止网关服务）
// - status: 查看PicoClaw的运行状态
// - cron: 管理定时任务（添加、列出、删除定时任务）
// - migrate: 数据迁移工具（从其他系统迁移配置）
// - skills: 技能管理（安装、列出、搜索技能）
// - version: 显示版本信息
func NewPicoclawCommand() *cobra.Command {
	// 简短描述字符串，包含Logo和版本号
	// internal.Logo 是PicoClaw的标志（小龙虾 ��）
	// internal.GetVersion() 获取当前版本号
	short := fmt.Sprintf("%s picoclaw - Personal AI Assistant v%s\n\n", internal.Logo, internal.GetVersion())

	// 创建Cobra命令对象
	// Use: 命令的使用方式，用户在终端输入的命令
	// Short: 简短描述，在帮助信息中显示
	// Example: 使用示例
	cmd := &cobra.Command{
		Use:     "picoclaw",
		Short:   short,
		Example: "picoclaw list",
	}

	// =============================================================================
	// 添加子命令 (Add Subcommands)
	// =============================================================================
	// 每个子命令都通过对应的NewXxxCommand()函数创建
	// 这些函数分布在不同的internal包中，每个包负责特定的功能领域
	cmd.AddCommand(
		onboard.NewOnboardCommand(), // 引导命令
		agent.NewAgentCommand(),     // 代理管理命令
		auth.NewAuthCommand(),       // 认证命令
		gateway.NewGatewayCommand(), // 网关命令
		status.NewStatusCommand(),   // 状态命令
		cron.NewCronCommand(),       // 定时任务命令
		migrate.NewMigrateCommand(), // 迁移命令
		skills.NewSkillsCommand(),   // 技能命令
		version.NewVersionCommand(), // 版本命令
	)

	return cmd
}

// =============================================================================
// 常量定义 - ASCII艺术横幅
// =============================================================================
// 这些常量定义了PicoClaw启动时显示的彩色ASCII艺术横幅
// 使用ANSI转义序列来实现彩色输出
//
// ANSI转义序列说明：
// \033[1;38;2;R;G;Bm - 设置前景色为RGB值
// \033[0m - 重置所有格式
//
// 图案说明：
// 左边蓝色部分显示"PicoClaw"的ASCII艺术
// 右边红色部分显示"AI Assistant"的ASCII艺术
const (
	colorBlue = "\033[1;38;2;62;93;185m" // 蓝色，用于PicoClaw文字
	colorRed  = "\033[1;38;2;213;70;70m" // 红色，用于"AI Assistant"文字
	banner    = "\r\n" +
		colorBlue + "██████╗ ██╗ ██████╗ ██████╗ " + colorRed + " ██████╗██╗      █████╗ ██╗    ██╗\n" +
		colorBlue + "██╔══██╗██║██╔════╝██╔═══██╗" + colorRed + "██╔════╝██║     ██╔══██╗██║    ██║\n" +
		colorBlue + "██████╔╝██║██║     ██║   ██║" + colorRed + "██║     ██║     ███████║██║ █╗ ██║\n" +
		colorBlue + "██╔═══╝ ██║██║     ██║   ██║" + colorRed + "██║     ██║     ██╔══██║██║███╗██║\n" +
		colorBlue + "██║     ██║╚██████╗╚██████╔╝" + colorRed + "╚██████╗███████╗██║  ██║╚███╔███╔╝\n" +
		colorBlue + "╚═╝     ╚═╝ ╚═════╝ ╚═════╝ " + colorRed + " ╚═════╝╚══════╝╚═╝  ╚═╝ ╚══╝╚══╝\n " +
		"\033[0m\r\n"
)

// =============================================================================
// main 函数 - 程序入口点
// =============================================================================
// 这是PicoClaw程序的入口点，当用户运行picoclaw命令时首先执行此函数。
//
// 程序执行流程：
// 1. 打印ASCII艺术横幅（显示PicoClaw标志和版本）
// 2. 创建主命令对象（包含所有子命令）
// 3. 执行命令并处理可能的错误
// 4. 如果执行出错，以退出码1退出程序
//
// 重要提示：
// - Go程序的main函数不能有返回值，所以错误处理通过os.Exit实现
// - cmd.Execute()会解析命令行参数并调用相应的子命令
// - 退出码0表示成功，非0表示失败
func main() {
	fmt.Printf("%s", banner) // 打印启动横幅
	cmd := NewPicoclawCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1) // 执行出错，退出程序
	}
}
