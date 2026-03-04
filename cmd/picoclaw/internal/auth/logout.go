// =============================================================================
// logout.go - 登出命令实现
// =============================================================================
// 这个文件实现了PicoClaw的登出功能，用于清除保存的认证凭证。
// 用户可以选择登出特定提供商或所有提供商。
// =============================================================================

package auth

// =============================================================================
// 导入 (Imports)
// =============================================================================
import "github.com/spf13/cobra"

// =============================================================================
// newLogoutCommand - 创建登出命令
// =============================================================================
// 创建"picoclaw auth logout"子命令。
//
// 命令行参数：
//   - --provider, -p: 指定要登出的AI服务提供商（可选）
//     支持的值：openai, anthropic
//     如果不指定，则登出所有提供商
//
// 使用示例：
//
//	# 登出所有提供商
//	picoclaw auth logout
//
//	# 仅登出OpenAI
//	picoclaw auth logout --provider openai
func newLogoutCommand() *cobra.Command {
	var provider string // AI服务提供商（可选）

	// 创建登出命令
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Remove stored credentials",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return authLogoutCmd(provider)
		},
	}

	// 添加命令行标志
	cmd.Flags().StringVarP(&provider, "provider", "p", "", "Provider to logout from (openai, anthropic); empty = all")

	return cmd
}
