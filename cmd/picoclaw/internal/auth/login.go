// =============================================================================
// login.go - 登录命令实现
// =============================================================================
// 这个文件实现了PicoClaw的登录功能，支持多种AI服务提供商的OAuth认证。
// 用户可以通过浏览器登录或使用设备码流程（适用于无头环境）。
// =============================================================================

package auth

// =============================================================================
// 导入 (Imports)
// =============================================================================
import "github.com/spf13/cobra"

// =============================================================================
// newLoginCommand - 创建登录命令
// =============================================================================
// 创建"picoclaw auth login"子命令。
//
// 命令行参数：
//   - --provider, -p: 指定要登录的AI服务提供商（必填）
//     支持的值：openai, anthropic, google-antigravity
//   - --device-code: 使用设备码流程登录（适用于无头环境/服务器）
//
// 使用示例：
//
//	# 使用浏览器登录OpenAI
//	picoclaw auth login --provider openai
//
//	# 使用设备码登录（无头环境）
//	picoclaw auth login --provider openai --device-code
func newLoginCommand() *cobra.Command {
	// 定义命令行标志变量
	var (
		provider      string // AI服务提供商
		useDeviceCode bool   // 是否使用设备码流程
	)

	// 创建登录命令
	// Use: 命令名称
	// Short: 简短描述
	// Args: 参数验证，cobra.NoArgs表示不接受任何位置参数
	// RunE: 执行函数，调用authLoginCmd处理实际登录逻辑
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Login via OAuth or paste token",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return authLoginCmd(provider, useDeviceCode)
		},
	}

	// 添加命令行标志
	// StringVarP: 字符串类型标志，-p简写形式
	// BoolVar: 布尔类型标志
	cmd.Flags().StringVarP(&provider, "provider", "p", "", "Provider to login with (openai, anthropic)")
	cmd.Flags().BoolVar(&useDeviceCode, "device-code", false, "Use device code flow (for headless environments)")

	// 标记provider为必填参数
	_ = cmd.MarkFlagRequired("provider")

	return cmd
}
