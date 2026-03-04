// =============================================================================
// enable.go - 启用定时任务命令实现
// =============================================================================
// 这个文件实现了PicoClaw的启用定时任务功能。
// =============================================================================

package cron

// =============================================================================
// 导入 (Imports)
// =============================================================================
import "github.com/spf13/cobra"

// =============================================================================
// newEnableCommand - 创建启用命令
// =============================================================================
// 创建"picoclaw cron enable"子命令。
// 根据任务ID启用定时任务。
//
// 参数：
//   - storePath: 返回定时任务存储文件路径的函数
//
// 位置参数：
//   - args[0]: 任务ID
//
// 使用示例：
//
//	picoclaw cron enable 1
func newEnableCommand(storePath func() string) *cobra.Command {
	return &cobra.Command{
		Use:     "enable",
		Short:   "Enable a job",
		Args:    cobra.ExactArgs(1), // 必须提供1个参数（任务ID）
		Example: `picoclaw cron enable 1`,
		RunE: func(_ *cobra.Command, args []string) error {
			cronSetJobEnabled(storePath(), args[0], true)
			return nil
		},
	}
}
