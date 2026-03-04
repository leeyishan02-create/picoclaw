// =============================================================================
// auth 包 - 认证管理模块
// =============================================================================
// 这个包负责PicoClaw的用户认证功能。
// 支持多种AI提供商的OAuth登录方式，包括：
// - OpenAI
// - Anthropic (Claude)
// - Google Antigravity
//
// 主要功能：
// - 登录：通过OAuth或粘贴token进行认证
// - 登出：清除保存的认证凭证
// - 状态：查看当前认证状态
// - 模型：查看可用的模型列表
// =============================================================================

package auth

// =============================================================================
// 导入 (Imports)
// =============================================================================
import "github.com/spf13/cobra"

// =============================================================================
// NewAuthCommand - 创建认证命令
// =============================================================================
// 这是auth模块的入口函数，创建认证相关的子命令。
// 当用户运行"picoclaw auth"时会调用此函数。
//
// 子命令说明：
// - login: 登录到AI服务提供商
// - logout: 登出并清除认证凭证
// - status: 查看当前认证状态
// - models: 查看可用的模型列表
func NewAuthCommand() *cobra.Command {
	// 创建认证命令
	// Use: 命令名称，用户输入"picoclaw auth"
	// Short: 简短描述
	// RunE: 默认行为，显示帮助信息
	cmd := &cobra.Command{
		Use:   "auth",
		Short: "Manage authentication (login, logout, status)",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return cmd.Help()
		},
	}

	// 添加子命令
	cmd.AddCommand(
		newLoginCommand(),  // 登录命令
		newLogoutCommand(), // 登出命令
		newStatusCommand(), // 状态命令
		newModelsCommand(), // 模型命令
	)

	return cmd
}
