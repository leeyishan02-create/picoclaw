// =============================================================================
// gateway 包 - 网关模块
// =============================================================================
// 这个包负责PicoClaw的网关（Gateway）功能。
// 网关是PicoClaw的核心服务，负责：
// - 接收来自各种消息通道的消息
// - 将消息路由到相应的AI代理
// - 处理AI响应并发送回消息通道
// - 管理消息队列和速率限制
// =============================================================================

package gateway

// =============================================================================
// 导入 (Imports)
// =============================================================================
import (
	"github.com/spf13/cobra" // CLI命令库
)

// =============================================================================
// NewGatewayCommand - 创建网关命令
// =============================================================================
// 这是gateway模块的入口函数，创建网关启动命令。
// 当用户运行"picoclaw gateway"时会调用此函数。
//
// 命令行参数：
//   - --debug, -d: 启用调试日志
//
// 使用示例：
//
//	# 启动网关
//	picoclaw gateway
//	picoclaw g
//
//	# 启用调试模式启动网关
//	picoclaw gateway --debug
func NewGatewayCommand() *cobra.Command {
	var debug bool // 是否启用调试模式

	// 创建网关命令
	cmd := &cobra.Command{
		Use:     "gateway",
		Aliases: []string{"g"}, // 命令别名
		Short:   "Start picoclaw gateway",
		Args:    cobra.NoArgs, // 不接受任何位置参数
		RunE: func(_ *cobra.Command, _ []string) error {
			return gatewayCmd(debug)
		},
	}

	// 添加命令行标志
	cmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug logging")

	return cmd
}
