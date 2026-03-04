// =============================================================================
// version 包 - 版本信息模块
// =============================================================================
// 这个包负责PicoClaw的版本信息显示功能。
// =============================================================================

package version

// =============================================================================
// 导入 (Imports)
// =============================================================================
import (
	"fmt" // 格式化输出

	"github.com/spf13/cobra" // CLI命令库

	"github.com/sipeed/picoclaw/cmd/picoclaw/internal" // 内部工具函数
)

// =============================================================================
// NewVersionCommand - 创建版本命令
// =============================================================================
// 这是version模块的入口函数，创建版本信息命令。
// 当用户运行"picoclaw version"时会调用此函数。
//
// 功能说明：
// - 显示PicoClaw版本号
// - 显示Git提交哈希（如果有）
// - 显示构建时间（如果有）
// - 显示Go版本
//
// 使用示例：
//
//	picoclaw version
//	picoclaw v
func NewVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"}, // 命令别名
		Short:   "Show version information",
		Run: func(_ *cobra.Command, _ []string) {
			printVersion()
		},
	}

	return cmd
}

// =============================================================================
// printVersion - 打印版本信息
// =============================================================================
// 打印PicoClaw的完整版本信息。
// 包括版本号、Git提交哈希、构建时间和Go版本。
func printVersion() {
	// 打印版本号和Logo
	fmt.Printf("%s picoclaw %s\n", internal.Logo, internal.FormatVersion())

	// 获取并打印构建信息
	build, goVer := internal.FormatBuildInfo()
	if build != "" {
		fmt.Printf("  Build: %s\n", build)
	}
	if goVer != "" {
		fmt.Printf("  Go: %s\n", goVer)
	}
}
