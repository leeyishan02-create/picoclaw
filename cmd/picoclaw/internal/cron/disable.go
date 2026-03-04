// =============================================================================
// disable.go - 禁用定时任务命令实现
// =============================================================================
// 这个文件实现了PicoClaw的禁用定时任务功能。
// =============================================================================

package cron

// =============================================================================
// 导入 (Imports)
// =============================================================================
import "github.com/spf13/cobra"

// =============================================================================
// newDisableCommand - 创建禁用命令
// =============================================================================
// 创建"picoclaw cron disable"子命令。
// 根据任务ID禁用定时任务。
//
// 参数：
//   - storePath: 返回定时任务存储文件路径的函数
//
// 位置参数：
//   - args[0]: 任务ID
//
// 使用示例：
//
//	picoclaw cron disable 1
func newDisableCommand(storePath func() string) *cobra.Command {
	return &cobra.Command{
		Use:     "disable",
		Short:   "Disable a job",
		Args:    cobra.ExactArgs(1), // 必须提供1个参数（任务ID）
		Example: `picoclaw cron disable 1`,
		RunE: func(_ *cobra.Command, args []string) error {
			cronSetJobEnabled(storePath(), args[0], false)
			return nil
		},
	}
}
